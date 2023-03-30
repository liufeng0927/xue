package taskx

import (
	"io"
	"log"
	"time"
)

type StartFn func()
type StopFn func()
type UpdateFn func() bool

type TaskerOption func(*TaskerOptions)
type TaskerOptions struct {
	uniqueId       int32
	startFns       []StartFn // 任务开始callback
	stopFns        []StopFn  // 任务停止callback
	updateFn       UpdateFn  // 默认更新callback
	timer          *time.Timer
	d              time.Duration // 超时时间
	updateInterval time.Duration // 更新间隔时间
	executeTimeout time.Duration // 执行超时时间
	chanBufSize    int           // 通道缓冲区大小
	logger         *log.Logger
}

func defaultTaskerOptions() *TaskerOptions {
	return &TaskerOptions{
		uniqueId:       0,
		d:              DefaultTimeout,
		startFns:       make([]StartFn, 0, 5),
		stopFns:        make([]StopFn, 0, 5),
		updateFn:       nil,
		timer:          time.NewTimer(DefaultTimeout),
		updateInterval: DefaultUpdateInterval,
		executeTimeout: DefaultExecuteTimeout,
		chanBufSize:    DefaultChannelSize,
		logger:         log.Default(),
	}
}

func WithStartFns(f ...StartFn) TaskerOption {
	return func(o *TaskerOptions) {
		o.startFns = o.startFns[:0]
		o.startFns = append(o.startFns, f...)
	}
}

func WithStopFns(f ...StopFn) TaskerOption {
	return func(o *TaskerOptions) {
		o.stopFns = o.stopFns[:0]
		o.stopFns = append(o.stopFns, f...)
	}
}

func WithUpdateFn(f UpdateFn) TaskerOption {
	return func(o *TaskerOptions) {
		o.updateFn = f
	}
}

func WithTimeout(d time.Duration) TaskerOption {
	return func(o *TaskerOptions) {
		o.d = d
	}
}

func WithUpdateInterval(d time.Duration) TaskerOption {
	return func(o *TaskerOptions) {
		o.updateInterval = d
	}
}

func WithExecuteTimeout(d time.Duration) TaskerOption {
	return func(o *TaskerOptions) {
		o.executeTimeout = d
	}
}

func WithChannelBufferSize(sz int) TaskerOption {
	return func(o *TaskerOptions) {
		o.chanBufSize = sz
	}
}

func WithUniqueId(id int32) TaskerOption {
	return func(o *TaskerOptions) {
		o.uniqueId = id
	}
}

func WithOutput(output io.Writer) TaskerOption {
	return func(o *TaskerOptions) {
		o.logger = log.New(output, "tasker: ", log.Lmsgprefix|log.LstdFlags)
		log.SetOutput(output)
	}
}
