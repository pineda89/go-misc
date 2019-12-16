package lib_gc_cache

import (
	CACHEHELPER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_cache/lib_gc_cache_helpers"
	CACHESOURCE "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_cache/lib_gc_cache_source"
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	IDX_VALUE "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_idx/lib_gc_idx_value/protobuf_protocols/capnp"
	LOG "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
	UTIL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_util"

	"bytes"
	"errors"
	"fmt"
	"reflect"
)

func init() {

}

var DOCache DOCacheAdapter

var cacheName string

// Get cache adapter
func GetCacheAdapter() (DOCacheAdapter, error) {
	if icache, err := CACHESOURCE.CacheSourceContainer.GetICacheSource(cacheName); err != nil {
		panic(err)
	} else {
		return &doCache{icache}, err
	}
}

//
// ----------- Serializable interface
//
type MakeSerializable interface {
	ToBytes() (*[]byte, error)
	FromBytes(_b *[]byte) (interface{}, error)
}

//
// ----------- MakeCacheable interface
//
type MakeCacheable interface {
	GetPK() string
	GetName() string
	GetMetadata() *map[string]interface{}
	SetMetadataField(field string, value interface{}) error
	Clone() interface{}
}

//
// ----------- Cache Adapter interface
//
type DOCacheAdapter interface {
	getDaoPKFromCacheable(cacheable MakeCacheable) (*[]byte, error)
	getDaoPK(key, name string) (*[]byte, error)

	// Init source
	InitSource() error

	// CO methods
	Put(data interface{}) error
	GetFromCacheable(cacheable MakeCacheable) (interface{}, error)
	MGetFromCacheable(cacheables []MakeCacheable) (*[]interface{}, error)
	GetFromKey(key string, dataType interface{}) (interface{}, error)
	RemoveFromCacheable(cacheable MakeCacheable) (bool, error)

	// Index methods
	AddValueToIndexKey(indexkey *string, value MakeSerializable, indexname string) error
	GetValuesFromIndexKey(indexkey *string, dataType MakeSerializable, indexname string) (interface{}, error)
	RemoveValueFromIndexKey(indexkey *string, value MakeSerializable, indexname *string) error
	RemoveIndexKey(indexkey, indexname *string) error

	Close() error
}

//
// ----------- DOCacheAdapter Implementation
//

func makeDaoPK(cacheableName, key string) (*[]byte, error) {
	var buff bytes.Buffer
	buff.WriteString(cacheableName)
	buff.WriteString(">>")
	buff.WriteString(key)
	//s_key := buff.String()

	s_key := []byte(buff.String())

	return &s_key, nil
}

// Cache struct
type doCache struct {
	cache CACHESOURCE.ICacheSource
}

// Close the cache
func (docache *doCache) Close() error {
	docache.cache.CloseSource()
	return nil
}

// Get a cache specific PK
func (docache *doCache) getDaoPKFromCacheable(cacheable MakeCacheable) (*[]byte, error) {
	name := cacheable.GetName()
	pk := cacheable.GetPK()
	return docache.getDaoPK(pk, name)
}

func (docache *doCache) getDaoPK(pk, name string) (*[]byte, error) {
	if len(pk) > 0 && len(name) > 0 {
		return makeDaoPK(name, pk)
	} else {
		msg, _ := EVENT.NotifyEvent("013-008", "", &[]string{pk, name})
		return nil, errors.New(msg)
	}
}

//
// ----- Init source
//
func (docache *doCache) InitSource() error {

	// Take environment variables
	_cacheName := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_CACHE__CACHENAME", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_STRING}
	evs := []*UTIL.EnvironmentVariable{_cacheName}
	if _, err := UTIL.GetEnvironmentVariables(evs); err != nil {
		panic(err)
	} else {
		cacheName = _cacheName.Var_value.(string)
	}

	// Start the DO cache adapter
	if icache, err := CACHESOURCE.CacheSourceContainer.GetICacheSource(cacheName); err != nil {
		return err
	} else {
		DOCache = &doCache{icache}
	}

	if err := docache.cache.InitSource(); err != nil {
		return err
	} else {
		LOG.Info.Println("lib_gc_cache package initialized. ")
		return nil
	}
}

//
// ----- Index methods
//

