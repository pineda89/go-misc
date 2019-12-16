package lib_gc_cache_source

import (
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"

	"errors"
	"sync"
)

func init() {
	CacheSourceContainer = &cacheSourceContainer{make(map[string]ICacheSource)}
}

var CacheSourceContainer ICacheSourceContainer
var mutex1 *sync.Mutex = &sync.Mutex{}
var mutex2 *sync.Mutex = &sync.Mutex{}

type ICacheSource interface {
	InitSource() error
	CloseSource() error
	AddData(key, value *[]byte) error
	AddMData(data *[][]byte) error
	GetData(key *[]byte) (*[]byte, error)
	MGetData(keys *[][]byte) (*[][]byte, error)
	DeleteData(key *[]byte) (bool, error)

	// Hash methdos to manage index keys
	AddToHash(hash *string, key, value *[]byte) error
	GetKeysFromHash(hash *string) (*[][]byte, error)
	RemoveKeyFromHash(hash *string, key *[]byte) error
	RemoveHash(hash *string) error
}

type ICacheSourceContainer interface {
	AddICacheSoure(key string, cacheSource ICacheSource) error
	GetICacheSource(key string) (ICacheSource, error)
	RemoveICacheSource(key string) error
}

type cacheSourceContainer struct {
	container map[string]ICacheSource
}

func (csc *cacheSourceContainer) AddICacheSoure(key string, cacheSource ICacheSource) error {
	mutex1.Lock()
	defer mutex1.Unlock()

	csc.container[key] = cacheSource
	return nil
}

func (csc *cacheSourceContainer) GetICacheSource(key string) (ICacheSource, error) {

	if cacheSource, ok := csc.container[key]; !ok {
		msg, _ := EVENT.NotifyEvent("013-018", "", &[]string{key})
		return nil, errors.New(msg)
	} else {
		return cacheSource, nil
	}
}

func (csc *cacheSourceContainer) RemoveICacheSource(key string) error {
	mutex2.Lock()
	defer mutex2.Unlock()

	if _, ok := csc.container[key]; !ok {
		msg, _ := EVENT.NotifyEvent("013-017", "", &[]string{key})
		return errors.New(msg)
	} else {
		delete(csc.container, key)
		return nil
	}
}
