package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type FileCacheDriver struct {
	path string
}

func NewFileCacheDriver(path string) *FileCacheDriver {
	return &FileCacheDriver{path: path}
}

func (c *FileCacheDriver) Set(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	filename := c.filenameForKey(key)
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	if ttl > 0 {
		go func() {
			time.Sleep(ttl)
			os.Remove(filename)
		}()
	}
	return nil
}

func (c *FileCacheDriver) Get(key string) (interface{}, error) {
	filename := c.filenameForKey(key)
	data, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var value interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, err
	}
	return value, nil
}

func (c *FileCacheDriver) Delete(key string) error {
	filename := c.filenameForKey(key)
	return os.Remove(filename)
}

func (c *FileCacheDriver) filenameForKey(key string) string {
	return filepath.Join(c.path, key+".json")
}
