package storage

import "sync"

type InmemoryStorage struct {
	data map[string]string

	lock sync.RWMutex
}

func NewInmemoryStorage() (Storage, error) {
	return InmemoryStorage{
		data: make(map[string]string, 1000),
	}, nil
}

func (s InmemoryStorage) Get(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return val, nil
}

func (s InmemoryStorage) Put(key, value string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.data[key]
	if ok {
		return ErrKeyAlreadyExists
	}

	s.data[key] = value
	return nil
}
