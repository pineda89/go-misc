package lib_gc_idx

import (
	CACHE "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_cache"
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	EVENT_PUBLISHER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event_publisher"
	IDX_VALUE "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_idx/lib_gc_idx_value/protobuf_protocols/capnp"
	LOG "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"

	UTIL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_util"

	C "github.com/glycerine/go-capnproto"

	"bytes"
	"errors"
	"fmt"
	"sync"
)

func init() {

	// Create the indexes manager
	DOIndexesManager = &doIndexesContainer{}

	// Create the default entry maker functions
	default_keyFromDoMaker := func(indexHeaderName string, keyFields *[]string, do MakeIndexable) (string, error) {
		attributesMap, err := do.GetAsMap()
		if err == nil {
			var buff bytes.Buffer
			buff.WriteString(indexHeaderName)
			buff.WriteString("::")
			for i, f := range *keyFields {
				if i > 0 {
					buff.WriteString(":")
				}
				buff.WriteString(UTIL.GetAsString((*attributesMap)[f]))
			}
			return buff.String(), nil
		} else {
			return "", err
		}
	}

	default_keyMaker := func(indexHeaderName string, keyFields *[]string) (string, error) {
		var buff bytes.Buffer
		buff.WriteString(indexHeaderName)
		buff.WriteString("::")
		for i, f := range *keyFields {
			if i > 0 {
				buff.WriteString(":")
			}
			buff.WriteString(UTIL.GetAsString(f))
		}
		return buff.String(), nil
	}

	default_ValuesFromDoMaker := func(valueFields *[]string, do MakeIndexable) (*[]interface{}, error) {
		attributesMap, err := do.(MakeIndexable).GetAsMap()
		if err == nil {
			v := make([]interface{}, len(*valueFields))
			for i, f := range *valueFields {
				if _v, ok := (*attributesMap)[f]; !ok {
					return nil, errors.New(fmt.Sprint("field:", f, " is nil "))
				} else {
					v[i] = _v
				}
			}
			return &v, nil
		} else {
			return nil, err
		}
	}

	// Create the entry maker
	DefaultEntryMaker = &doIndexEntry{
		default_keyFromDoMaker,
		default_keyMaker,
		default_ValuesFromDoMaker,
	}

	// Create the index factory
	DOIndexFactory = &doIndexFactory{}

	LOG.Info.Println("lib_idx package initialiated. ")
}

// Indexes manager
var DOIndexesManager doIndexesManager

// Default index entries maker
var DefaultEntryMaker DOIndexEntryMaker

// Index factory
var DOIndexFactory *doIndexFactory

// Mutexes
var getAsSliceMutex *sync.Mutex = &sync.Mutex{}
var getAsMap *sync.Mutex = &sync.Mutex{}
var removeKey *sync.Mutex = &sync.Mutex{}

// ----------- INDEX DEFINITION -----------------------------------------------
type IdxOrder int8

const (
	IDX_ORDER_ASC = iota
	IDX_ORDER_DESC
)

type DOIndexHeader struct {
	name       string
	keyfield   []string
	valuefield []string
	unique     bool
	order      IdxOrder
}

// ----------- INDEX FACTORY --------------------------------------------------
type doIndexFactory struct {
}
type DOIndexMaker interface {
	CreateDOIndexer(name string, keyField []string, valueField []string, unique bool, order IdxOrder) (DOIndexer, error)
}

func (doIdxFactory *doIndexFactory) CreateDOIndexer(name string, keyField []string, valueField []string, unique bool, order IdxOrder) (DOIndexer, error) {
	indexer := DOIndexHeader{name: name, keyfield: keyField, valuefield: valueField, unique: unique, order: order}
	return &indexer, nil
}

// ------- INDEXES CONTAINER --------------------------------------------------
type doIndexesManager interface {
	PutDOIndex(indexer DOIndexer) error
	GetDOIndex(name string, keyField []string, valueField []string, unique bool, order IdxOrder) (DOIndexer, error)
	GetDOIndexFromDO(indexer DOIndexer) (DOIndexer, error)
	RemoveDOIndex(indexer DOIndexer) (bool, error)
}

type doIndexesContainer struct {
}

/*
	doIndexesManager implementation -------------------------------------------
*/

// Add a new index (DOIndexHeader) to the cache
func (doIdxC *doIndexesContainer) PutDOIndex(indexer DOIndexer) error {
	//return CACHE.DOCache.Put(indexer)
	return errors.New("method not implemented !!!")
}

