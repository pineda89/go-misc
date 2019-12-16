package lib_idx_value

import (
	CACHEHELPER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_cache/lib_gc_cache_helpers"
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	PROTO "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_protobuf_protocols/capnp/lib_gc_capnp_common"
	PROTO_HELPER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_protobuf_protocols/capnp/lib_gc_capnp_helpers"

	C "github.com/glycerine/go-capnproto"

	"bytes"
	"errors"
	"fmt"
)

//---- DOIndexValue MakeCacheable implementation-------
func (doIdxValue *DOIndexValue) GetName() string {
	return "DOIndexValue::"
}

func (doIdxValue *DOIndexValue) GetPK() string {
	var buff bytes.Buffer
	var v interface{}
	for i, item := range doIdxValue.Items().ToArray() {
		if i > 0 {
			buff.WriteString("::")
		}
		v, _ = HGetValueFromDOIndexValueItem(&item)
		buff.WriteString(fmt.Sprint(v))
	}
	return buff.String()
}

func (doIdxValue *DOIndexValue) Clone() interface{} {
	return nil
}

//---- DOIndexValue MakeSerializable implementation-------
func (doIdxValue *DOIndexValue) ToBytes() (*[]byte, error) {
	_b, err := PROTO_HELPER.HSerializeAsCapNProto(doIdxValue.Segment)
	return _b, err
}

func (doIdxValue *DOIndexValue) FromBytes(_b *[]byte) (interface{}, error) {
	var buf = bytes.Buffer{}

	buf.Write(*_b)
	if capnp_seg, err := C.ReadFromStream(&buf, nil); err == nil {
		item := ReadRootDOIndexValue(capnp_seg)
		return &item, nil
	} else {
		return nil, err
	}
}

func (doIdxValue *DOIndexValue) GetMetadata() *map[string]interface{} {
	m := make(map[string]interface{})

	metadata := doIdxValue.Metadata()
	m[CACHEHELPER.METADATA_VERSION] = metadata.Version()
	m[CACHEHELPER.METADATA_STATUS] = metadata.Status()

	return &m
}

func (doIdxValue *DOIndexValue) SetMetadataField(field string, value interface{}) error {

	metadata := PROTO.NewRootMetadata(C.NewBuffer(nil))
	version := doIdxValue.Metadata().Version()
	status := doIdxValue.Metadata().Status()

	switch {
	case field == CACHEHELPER.METADATA_VERSION:
		metadata.SetVersion(value.(uint64))
		metadata.SetStatus(status)
	case field == CACHEHELPER.METADATA_STATUS:
		metadata.SetStatus(value.(int16))
		metadata.SetVersion(version)
	default:
		msg, _ := EVENT.NotifyEvent("014-014", "", &[]string{field})
		return errors.New(msg)
	}
	doIdxValue.SetMetadata(metadata)

	return nil
}

func HGetValueFromDOIndexValueItem(item *DOIndexValueItem) (interface{}, error) {

	switch item.ValueType() {
	case DOINDEXVALUEITEMATTRIBUTETYPE_TEXT:
		v, _ := item.GetValue_string()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT8:
		v, _ := item.GetValue_int8()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT16:
		v, _ := item.GetValue_int16()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT32:
		v, _ := item.GetValue_int32()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT64:
		v, _ := item.GetValue_int64()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT32:
		v, _ := item.GetValue_float32()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT64:
		v, _ := item.GetValue_float64()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_BOOL:
		v, _ := item.GetValue_bool()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT8:
		v, _ := item.GetValue_uint8()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT16:
		v, _ := item.GetValue_uint16()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT32:
		v, _ := item.GetValue_uint32()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT64:
		v, _ := item.GetValue_uint64()
		return v, nil
	case DOINDEXVALUEITEMATTRIBUTETYPE_DATA:
		v, _ := item.GetValue_data()
		return v, nil
	default:
		msg, _ := EVENT.NotifyEvent("014-006", "", nil)
		err := errors.New(msg)
		return nil, err
	}
}

// ----------- INDEX VALUE --------------------------------------------------
type DOIndexValuator interface {
	GetValue_uint8() (uint8, error)
	GetValue_uint16() (uint16, error)
	GetValue_uint32() (uint32, error)
	GetValue_uint64() (uint64, error)
	GetValue_int8() (int8, error)
	GetValue_int16() (int16, error)
	GetValue_int32() (int32, error)
	GetValue_int64() (int64, error)
	GetValue_float32() (float32, error)
	GetValue_float64() (float64, error)
	GetValue_data() (*[]byte, error)
	GetValue_bool() (bool, error)
}

func (doIdxValueItem *DOIndexValueItem) GetValue_string() (string, error) {
	return doIdxValueItem.Value().VText(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_uint8() (uint8, error) {
	return doIdxValueItem.Value().VUint8(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_uint16() (uint16, error) {
	return doIdxValueItem.Value().VUint16(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_uint32() (uint32, error) {
	return doIdxValueItem.Value().VUint32(), nil
}
func (doIdxValueItem *DOIndexValueItem) GetValue_uint64() (uint64, error) {
	return doIdxValueItem.Value().VUint64(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_int8() (int8, error) {
	return doIdxValueItem.Value().VInt8(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_int16() (int16, error) {
	return doIdxValueItem.Value().VInt16(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_int32() (int32, error) {
	return doIdxValueItem.Value().VInt32(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_int64() (int64, error) {
	return doIdxValueItem.Value().VInt64(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_float32() (float32, error) {
	return doIdxValueItem.Value().VFloat32(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_float64() (float64, error) {
	return doIdxValueItem.Value().VFloat64(), nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_data() (*[]byte, error) {
	v := doIdxValueItem.Value().VData()
	return &v, nil
}

func (doIdxValueItem *DOIndexValueItem) GetValue_bool() (bool, error) {
	return doIdxValueItem.Value().VBool(), nil
}
