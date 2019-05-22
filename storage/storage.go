package storage

import (
	"fmt"
	"time"
)

type Storage interface {
	Get(string) (string, error)
	Put(string, string) error
}

var (
	ErrKeyAlreadyExists = fmt.Errorf("key already exists")
	ErrKeyNotFound      = fmt.Errorf("key does not exists")

	s Storage
)

func Init(useDB bool, addr string, pwd string, db int, rwTimeout time.Duration) error {
	var err error

	if useDB {
		ss, err := NewDBStorage(addr, pwd, db, rwTimeout)
		if err != nil {
			return err
		}
		s = *ss
	} else {
		s, err = NewInmemoryStorage()
		if err != nil {
			return err
		}
	}

	s.Put("testhash", "full-url/for-testing-this")
	return nil
}

func Get(key string) (string, error) {
	return s.Get(key)
}

func Put(key, value string) error {
	return s.Put(key, value)
}