// Retrieve an index (DOIndexer)
func (doIdxC *doIndexesContainer) GetDOIndex(name string, keyField []string, valueField []string, unique bool, order IdxOrder) (DOIndexer, error) {
	idx, _ := DOIndexFactory.CreateDOIndexer(name, keyField, valueField, unique, order)
	return idx, nil
}

// Retrieve an existing index (DOIndexHeader) from the cache
func (doIdxC *doIndexesContainer) GetDOIndexFromDO(indexer DOIndexer) (DOIndexer, error) {
	return nil, errors.New("*doIndexesContainer.GetDOIndexFromDO method not implemented !!!")
}

// Remove an existing index (DOIndexHeader) from the cache
func (doIdxC *doIndexesContainer) RemoveDOIndex(indexer DOIndexer) (bool, error) {
	return false, errors.New("*doIndexesContainer.RemoveDOIndex method not implemented !!!")
}

//-------. INDEXABLE INTERFACE -----------------------------------------------
type MakeIndexable interface {
	GetAsMap() (*map[string]interface{}, error)
}

//-------. INDEX HEADER -------------------------------------------------------
//
// A DOIndexHeader instance performs a specific index.
//	Each key of this index, will be an DOIndexKey instance
//
type DOIndexer interface {

	// Make it cacheable
	GetName() string
	GetPK() string

	// Attribute Getters
	GetOrdered() IdxOrder
	GetUnique() bool
	GetKeyField() *[]string
	GetValueField() *[]string

	// Entry maker getter
	GetEntryMaker() DOIndexEntryMaker

	// Entry A/D operations
	AddEntry(do MakeIndexable) error
	RemoveIndexKeyEntryFromDo(do MakeIndexable) error
	RemoveIndexKeyEntry(key *[]string, entryValue ...interface{}) error
	RemoveIndexKey(key *[]string) error

	// Read operations
	GetAsSlice(key *[]string) (*[]IDX_VALUE.DOIndexValue, error)
	GetAsSliceFromDo(do MakeIndexable) (*[]IDX_VALUE.DOIndexValue, error)
	GetAsMap(key *[]string) (*map[string]*IDX_VALUE.DOIndexValue, error)
	GetAsMapFromDo(do MakeIndexable) (*map[string]*IDX_VALUE.DOIndexValue, error)
}

// ---------- MakeCacheable implementation -----------------------------------

//---- DOIndexHeader MakeCacheable implementation-------
func (doIdxHeader *DOIndexHeader) GetName() string {
	var buff bytes.Buffer
	buff.WriteString(doIdxHeader.name)
	return buff.String()
}

func (doIdxHeader *DOIndexHeader) GetPK() string {
	return doIdxHeader.GetName()
}

/*
func (doIdxHeader *DOIndexHeader) Clone() interface{} {
	return nil
}

func (doIdxHeader *DOIndexHeader) GetMetadata() *map[string]interface{} {
	m := make(map[string]interface{})

	metadata := doIdxHeader.Metadata()
	m[PROTO_HELPER..METADATA_VERSION] = metadata.Version()
	m[PROTO_HELPER..METADATA_STATUS] = metadata.Status()

	return &m
}

func (doIdxHeader *DOIndexHeader) SetMetadataField(field string, value interface{}) error {

	metadata := CAPNP_COMMON.NewRootMetadata(C.NewBuffer(nil))
	version := doIdxHeader.Metadata().Version()
	status := doIdxHeader.Metadata().Status()

	switch {
	case field == PROTO_HELPER..METADATA_VERSION:
		metadata.SetVersion(value.(uint64))
		metadata.SetStatus(status)
	case field == PROTO_HELPER..METADATA_STATUS:
		metadata.SetStatus(value.(CAPNP_COMMON.CacheObjectStatus))
		metadata.SetVersion(version)
	default:
		msg, _ := EVENT.NotifyEvent("028-001", "", &[]string{field})
		return errors.New(msg)
	}
	doIdxHeader.SetMetadata(metadata)

	return nil
}

// ---------- MakeSerializable implementation -------------------------------

//---- DOIndexHeader MakeSerializable implementation-------
func (doIdxHeader *DOIndexHeader) ToBytes() (*[]byte, error) {
	_b, err := PROTO_HELPER..HSerializeAsCapNProto(doIdxHeader.Segment)
	return _b, err
}

func (doIdxHeader *DOIndexHeader) FromBytes(_b *[]byte) (interface{}, error) {
	var buf = bytes.Buffer{}

	buf.Write(*_b)
	if capnp_seg, err := C.ReadFromStream(&buf, nil); err == nil {
		item := ReadRootDOIndexHeader(capnp_seg)
		return &item, nil
	} else {
		return nil, err
	}
}
*/

