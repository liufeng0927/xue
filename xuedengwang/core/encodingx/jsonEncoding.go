package encodingx

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/golang/snappy"
)

// JsonEncoding json格式
type JsonEncoding struct{}

// Marshal json encode
func (j JsonEncoding) Marshal(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	return buf, err
}

// Unmarshal json decode
func (j JsonEncoding) Unmarshal(data []byte, value interface{}) error {
	err := json.Unmarshal(data, value)
	if err != nil {
		return err
	}
	return nil
}

// JsonGzipEncoding json and gzip
type JsonGzipEncoding struct{}

// Marshal json encode and gzip
func (jz JsonGzipEncoding) Marshal(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	// var bufSizeBefore = len(buf)

	buf, err = GzipEncode(buf)
	// log.Infof("gzip_json_compress_ratio=%d/%d=%.2f", bufSizeBefore, len(buf), float64(bufSizeBefore)/float64(len(buf)))
	return buf, err
}

// Unmarshal json encode and gzip
func (jz JsonGzipEncoding) Unmarshal(data []byte, value interface{}) error {
	jsonData, err := GzipDecode(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, value)
	if err != nil {
		return err
	}
	return nil
}

// GzipEncode 编码
func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer, err := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(in)
	if err != nil {
		err = writer.Close()
		if err != nil {
			return out, err
		}
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

// GzipDecode 解码
func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer func() {
		err = reader.Close()
		if err != nil {
			fmt.Printf("reader close err: %+v", err)
		}
	}()

	return ioutil.ReadAll(reader)
}

// JsonSnappyEncoding json格式和snappy压缩
type JsonSnappyEncoding struct{}

// Marshal 序列化
func (s JsonSnappyEncoding) Marshal(v interface{}) (data []byte, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	d := snappy.Encode(nil, b)
	return d, nil
}

// Unmarshal 反序列化
func (s JsonSnappyEncoding) Unmarshal(data []byte, value interface{}) error {
	b, err := snappy.Decode(nil, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, value)
}
