package storage

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

type DBStorage struct {
	client *redis.Client
}

func NewDBStorage(addr string, pwd string, db int, rwTimeout time.Duration) (*DBStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd, // no password set
		DB:           db,  // use default DB
		ReadTimeout:  rwTimeout,
		WriteTimeout: rwTimeout,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &DBStorage{client}, nil
}

func (s DBStorage) Put(key, value string) error {
	cmdres := s.client.SetNX(key, value, 0)
	ok, err := cmdres.Result()
	if err != nil {
		log.Printf("redis error: %s", err)
		return err
	}
	if !ok {
		return ErrKeyAlreadyExists
	}

	return nil
}

// todo add url store duration

func (s DBStorage) Get(key string) (string, error) {
	val, err := s.client.Get(key).Result()
	if err == redis.Nil {
		return "", ErrKeyNotFound
	} else if err != nil {
		log.Printf("redis error: %s", err)
		return "", err
	}
	return val, nil
}