// Add a value to index key
func (docache *doCache) AddValueToIndexKey(indexkey *string, value MakeSerializable, indexname string) error {

	if b_value, err2 := value.ToBytes(); err2 == nil {
		//		t := time.Now()
		pk := []byte(value.(MakeCacheable).GetPK())
		if err3 := docache.cache.AddToHash(indexkey, b_value, &pk); err3 != nil {
			return err3
		} else {
			// Note indicators. Send the indicators
			/*	TO DO : Update indicators
				if indicators_channel, err := SM.HGetIndicatorsChannel(MLS.ML_MONITOR_NAME); err != nil {
					return err
				} else {

						indicators_channel <- &SM.StatusAlteraion{MLS.TOTAL_IDX_BYTES_WRITED, SM.INDICATOR_OPERATION_INC, int64(len(*b_value) + len(*b_value) + len(pk))}
						indicators_channel <- &SM.StatusAlteraion{MLS.TOTAL_IDX_PUTS, SM.INDICATOR_OPERATION_INC, int64(1)}
						indicators_channel <- &SM.StatusAlteraion{MLS.LAST_IDX_PUT_TIME, SM.INDICATOR_OPERATION_SET, time.Now().Sub(t)}

				}
			*/
		}
	} else {
		// It not has been possible to add the data
		msg, _ := EVENT.NotifyEvent("013-001", "", &[]string{err2.Error()})
		return errors.New(msg)
	}
	return nil
}

// Get values from index key
func (docache *doCache) GetValuesFromIndexKey(indexkey *string, dataType MakeSerializable, indexname string) (interface{}, error) {

	if values, err1 := docache.cache.GetKeysFromHash(indexkey); err1 == nil {
		__values := make([]IDX_VALUE.DOIndexValue, len(*values))
		var ds_value interface{}
		for i, ___value := range *values {
			ds_value, _ = dataType.FromBytes(&___value)
			__values[i] = *ds_value.(*IDX_VALUE.DOIndexValue)
		}
		return &__values, nil
	} else {
		// It not has been possible to marshal the data
		msg, _ := EVENT.NotifyEvent("013-009", "", &[]string{err1.Error()})
		return nil, errors.New(msg)
	}
	return nil, nil
}

func (docache *doCache) RemoveValueFromIndexKey(indexkey *string, value MakeSerializable, indexname *string) error {

	if b_pk, err2 := value.ToBytes(); err2 == nil {
		if err3 := docache.cache.RemoveKeyFromHash(indexkey, b_pk); err3 != nil {
			return err3
		} else {
			return nil
		}
	} else {
		// It not has been possible to add the data
		msg, _ := EVENT.NotifyEvent("013-015", "", &[]string{*indexkey, err2.Error()})
		return errors.New(msg)
	}
	return nil
}

func (docache *doCache) RemoveIndexKey(indexkey, indexname *string) error {

	if err := docache.cache.RemoveHash(indexkey); err != nil {
		return err
	} else {
		return nil
	}

	return nil
}

//
// ------------------- CO methods.
//

// Put a DO into cache
func (docache *doCache) Put(data interface{}) error {

	// Check for cacheable
	if cacheable, ok := data.(MakeCacheable); ok {
		// Check for record metadata

		if err := docache.checkForCacheableMetadataStatus(cacheable); err != nil {
			return err
		} else {
			if b_pk, err1 := docache.getDaoPKFromCacheable(cacheable); err1 == nil {
				if b, err2 := data.(MakeSerializable).ToBytes(); err2 == nil {
					//					t := time.Now()
					if err3 := docache.cache.AddData(b_pk, b); err3 == nil {
						// Note indicators. Send the indicators
						/*	TO DO : Update indicators
							if indicators_channel, err := SM.HGetIndicatorsChannel(MLS.ML_MONITOR_NAME); err != nil {
								return err
							} else {
								indicators_channel <- &SM.StatusAlteraion{MLS.TOTAL_BYTES_WRITED, SM.INDICATOR_OPERATION_INC, int64(len(*b))}
								indicators_channel <- &SM.StatusAlteraion{MLS.TOTAL_PUTS, SM.INDICATOR_OPERATION_INC, int64(1)}
								indicators_channel <- &SM.StatusAlteraion{MLS.LAST_CACHE_PUT_TIME, SM.INDICATOR_OPERATION_SET, time.Now().Sub(t)}
							}
						*/
					} else {
						// It not has been possible to add the data
						msg, _ := EVENT.NotifyEvent("013-001", "", &[]string{err3.Error()})
						return errors.New(msg)
					}
				} else {
					// It not has been possible to marshal the data
					msg, _ := EVENT.NotifyEvent("013-002", "", &[]string{err2.Error()})
					return errors.New(msg)
				}
			} else {
				// It has not been posssible to marshal the pk
				msg, _ := EVENT.NotifyEvent("013-003", "", &[]string{err1.Error()})
				return errors.New(msg)
			}
		}
	} else {
		// The DO does not implements the MakeCacheable interface
		msg, _ := EVENT.NotifyEvent("013-010", "", &[]string{fmt.Sprint(reflect.TypeOf(data))})
		return errors.New(msg)
	}
	return nil
}

