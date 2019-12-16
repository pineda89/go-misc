package lib_gc_cache_source

import (
	"errors"
	"sync"

	LOG "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
)

var mutex *sync.Mutex = &sync.Mutex{}
var dmutex *sync.Mutex = &sync.Mutex{}

func init() {
	ICacheMap = &MapCache{make(map[string][]byte)}

	// Register the cache source
	CacheSourceContainer.AddICacheSoure("MAP", ICacheMap)

	LOG.Info.Println("lib_cache_source (map adapter) package initialized. ")
}

var ICacheMap ICacheSource

type MapCache struct {
	cache map[string][]byte
}

func (mp *MapCache) InitSource() error { return nil }

func (mp *MapCache) CloseSource() error { return nil }

func (mp *MapCache) AddToHash(hash *string, key, value *[]byte) error {
	return errors.New("method not implemented !!!")
}

func (mp *MapCache) GetKeysFromHash(hash *string) (*[][]byte, error) {
	return nil, errors.New("method not implemented !!!")
}

func (mp *MapCache) RemoveKeyFromHash(hash *string, key *[]byte) error {
	return errors.New("method not implemented !!!")
}

func (mp *MapCache) RemoveHash(hash *string) error {
	return errors.New("method not implemented !!!")
}

func (mp *MapCache) AddData(key, value *[]byte) error {
	mutex.Lock()
	defer mutex.Unlock()

	mp.cache[string(*key)] = *value
	return nil
}

func (mp *MapCache) AddMData(data *[][]byte) error {
	return errors.New("method not implemented !!!")
}

func (mp *MapCache) GetData(key *[]byte) (*[]byte, error) {
	if data, ok := mp.cache[string(*key)]; ok {
		return &data, nil
	} else {
		return nil, errors.New("The key" + string(*key) + " does not exists !!!")
	}
}

func (mp *MapCache) MGetData(keys *[][]byte) (*[][]byte, error) {
	return nil, errors.New("method not implemented !!!")
}

func (mp *MapCache) DeleteData(key *[]byte) (bool, error) {
	dmutex.Lock()
	defer dmutex.Unlock()

	skey := string(*key)
	if _, ok := mp.cache[skey]; ok {
		delete(mp.cache, skey)
		return true, nil
	} else {
		return false, nil
	}
}