// ----------- doIndexer implementation --------------------------------------

func (doIdxHeader *DOIndexHeader) GetOrdered() IdxOrder {
	return doIdxHeader.order
}

func (doIdxHeader *DOIndexHeader) GetUnique() bool {
	return doIdxHeader.unique
}

func (doIdxHeader *DOIndexHeader) GetKeyField() *[]string {
	keyField := doIdxHeader.keyfield
	return &keyField
}

func (doIdxHeader *DOIndexHeader) GetValueField() *[]string {
	valueField := doIdxHeader.valuefield
	return &valueField
}

func (doIdxHeader *DOIndexHeader) GetEntryMaker() DOIndexEntryMaker {
	return DefaultEntryMaker
}

/*
	Add a new entry to the index
*/

// ---------------- Index key locking ......?

func (doIdxHeader *DOIndexHeader) AddEntry(do MakeIndexable) error {
	//addEntryMutex.Lock()
	//defer addEntryMutex.Unlock()

	var _err error

	// Take the key
	var entryMaker DOIndexEntryMaker = doIdxHeader.GetEntryMaker()
	var key_field *[]string = doIdxHeader.GetKeyField()

	if key, err := entryMaker.GetKeyFromDo(doIdxHeader.GetName(), key_field, do); err == nil {
		if value, err2 := doIdxHeader.GetEntryMaker().GetValueFromDo(doIdxHeader.GetValueField(), do); err2 == nil {

			doIndexValue := IDX_VALUE.NewRootDOIndexValue(C.NewBuffer(nil))

			// Make the list of value items
			doIndexValueItem_list, err22 := getDOIndexValueItemList(value, C.NewBuffer(nil))
			if err22 != nil {
				return err
			}

			// Set the list of value items to the index key value
			doIndexValue.SetItems(*doIndexValueItem_list)

			// Store the index key
			if cache_adapter, err := CACHE.GetCacheAdapter(); err == nil {
				if err := cache_adapter.AddValueToIndexKey(&key, &doIndexValue, doIdxHeader.GetName()); err != nil {
					_err = err
				}
			} else {
				_err = err
			}
		} else {
			_err = err2
		}
	} else {
		_err = err
	}

	return _err
}

/*
	Takes a slice entry indexed by the "key" parameter.
*/
func (doIdxHeader *DOIndexHeader) GetAsSlice(key *[]string) (*[]IDX_VALUE.DOIndexValue, error) {
	//getAsSliceMutex.Lock()
	//defer getAsSliceMutex.Unlock()

	if cache_adapter, err := CACHE.GetCacheAdapter(); err != nil {

		return nil, err

	} else {

		// Take the key
		_key, _ := doIdxHeader.GetEntryMaker().GetKey(doIdxHeader.GetName(), key)

		// Take the index key values
		var dataType IDX_VALUE.DOIndexValue
		if values, err := cache_adapter.GetValuesFromIndexKey(&_key, &dataType, doIdxHeader.GetName()); err != nil {
			msg, _ := EVENT.NotifyEvent("014-013", "", &[]string{err.Error()})
			return nil, errors.New(msg)
		} else {
			return values.(*[]IDX_VALUE.DOIndexValue), nil
		}
	}
}

/*
	Takes the index entry for the "key" parameter, as a map.
*/
func (doIdxHeader *DOIndexHeader) GetAsMap(key *[]string) (*map[string]*IDX_VALUE.DOIndexValue, error) {
	//t0 := time.Now()
	m, err := doIdxHeader.getAsMap(key)
	//LOG.Warning.Println("GetAsMap in ", time.Now().Sub(t0))
	return m, err
}

/*
	Takes the index entry for the DO's "key" parameter, as a slice.
*/
func (doIdxHeader *DOIndexHeader) GetAsSliceFromDo(do MakeIndexable) (*[]IDX_VALUE.DOIndexValue, error) {
	//t0 := time.Now()
	attributesMap, err := do.GetAsMap()
	if err == nil {
		s := make([]string, len(*doIdxHeader.GetKeyField()))
		for i, f := range *doIdxHeader.GetKeyField() {
			s[i] = UTIL.GetAsString((*attributesMap)[f])
		}

		m, err := doIdxHeader.GetAsSlice(&s)
		//LOG.Warning.Println("GetAsSliceFromDo in ", time.Now().Sub(t0))
		return m, err

	} else {
		return nil, err
	}

}