// Get a DO from cache using a cacheable struct of the same type we want to retrieve.
func (docache *doCache) GetFromCacheable(cacheable MakeCacheable) (interface{}, error) {
	return docache.get(cacheable.GetPK(), cacheable.GetName(), cacheable)
}

// Get a DO from cache using a cacheable struct of the same type we want to retrieve.
func (docache *doCache) MGetFromCacheable(cacheables []MakeCacheable) (*[]interface{}, error) {
	skeys := make([]string, len(cacheables))
	for i, _ := range cacheables {
		skeys[i] = cacheables[i].GetPK()
	}
	return docache.mget(skeys, cacheables[0].GetName(), cacheables[0])
}

// Get a DO from cache specifing its key and type
func (docache *doCache) GetFromKey(key string, dataType interface{}) (interface{}, error) {
	if cacheable, ok := dataType.(MakeCacheable); ok {
		return docache.get(key, cacheable.GetName(), dataType)
	} else {
		// The DO does not implements the MakeCacheable interface
		msg, _ := EVENT.NotifyEvent("013-004", "", &[]string{fmt.Sprint(reflect.TypeOf(dataType))})
		return nil, errors.New(msg)
	}
}

// Underlying method used to retrieve a DO from the cache.
func (docache *doCache) get(key string, name string, dataType interface{}) (interface{}, error) {

	if b_key, err := makeDaoPK(name, key); err == nil {
		if b, err2 := docache.cache.GetData(b_key); err2 == nil {
			if item, err3 := dataType.(MakeSerializable).FromBytes(b); err3 == nil {
				if (*item.(MakeCacheable).GetMetadata())[CACHEHELPER.METADATA_STATUS] == CACHEHELPER.CACHEOBJECTSTATUS_DELETED {
					// The current CO has been deleted.
					msg, _ := EVENT.NotifyEvent("013-012", "", &[]string{fmt.Sprint(item.(MakeCacheable).GetName()), item.(MakeCacheable).(MakeCacheable).GetPK()})
					return nil, errors.New(msg)
				} else {
					return item, nil
				}
			} else {
				// It has not been possible to unmarshal the retrieved data
				msg, _ := EVENT.NotifyEvent("013-007", "", &[]string{err3.Error()})
				LOG.Error.Println(msg)
				return nil, errors.New(msg)
			}
		} else {
			// It has not been possible retrieve the data from cache
			msg, _ := EVENT.NotifyEvent("013-006", "", &[]string{key, name, err2.Error()})
			return nil, errors.New(msg)
		}
	} else {
		// It has not been possible to marshal the key
		msg, _ := EVENT.NotifyEvent("013-005", "", &[]string{err.Error()})
		LOG.Error.Println(msg)
		return nil, errors.New(msg)
	}
	return nil, nil
}

func (docache *doCache) mget(skeys []string, name string, dataType interface{}) (*[]interface{}, error) {

	b_keys := make([][]byte, len(skeys))
	for i, skey := range skeys {
		if b_key, err := makeDaoPK(name, skey); err == nil {
			b_keys[i] = *b_key
		} else {
			// It has not been possible retrieve the data from cache
			msg, _ := EVENT.NotifyEvent("013-006", "", &[]string{skey, name, err.Error()})
			return nil, errors.New(msg)
		}
	}

	if b_values, err2 := docache.cache.MGetData(&b_keys); err2 == nil {
		//values := make([]interface{}, len(*b_values))
		var values []interface{}
		for _, b_value := range *b_values {
			if len(b_value) > 0 {
				if item, err3 := dataType.(MakeSerializable).FromBytes(&b_value); err3 == nil {
					//values[_i] = item
					values = append(values, item)
				} else {
					// It has not been possible to unmarshal the retrieved data
					msg, _ := EVENT.NotifyEvent("013-007", "", &[]string{err3.Error()})
					LOG.Error.Println(msg)
					return nil, errors.New(msg)
				}
			}
		}
		return &values, nil
	} else {
		// It has not been possible retrieve the data from cache
		msg, _ := EVENT.NotifyEvent("013-006", "", &[]string{fmt.Sprint(skeys), name, err2.Error()})
		return nil, errors.New(msg)
	}
	return nil, nil
}

