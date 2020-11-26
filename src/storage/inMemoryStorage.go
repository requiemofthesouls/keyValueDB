package storage

import (
	"errors"
	"fmt"
)

type InMemoryKVStorage struct {
	storage map[string]string
}

func (K InMemoryKVStorage) NewStorage() KVStorage {
	fmt.Println("Building new inmemory storage")
	return InMemoryKVStorage{storage: make(map[string]string)}
}

// check if key exists
func (K InMemoryKVStorage) keyExists(key string) bool {
	_, found := K.storage[key]
	return found
}

func (K InMemoryKVStorage) CreateString(key, value string) error {
	if exists := K.keyExists(key); exists {
		return errors.New("key already exists")
	}

	K.storage[key] = value
	return nil
}

func (K InMemoryKVStorage) ReadString(key string) (string, error) {
	if exists := K.keyExists(key); !exists {
		return "", errors.New("key not found")
	}

	v, _ := K.storage[key]
	return v, nil
}

func (K InMemoryKVStorage) UpdateString(key, value string) error {
	if exists := K.keyExists(key); !exists {
		return errors.New("key not found")
	}

	K.storage[key] = value
	return nil
}

func (K InMemoryKVStorage) DeleteString(key string) error {
	if exists := K.keyExists(key); !exists {
		return errors.New("key not found")
	}

	delete(K.storage, key)
	return nil
}

func (K InMemoryKVStorage) PrintDump() error {
	fmt.Println("==================================================")
	fmt.Println(K.storage)
	fmt.Println("==================================================")
	return nil
}

func (K InMemoryKVStorage) String() string {
	return "InMemory storage"
}
