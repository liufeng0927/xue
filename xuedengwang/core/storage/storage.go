package storage

import (
	"errors"
	"io"
	"os"
	"time"
)

// Config COS 存储配置
type Config struct {
	SecretID  string
	SecretKey string
	Host      string
	Bucket    string
	Protocol  string
}

type FileInfo struct {
	Size    int64
	ModTime time.Time
	Mode    os.FileMode
}

type Meta struct {
	ContentType   string
	ContentLength int32
	Etag          string
	ReqID         string
}

var (
	ErrObjectNotFound            = errors.New("object not found")
	ErrObjectKeyInvalid          = errors.New("invalid object key")
	ErrObjectWritePermissionDeny = errors.New("no write permission")
	ErrObjectReadPermissionDeny  = errors.New("no read permission")
	ErrObjectEmptyContent        = errors.New("zero content")
	ErrStorageUnRegister         = errors.New("unregister storage")
)

type Storage interface {

	// Put 保存data至某个文件
	Put(key string, r io.Reader) error

	// PutByPath 根据某个文件写入到另一个文件里
	PutByPath(key string /*要带上路径，相对路径或绝对路径都行*/, src string) error

	// FileStream 获取语音流
	FileStream(key string) (io.ReadCloser, *FileInfo, error)

	// Get 获取数据
	Get(key string, wa io.WriterAt) error

	// GetToPath 获取到某个文件
	GetToPath(key string /*要带上路径，相对路径或绝对路径都行*/, dest string) error

	// Stat 获取文件信息  大小，修改时间，权限
	Stat(key string) (*FileInfo, error)

	// Del 删除文件
	Del(key string) error

	// Size 获取文件大小
	Size(key string) (int64, error)

	// IsExist 判断文件是否存在
	IsExist(key string) (bool, error)

	// Meta 获取对象元数据
	Meta(key string) (*Meta, error)
}
