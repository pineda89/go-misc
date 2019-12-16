package lib_gc_idx

import (
	IDX_VALUE "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_idx/lib_gc_idx_value/protobuf_protocols/capnp"
	UTIL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_util"
)

// Retrieves the DOIndexesManager
func HGetDOIndexesManager() doIndexesManager {
	return DOIndexesManager
}

// Given an index, this method checks for an specific value as a part of the slice of values of a given index key.
// If the index key does not exist into index, the method return an error.
// If the index key does exist into index, but the specific value does not exist, the method returns false.
// Remember that a value is a slice
func HGetValueAsSlice(indexer DOIndexer, index_key *[]string, valueKey ...interface{}) (*[]IDX_VALUE.DOIndexValueItem, bool, error) {
	if _map, err := indexer.GetAsMap(index_key); err == nil {
		value, ok := HGetMapValueAsSlice(_map, &valueKey)
		if ok {
			items := value.Items().ToArray()
			return &items, ok, nil
		} else {
			return nil, ok, nil
		}
	} else {
		return nil, false, err
	}
}

// Given an index, a domain object with the index key files filled, and a specific value,
// this method checks if into the index exist an entry with the given index key. If it exist,
// then the method checks if between its map of values it is the given value.
//
// If the index key does not exist into the index, the method returns an error.
// If the index exist, but the value is not between their values, it returns false
func HGetValueAsSliceFromDo(indexer DOIndexer, do MakeIndexable, valueKey ...interface{}) (*[]IDX_VALUE.DOIndexValueItem, bool, error) {

	attributesMap, err := do.GetAsMap()
	if err == nil {
		keyFields := indexer.GetKeyField()
		s_key := make([]string, len(*keyFields))
		for i, f := range *keyFields {
			s_key[i] = UTIL.GetAsString((*attributesMap)[f])
		}
		if _map, err := indexer.GetAsMap(&s_key); err == nil {
			value, ok := HGetMapValueAsSlice(_map, &valueKey)
			if ok {
				list := value.Items().ToArray()
				return &list, ok, nil
			} else {
				return nil, ok, nil
			}
		} else {
			return nil, false, err
		}
	} else {
		return nil, false, err
	}

}

// Given a map of values, linked to an index key, this method  retrieves its slice of values.
// Remember that a value is a slice.
// For exemple, this methos can be called after calling the HGetMapOfValuesForAnIndexKey.
func HGetMapValueAsSlice(_map *map[string]*IDX_VALUE.DOIndexValue, keyValues *[]interface{}) (*IDX_VALUE.DOIndexValue, bool) {
	key := UTIL.GetAsString(*keyValues)
	z, ok := (*_map)[key]
	if ok {
		return z, ok
	} else {
		return nil, ok
	}
}