/*
	Takes the index entry for the DO's "key" parameter, as a map.
*/
func (doIdxHeader *DOIndexHeader) GetAsMapFromDo(do MakeIndexable) (*map[string]*IDX_VALUE.DOIndexValue, error) {
	//t0 := time.Now()
	attributesMap, err := do.GetAsMap()
	if err == nil {
		s := make([]string, len(*doIdxHeader.GetKeyField()))
		for i, f := range *doIdxHeader.GetKeyField() {
			s[i] = UTIL.GetAsString((*attributesMap)[f])
		}

		m, err := doIdxHeader.GetAsMap(&s)
		//LOG.Warning.Println("GetAsMapFromDo in ", time.Now().Sub(t0))
		return m, err

	} else {
		return nil, err
	}

}

func (doIdxHeader *DOIndexHeader) getAsMap(key *[]string) (*map[string]*IDX_VALUE.DOIndexValue, error) {

	if values, err := doIdxHeader.GetAsSlice(key); err == nil {

		m_values := make(map[string]*IDX_VALUE.DOIndexValue)
		for _, v := range *values {
			value := v
			s_value := make([]interface{}, v.Items().Len())
			for z, item := range v.Items().ToArray() {
				if item_value, err := IDX_VALUE.HGetValueFromDOIndexValueItem(&item); err == nil {
					s_value[z] = item_value
				} else {
					return nil, err
				}
			}
			m_values[UTIL.GetAsString(s_value)] = &value
		}

		return &m_values, nil
	} else {
		msg, _ := EVENT.NotifyEvent("014-003", "", &[]string{doIdxHeader.GetName(), fmt.Sprint(key)})
		EVENT_PUBLISHER.EventPublisherChannel <- msg
		err := errors.New(msg)
		return nil, err
	}
}

/*
	Remove the index entry filled into the "do" parameter.
*/
func (doIdxHeader *DOIndexHeader) RemoveIndexKeyEntryFromDo(do MakeIndexable) error {

	var entryMaker DOIndexEntryMaker = doIdxHeader.GetEntryMaker()
	var key_field *[]string = doIdxHeader.GetKeyField()

	if key, err := entryMaker.GetKeyFromDo(doIdxHeader.GetName(), key_field, do); err == nil {
		if value, err := doIdxHeader.GetEntryMaker().GetValueFromDo(doIdxHeader.GetValueField(), do); err == nil {

			doIndexValue := IDX_VALUE.NewRootDOIndexValue(C.NewBuffer(nil))

			// Make the list of value items
			if doIndexValueItem_list, err22 := getDOIndexValueItemList(value, C.NewBuffer(nil)); err22 != nil {
				return err
			} else {

				// Set the list of value items to the index key value
				doIndexValue.SetItems(*doIndexValueItem_list)

				// Remove the entry from index key
				if cache_adapter, err := CACHE.GetCacheAdapter(); err == nil {
					_n := doIdxHeader.GetName()
					if err := cache_adapter.RemoveValueFromIndexKey(&key, &doIndexValue, &_n); err != nil {
						return err
					} else {
						return nil
					}
				} else {
					return err
				}
			}

		} else {
			return err
		}
	} else {
		return err
	}

}

/*
	Removes an index entry

*/
func (doIdxHeader *DOIndexHeader) RemoveIndexKeyEntry(key *[]string, entryValue ...interface{}) error {

	// take the index key
	if indexkey, err := doIdxHeader.GetEntryMaker().GetKey(doIdxHeader.GetName(), key); err == nil {

		// Create the doIndexValue
		doIndexValue := IDX_VALUE.NewDOIndexValue(C.NewBuffer(nil))

		// Make the list of value items
		doIndexValueItem_list, err22 := getDOIndexValueItemList(&entryValue, C.NewBuffer(nil))
		if err22 != nil {
			return err
		}

		// Set the list of value items to the index key value
		doIndexValue.SetItems(*doIndexValueItem_list)

		// Removing.
		_n := doIdxHeader.GetName()
		return CACHE.DOCache.RemoveValueFromIndexKey(&indexkey, &doIndexValue, &_n)

	} else {
		return err
	}
	return nil
}

