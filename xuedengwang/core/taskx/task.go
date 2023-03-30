package taskx

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	"go.uber.org/atomic"
)

var (
	DefaultChannelSize    = 64                   // task channel buffer size
	DefaultExecuteTimeout = time.Second * 5      // execute timeout
	DefaultTimeout        = time.Hour * 24 * 356 // default timeout
	DefaultUpdateInterval = time.Second * 1      // update interval
	ErrTimeout            = errors.New("time out")
	ErrTaskPanic          = errors.New("task panic")
	ErrTaskFailed         = errors.New("task failed")
	ErrTaskFulled         = errors.New("task fulled")
	ErrTaskNotRunning     = errors.New("task not running")
)

type any = interface{}

type Handler func(context.Context, ...any) error
type Task struct {
	ctx    context.Context // control run timeout
	fn     Handler         // handle function
	err    chan<- error    // error returned
	params []any           // params in order
}

type Tasker struct {
	opts     *TaskerOptions
	tasks    chan *Task
	stopChan chan struct{}
	ticker   *time.Ticker
	stopOnce *sync.Once
	running  atomic.Bool
}

func NewTasker() *Tasker {
	return &Tasker{
		opts:     &TaskerOptions{},
		tasks:    make(chan *Task, DefaultChannelSize),
		stopChan: make(chan struct{}, 1),
		ticker:   time.NewTicker(DefaultUpdateInterval),
		stopOnce: new(sync.Once),
	}
}

func (t *Tasker) Init(opts ...TaskerOption) {
	t.opts = defaultTaskerOptions()

	for _, o := range opts {
		o(t.opts)
	}

	if t.opts.chanBufSize > DefaultChannelSize {
		t.tasks = make(chan *Task, t.opts.chanBufSize)
	}

	t.running.Store(true)
	t.ticker.Reset(t.opts.updateInterval)
}

func (t *Tasker) ResetTimer() {
	tm := t.opts.timer
	if tm != nil && !tm.Stop() {
		<-tm.C
	}
	tm.Reset(t.opts.d)
}

func (t *Tasker) GetTaskNum() int {
	return len(t.tasks)
}

func (t *Tasker) IsRunning() bool {
	return t.running.Load()
}

func (t *Tasker) AddWait(ctx context.Context, f Handler, p ...any) error {
	if !t.IsRunning() {
		return ErrTaskNotRunning
	}

	if len(t.tasks) >= t.opts.chanBufSize {
		return ErrTaskFulled
	}

	subCtx, cancel := context.WithTimeout(ctx, t.opts.executeTimeout)
	defer cancel()

	e := make(chan error, 1)
	tk := &Task{
		ctx:    subCtx,
		fn:     f,
		err:    e,
		params: make([]any, 0, len(p)),
	}
	tk.params = append(tk.params, p...)

	// send channel
	select {
	case t.tasks <- tk:
		break
	case <-subCtx.Done():
		if subCtx.Err() == nil {
			return nil
		}
		return fmt.Errorf("task send to channel with timeout:%w, chan buff size:%d", subCtx.Err(), len(t.tasks))
	}

	// wait channel result
	select {
	case err := <-e:
		if err == nil {
			return nil
		}
		return fmt.Errorf("task wait channel result with error:%w, chan buff size:%d", err, len(t.tasks))
	case <-subCtx.Done():
		if subCtx.Err() == nil {
			return nil
		}
		return fmt.Errorf("task wait channel result with timeout:%w, chan buff size:%d", subCtx.Err(), len(t.tasks))
	}
}

func (t *Tasker) Add(ctx context.Context, f Handler, p ...any) {
	if !t.IsRunning() {
		return
	}

	if len(t.tasks) >= t.opts.chanBufSize {
		return
	}

	tk := &Task{
		ctx:    ctx,
		fn:     f,
		err:    nil,
		params: make([]any, 0, len(p)),
	}
	tk.params = append(tk.params, p...)
	t.tasks <- tk
}

func (t *Tasker) Run(ctx context.Context) (reterr error) {
	t.ResetTimer()

	defer func() {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			t.opts.logger.Printf("catch exception:%v, panic recovered with stack:%s", err, stack)
			reterr = ErrTaskPanic
		}
		t.stop()
	}()

	if len(t.opts.startFns) > 0 {
		for _, fn := range t.opts.startFns {
			fn()
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case h, ok := <-t.tasks:
			if !ok {
				return nil
			} else {
				select {
				case <-h.ctx.Done():
					continue
				default:
				}

				err := h.fn(h.ctx, h.params...)
				if h.err != nil {
					h.err <- err // handle result
				}

				if err != nil {
					funcName := runtime.FuncForPC(reflect.ValueOf(h.fn).Pointer()).Name()
					t.opts.logger.Printf("execute %s with error:%v\n", funcName, err)
				}
			}

		case <-t.opts.timer.C:
			return ErrTimeout

		case <-t.ticker.C:
			if t.opts.updateFn != nil {
				// grace stop task when update
				if !t.opts.updateFn() {
					return
				}
			}
		}

	}
}

func (t *Tasker) Stop() {
	t.stopOnce.Do(func() {
		close(t.tasks)
		t.ticker.Stop()
		<-t.stopChan
	})
}

func (t *Tasker) stop() {
	if len(t.opts.stopFns) > 0 {
		for _, fn := range t.opts.stopFns {
			fn()
		}
	}

	t.opts.timer.Stop()
	t.running.Store(false)

	select {
	case <-t.stopChan:
		return
	default:
		close(t.stopChan)
	}
}