// The cachable to be deleted will not be deleted from the cache.
// Instead, the cacheable will be marked as a deleted item.
func (docache *doCache) RemoveFromCacheable(cacheable MakeCacheable) (bool, error) {

	//println("*jas* Remove from cacheable 1: ", cacheable.GetName(), cacheable.GetPK(), fmt.Sprint(cacheable.GetMetadata()))

	// Check for record metadata
	if err := docache.checkForCacheableMetadataStatus(cacheable); err != nil {
		return false, err
	} else {

		// The cacheable to be deleted, exist into cache?
		if to_be_deleted, err := docache.GetFromCacheable(cacheable); err != nil {
			// The cacheable can't be retrieved from cache. Does it exist into cache?
			// It can't be deleted.
			return false, err
		} else {
			// Mark the retrieved cacheable as a deleted item.
			_cacheable := to_be_deleted.(MakeCacheable)
			new_version := (*cacheable.GetMetadata())[CACHEHELPER.METADATA_VERSION].(uint64) + 1
			_cacheable.SetMetadataField(CACHEHELPER.METADATA_STATUS, CACHEHELPER.CACHEOBJECTSTATUS_DELETED)
			_cacheable.SetMetadataField(CACHEHELPER.METADATA_VERSION, new_version)

			// Set the new cache status for the item.
			if err := docache.Put(_cacheable); err != nil {
				return false, err
			} else {
				return true, nil
			}
		}
	}

	//return docache.RemoveFromKey(cacheable.GetPK(), cacheable.GetName())
}

//
// -------------------------------- Check for metadata values --------------------------
//

// This checks for the cacheable
func (docache *doCache) checkForCacheableMetadataStatus(cacheable MakeCacheable) error {
	// Check for Cacheable state timestamp
	if current, err6 := docache.GetFromCacheable(cacheable); err6 != nil {
		// The cacheable does not exist into cache yet.
		return nil
	} else {

		// 26-03-2015 : A deleted entity/PK may be set again.
		//		if (*current.(MakeCacheable).GetMetadata())[METADATA_STATUS] == CACHEOBJECTSTATUS_DELETED {
		//			// The current CO has been deleted. It can't be updated
		//			msg, _ := EVENT.NotifyEvent("013-012", "", &[]string{fmt.Sprint(cacheable.GetName()), current.(MakeCacheable).GetPK()})
		//			return errors.New(msg)
		//
		//		} else
		if !docache.checkForCachableTS(cacheable, current.(MakeCacheable)) {
			// The timestamp of the current co into cache is later than the new co state.
			// The put will not be done.
			msg, _ := EVENT.NotifyEvent("013-011", "", &[]string{fmt.Sprint(cacheable.GetName()), fmt.Sprint((*cacheable.GetMetadata())[CACHEHELPER.METADATA_VERSION]), fmt.Sprint((*current.(MakeCacheable).GetMetadata())[CACHEHELPER.METADATA_VERSION])})
			return errors.New(msg)
		} else {
			return nil
		}
	}
}

// This method returns false if the timestamp of the metadata cached co is newer than the new co
func (docache *doCache) checkForCachableTS(new_state, current_state MakeCacheable) bool {
	//    v1 := (*current_state.GetMetadata())[CACHEHELPER.METADATA_VERSION].(uint64)
	//    v2 := (*new_state.GetMetadata())[CACHEHELPER.METADATA_VERSION].(uint64)
	//    println("*jas* cache. version ", v1, "<", v2, v1 < v2)
	return (*current_state.GetMetadata())[CACHEHELPER.METADATA_VERSION].(uint64) < (*new_state.GetMetadata())[CACHEHELPER.METADATA_VERSION].(uint64)
}
