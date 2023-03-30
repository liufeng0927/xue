package cos

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"xuedengwang/core/storage"

	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type client struct {
	ctx    context.Context
	client *cos.Client
}

func NewClient(cfg storage.Config, ctx context.Context) storage.Storage {

	storageUrl := fmt.Sprintf("%s://%s.%s", cfg.Protocol, cfg.Bucket, cfg.Host)

	u, _ := url.Parse(storageUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretID,
			SecretKey: cfg.SecretKey,
		},
	})

	return &client{
		ctx:    ctx,
		client: c,
	}
}

func (s *client) PutByPath(key string, src string) error {
	fmt.Printf("[COS PUT BY PATH] object: %s \n", key)

	fd, _, err := storage.OpenLocal(src)
	if err != nil {
		return err
	}
	defer fd.Close()

	return s.Put(key, fd)
}

func (s *client) Put(key string, r io.Reader) error {
	fmt.Printf("[COS PUT] object: %s \n", key)

	opt := &cos.ObjectPutOptions{}
	_, err := s.client.Object.Put(s.ctx, key, r, opt)
	if err != nil {
		return err
	}

	return nil
}

func (s *client) GetToPath(key string, dest string) error {
	fmt.Printf("[COS GET TO PATH] object: %s \n", key)

	dir, _ := filepath.Split(dest)
	_ = os.MkdirAll(dir, 0766)
	fd, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
	if err != nil {
		return err
	}
	defer fd.Close()

	return s.Get(key, fd)
}

func (s *client) Get(key string, wa io.WriterAt) error {
	fmt.Printf("[COS GET] object: %s \n", key)

	if !storage.ValidKey(key) {
		return storage.ErrObjectKeyInvalid
	}

	opt := &cos.ObjectGetOptions{}

	v, err := s.client.Object.Get(s.ctx, key, opt)
	if err != nil {
		return err
	}

	return storage.Copy(wa, v.Body)
}

func (s *client) FileStream(key string) (io.ReadCloser, *storage.FileInfo, error) {
	fmt.Printf("[COS FILE STREAM] object: %s \n", key)

	if !storage.ValidKey(key) {
		return nil, nil, storage.ErrObjectKeyInvalid
	}

	opt := &cos.ObjectGetOptions{}

	output, err := s.client.Object.Get(s.ctx, key, opt)
	if err != nil {
		if output != nil {
			if output.StatusCode == http.StatusNotFound {
				return nil, nil, storage.ErrObjectNotFound
			}
		}
		return nil, nil, err
	}

	modTime, _ := time.Parse(time.RFC1123, output.Header.Get("Last-Modified"))

	return output.Body, &storage.FileInfo{
		Size:    output.ContentLength,
		ModTime: modTime.In(time.Local),
		Mode:    0666,
	}, nil
}

func (s *client) Stat(key string) (*storage.FileInfo, error) {
	fmt.Printf("[COS STAT] object: %s \n", key)

	if !storage.ValidKey(key) {
		return nil, storage.ErrObjectKeyInvalid
	}

	opt := &cos.ObjectGetOptions{}

	output, err := s.client.Object.Get(s.ctx, key, opt)
	if err != nil {
		if output != nil {
			if output.StatusCode == http.StatusNotFound {
				return nil, storage.ErrObjectNotFound
			}
		}
		return nil, err
	}

	modTime, _ := time.Parse(time.RFC1123, output.Header.Get("Last-Modified"))

	return &storage.FileInfo{
		Size:    output.ContentLength,
		ModTime: modTime.In(time.Local),
		Mode:    0666,
	}, nil
}

func (s *client) Size(key string) (int64, error) {
	fmt.Printf("[COS SIZE] object: %s \n", key)

	if !storage.ValidKey(key) {
		return 0, storage.ErrObjectKeyInvalid
	}

	opt := &cos.ObjectGetOptions{}

	output, err := s.client.Object.Get(s.ctx, key, opt)
	if err != nil {
		if output != nil {
			if output.StatusCode == http.StatusNotFound {
				return 0, storage.ErrObjectNotFound
			}
		}
		return 0, err
	}

	if output.ContentLength == 0 {
		return 0, fmt.Errorf("failed to get object size with code %d", output.StatusCode)
	}

	return output.ContentLength, nil
}

func (s *client) IsExist(key string) (bool, error) {
	fmt.Printf("[COS IS EXIST] object: %s \n", key)

	if !storage.ValidKey(key) {
		return false, storage.ErrObjectKeyInvalid
	}

	return s.client.Object.IsExist(s.ctx, key)
}

func (s *client) Del(key string) error {
	fmt.Printf("[COS DEL] object: %s \n", key)

	if !storage.ValidKey(key) {
		return storage.ErrObjectKeyInvalid
	}

	resp, err := s.client.Object.Delete(s.ctx, key)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return storage.ErrObjectNotFound
	}

	return nil
}

func (s *client) Meta(key string) (*storage.Meta, error) {
	opt := &cos.ObjectHeadOptions{}
	resp, err := s.client.Object.Head(s.ctx, key, opt)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, storage.ErrObjectNotFound
	}

	tag := resp.Header.Get("ETag")
	return &storage.Meta{
		ContentType:   resp.Header.Get("Content-Type"),
		ContentLength: cast.ToInt32(resp.Header.Get("Content-Length")),
		Etag:          tag[1 : len(tag)-1],
		ReqID:         resp.Header.Get("X-Cos-Request-Id"),
	}, nil
}