/*
	Remove an index header key
*/
func (doIdxHeader *DOIndexHeader) RemoveIndexKey(key *[]string) error {
	removeKey.Lock()
	defer removeKey.Unlock()

	//s_key, _ := doIdxHeader.GetEntryMaker().GetKey(doIdxHeader.GetName(), key)
	//if err := CACHE.DOCache.RemoveIndexKey(&s_key); err != nil {
	//	return err
	//} else {
	//	return nil
	//}

	return errors.New("method not implemented !!!")
}

// ----------- INDEX ENTRY MAKER ----------------------------------------------
type doIndexEntry struct {
	KeyFromDoMaker   func(indexHeaderName string, keyFields *[]string, do MakeIndexable) (string, error)
	KeyMaker         func(indexHeaderName string, keyFields *[]string) (string, error)
	ValueFromDoMaker func(valueFields *[]string, do MakeIndexable) (*[]interface{}, error)
}
type DOIndexEntryMaker interface {
	GetKeyFromDo(indexHeaderName string, key *[]string, do MakeIndexable) (string, error)
	GetKey(indexHeaderName string, key *[]string) (string, error)
	GetValueFromDo(value *[]string, do MakeIndexable) (*[]interface{}, error)
}

/*
	doIndexEntryMaker implementation -------------------------------------------
*/
func (doIdxEntry *doIndexEntry) GetKeyFromDo(indexHeaderName string, keyFields *[]string, do MakeIndexable) (string, error) {
	return doIdxEntry.KeyFromDoMaker(indexHeaderName, keyFields, do)
}

func (doIdxEntry *doIndexEntry) GetKey(indexHeaderName string, keyFields *[]string) (string, error) {
	return doIdxEntry.KeyMaker(indexHeaderName, keyFields)
}

func (doIdxEntry *doIndexEntry) GetValueFromDo(valueFields *[]string, do MakeIndexable) (*[]interface{}, error) {
	return doIdxEntry.ValueFromDoMaker(valueFields, do)
}

// ----------------------------------------------------------------------

func getDOIndexValueItemList(value *[]interface{}, capnp_segment *C.Segment) (*IDX_VALUE.DOIndexValueItem_List, error) {

	// Make the list of value items
	doIndexValueItem_list := IDX_VALUE.NewDOIndexValueItemList(C.NewBuffer(nil), len(*value))
	pdoIndexValueItem_list := C.PointerList(doIndexValueItem_list)

	for i, v := range *value {
		value_item := IDX_VALUE.NewDOIndexValueItem(capnp_segment)

		if _v, ok := v.(bool); ok {
			value_item.Value().SetVBool(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_BOOL)
		} else if _v, ok := v.(string); ok {
			value_item.Value().SetVText(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_TEXT)
		} else if _v, ok := v.(int8); ok {
			value_item.Value().SetVInt8(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_INT8)
		} else if _v, ok := v.(int16); ok {
			value_item.Value().SetVInt16(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_INT16)
		} else if _v, ok := v.(int32); ok {
			value_item.Value().SetVInt32(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_INT32)
		} else if _v, ok := v.(int64); ok {
			value_item.Value().SetVInt64(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_INT64)
		} else if _v, ok := v.(uint8); ok {
			value_item.Value().SetVUint8(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_UINT8)
		} else if _v, ok := v.(uint16); ok {
			value_item.Value().SetVUint16(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_UINT16)
		} else if _v, ok := v.(uint32); ok {
			value_item.Value().SetVUint32(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_UINT32)
		} else if _v, ok := v.(uint64); ok {
			value_item.Value().SetVUint64(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_UINT64)
		} else if _v, ok := v.(float32); ok {
			value_item.Value().SetVFloat32(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT32)
		} else if _v, ok := v.(float64); ok {
			value_item.Value().SetVFloat64(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT64)
		} else if _v, ok := v.([]byte); ok {
			value_item.Value().SetVData(_v)
			value_item.SetValueType(IDX_VALUE.DOINDEXVALUEITEMATTRIBUTETYPE_DATA)
		} else {
			msg, _ := EVENT.NotifyEvent("014-010", "", &[]string{fmt.Sprint(v)})
			EVENT_PUBLISHER.EventPublisherChannel <- msg
			//*jas*
			panic(msg)
			err := errors.New(msg)
			return nil, err
		}

		pdoIndexValueItem_list.Set(i, C.Object(value_item))
	}
	return &doIndexValueItem_list, nil
}
