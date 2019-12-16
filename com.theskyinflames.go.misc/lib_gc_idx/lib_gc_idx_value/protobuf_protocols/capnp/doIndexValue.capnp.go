package lib_idx_value

// AUTO GENERATED - DO NOT EDIT

import (
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_protobuf_protocols/capnp/lib_gc_capnp_common"
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
	"math"
)

type DOIndexValueItem C.Struct
type DOIndexValueItemValue DOIndexValueItem
type DOIndexValueItemValue_Which uint16

const (
	DOINDEXVALUEITEMVALUE_VDATA    DOIndexValueItemValue_Which = 0
	DOINDEXVALUEITEMVALUE_VBOOL    DOIndexValueItemValue_Which = 1
	DOINDEXVALUEITEMVALUE_VINT8    DOIndexValueItemValue_Which = 2
	DOINDEXVALUEITEMVALUE_VINT16   DOIndexValueItemValue_Which = 3
	DOINDEXVALUEITEMVALUE_VINT32   DOIndexValueItemValue_Which = 4
	DOINDEXVALUEITEMVALUE_VINT64   DOIndexValueItemValue_Which = 5
	DOINDEXVALUEITEMVALUE_VUINT8   DOIndexValueItemValue_Which = 6
	DOINDEXVALUEITEMVALUE_VUINT16  DOIndexValueItemValue_Which = 7
	DOINDEXVALUEITEMVALUE_VUINT32  DOIndexValueItemValue_Which = 8
	DOINDEXVALUEITEMVALUE_VUINT64  DOIndexValueItemValue_Which = 9
	DOINDEXVALUEITEMVALUE_VFLOAT32 DOIndexValueItemValue_Which = 10
	DOINDEXVALUEITEMVALUE_VFLOAT64 DOIndexValueItemValue_Which = 11
	DOINDEXVALUEITEMVALUE_VTEXT    DOIndexValueItemValue_Which = 12
)

func NewDOIndexValueItem(s *C.Segment) DOIndexValueItem { return DOIndexValueItem(s.NewStruct(24, 2)) }
func NewRootDOIndexValueItem(s *C.Segment) DOIndexValueItem {
	return DOIndexValueItem(s.NewRootStruct(24, 2))
}
func AutoNewDOIndexValueItem(s *C.Segment) DOIndexValueItem {
	return DOIndexValueItem(s.NewStructAR(24, 2))
}
func ReadRootDOIndexValueItem(s *C.Segment) DOIndexValueItem {
	return DOIndexValueItem(s.Root(0).ToStruct())
}
func (s DOIndexValueItem) Metadata() lib_gc_capnp_common.Metadata {
	return lib_gc_capnp_common.Metadata(C.Struct(s).GetObject(0).ToStruct())
}
func (s DOIndexValueItem) SetMetadata(v lib_gc_capnp_common.Metadata) {
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s DOIndexValueItem) ValueType() DOIndexValueItemAttributeType {
	return DOIndexValueItemAttributeType(C.Struct(s).Get16(0))
}
func (s DOIndexValueItem) SetValueType(v DOIndexValueItemAttributeType) {
	C.Struct(s).Set16(0, uint16(v))
}
func (s DOIndexValueItem) Value() DOIndexValueItemValue { return DOIndexValueItemValue(s) }
func (s DOIndexValueItemValue) Which() DOIndexValueItemValue_Which {
	return DOIndexValueItemValue_Which(C.Struct(s).Get16(2))
}
func (s DOIndexValueItemValue) VData() []byte { return C.Struct(s).GetObject(1).ToData() }
func (s DOIndexValueItemValue) SetVData(v []byte) {
	C.Struct(s).Set16(2, 0)
	C.Struct(s).SetObject(1, s.Segment.NewData(v))
}
func (s DOIndexValueItemValue) VBool() bool     { return C.Struct(s).Get1(32) }
func (s DOIndexValueItemValue) SetVBool(v bool) { C.Struct(s).Set16(2, 1); C.Struct(s).Set1(32, v) }
func (s DOIndexValueItemValue) VInt8() int8     { return int8(C.Struct(s).Get8(4)) }
func (s DOIndexValueItemValue) SetVInt8(v int8) {
	C.Struct(s).Set16(2, 2)
	C.Struct(s).Set8(4, uint8(v))
}
func (s DOIndexValueItemValue) VInt16() int16 { return int16(C.Struct(s).Get16(4)) }
func (s DOIndexValueItemValue) SetVInt16(v int16) {
	C.Struct(s).Set16(2, 3)
	C.Struct(s).Set16(4, uint16(v))
}
func (s DOIndexValueItemValue) VInt32() int32 { return int32(C.Struct(s).Get32(4)) }
func (s DOIndexValueItemValue) SetVInt32(v int32) {
	C.Struct(s).Set16(2, 4)
	C.Struct(s).Set32(4, uint32(v))
}
func (s DOIndexValueItemValue) VInt64() int64 { return int64(C.Struct(s).Get64(8)) }
func (s DOIndexValueItemValue) SetVInt64(v int64) {
	C.Struct(s).Set16(2, 5)
	C.Struct(s).Set64(8, uint64(v))
}
func (s DOIndexValueItemValue) VUint8() uint8       { return C.Struct(s).Get8(4) }
func (s DOIndexValueItemValue) SetVUint8(v uint8)   { C.Struct(s).Set16(2, 6); C.Struct(s).Set8(4, v) }
func (s DOIndexValueItemValue) VUint16() uint16     { return C.Struct(s).Get16(4) }
func (s DOIndexValueItemValue) SetVUint16(v uint16) { C.Struct(s).Set16(2, 7); C.Struct(s).Set16(4, v) }
func (s DOIndexValueItemValue) VUint32() uint32     { return C.Struct(s).Get32(4) }
func (s DOIndexValueItemValue) SetVUint32(v uint32) { C.Struct(s).Set16(2, 8); C.Struct(s).Set32(4, v) }
func (s DOIndexValueItemValue) VUint64() uint64     { return C.Struct(s).Get64(8) }
func (s DOIndexValueItemValue) SetVUint64(v uint64) { C.Struct(s).Set16(2, 9); C.Struct(s).Set64(8, v) }
func (s DOIndexValueItemValue) VFloat32() float32   { return math.Float32frombits(C.Struct(s).Get32(4)) }
func (s DOIndexValueItemValue) SetVFloat32(v float32) {
	C.Struct(s).Set16(2, 10)
	C.Struct(s).Set32(4, math.Float32bits(v))
}
func (s DOIndexValueItemValue) VFloat64() float64 { return math.Float64frombits(C.Struct(s).Get64(8)) }
func (s DOIndexValueItemValue) SetVFloat64(v float64) {
	C.Struct(s).Set16(2, 11)
	C.Struct(s).Set64(8, math.Float64bits(v))
}
func (s DOIndexValueItemValue) VText() string { return C.Struct(s).GetObject(1).ToText() }
func (s DOIndexValueItemValue) SetVText(v string) {
	C.Struct(s).Set16(2, 12)
	C.Struct(s).SetObject(1, s.Segment.NewText(v))
}
func (s DOIndexValueItem) IsNull() bool     { return C.Struct(s).Get1(128) }
func (s DOIndexValueItem) SetIsNull(v bool) { C.Struct(s).Set1(128, v) }
func (s DOIndexValueItem) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"metadata\":")
	if err != nil {
		return err
	}
	{
		s := s.Metadata()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"valueType\":")
	if err != nil {
		return err
	}
	{
		s := s.ValueType()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"value\":")
	if err != nil {
		return err
	}
	{
		s := s.Value()
		err = b.WriteByte('{')
		if err != nil {
			return err
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VDATA {
			_, err = b.WriteString("\"vData\":")
			if err != nil {
				return err
			}
			{
				s := s.VData()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VBOOL {
			_, err = b.WriteString("\"vBool\":")
			if err != nil {
				return err
			}
			{
				s := s.VBool()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT8 {
			_, err = b.WriteString("\"vInt8\":")
			if err != nil {
				return err
			}
			{
				s := s.VInt8()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT16 {
			_, err = b.WriteString("\"vInt16\":")
			if err != nil {
				return err
			}
			{
				s := s.VInt16()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT32 {
			_, err = b.WriteString("\"vInt32\":")
			if err != nil {
				return err
			}
			{
				s := s.VInt32()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT64 {
			_, err = b.WriteString("\"vInt64\":")
			if err != nil {
				return err
			}
			{
				s := s.VInt64()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT8 {
			_, err = b.WriteString("\"vUint8\":")
			if err != nil {
				return err
			}
			{
				s := s.VUint8()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT16 {
			_, err = b.WriteString("\"vUint16\":")
			if err != nil {
				return err
			}
			{
				s := s.VUint16()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT32 {
			_, err = b.WriteString("\"vUint32\":")
			if err != nil {
				return err
			}
			{
				s := s.VUint32()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT64 {
			_, err = b.WriteString("\"vUint64\":")
			if err != nil {
				return err
			}
			{
				s := s.VUint64()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VFLOAT32 {
			_, err = b.WriteString("\"vFloat32\":")
			if err != nil {
				return err
			}
			{
				s := s.VFloat32()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VFLOAT64 {
			_, err = b.WriteString("\"vFloat64\":")
			if err != nil {
				return err
			}
			{
				s := s.VFloat64()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VTEXT {
			_, err = b.WriteString("\"vText\":")
			if err != nil {
				return err
			}
			{
				s := s.VText()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		err = b.WriteByte('}')
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"isNull\":")
	if err != nil {
		return err
	}
	{
		s := s.IsNull()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DOIndexValueItem) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s DOIndexValueItem) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("metadata = ")
	if err != nil {
		return err
	}
	{
		s := s.Metadata()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("valueType = ")
	if err != nil {
		return err
	}
	{
		s := s.ValueType()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("value = ")
	if err != nil {
		return err
	}
	{
		s := s.Value()
		err = b.WriteByte('(')
		if err != nil {
			return err
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VDATA {
			_, err = b.WriteString("vData = ")
			if err != nil {
				return err
			}
			{
				s := s.VData()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VBOOL {
			_, err = b.WriteString("vBool = ")
			if err != nil {
				return err
			}
			{
				s := s.VBool()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT8 {
			_, err = b.WriteString("vInt8 = ")
			if err != nil {
				return err
			}
			{
				s := s.VInt8()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT16 {
			_, err = b.WriteString("vInt16 = ")
			if err != nil {
				return err
			}
			{
				s := s.VInt16()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT32 {
			_, err = b.WriteString("vInt32 = ")
			if err != nil {
				return err
			}
			{
				s := s.VInt32()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VINT64 {
			_, err = b.WriteString("vInt64 = ")
			if err != nil {
				return err
			}
			{
				s := s.VInt64()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT8 {
			_, err = b.WriteString("vUint8 = ")
			if err != nil {
				return err
			}
			{
				s := s.VUint8()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT16 {
			_, err = b.WriteString("vUint16 = ")
			if err != nil {
				return err
			}
			{
				s := s.VUint16()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT32 {
			_, err = b.WriteString("vUint32 = ")
			if err != nil {
				return err
			}
			{
				s := s.VUint32()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VUINT64 {
			_, err = b.WriteString("vUint64 = ")
			if err != nil {
				return err
			}
			{
				s := s.VUint64()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VFLOAT32 {
			_, err = b.WriteString("vFloat32 = ")
			if err != nil {
				return err
			}
			{
				s := s.VFloat32()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VFLOAT64 {
			_, err = b.WriteString("vFloat64 = ")
			if err != nil {
				return err
			}
			{
				s := s.VFloat64()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == DOINDEXVALUEITEMVALUE_VTEXT {
			_, err = b.WriteString("vText = ")
			if err != nil {
				return err
			}
			{
				s := s.VText()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		err = b.WriteByte(')')
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("isNull = ")
	if err != nil {
		return err
	}
	{
		s := s.IsNull()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DOIndexValueItem) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type DOIndexValueItem_List C.PointerList

func NewDOIndexValueItemList(s *C.Segment, sz int) DOIndexValueItem_List {
	return DOIndexValueItem_List(s.NewCompositeList(24, 2, sz))
}
func (s DOIndexValueItem_List) Len() int { return C.PointerList(s).Len() }
func (s DOIndexValueItem_List) At(i int) DOIndexValueItem {
	return DOIndexValueItem(C.PointerList(s).At(i).ToStruct())
}
func (s DOIndexValueItem_List) ToArray() []DOIndexValueItem {
	n := s.Len()
	a := make([]DOIndexValueItem, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s DOIndexValueItem_List) Set(i int, item DOIndexValueItem) {
	C.PointerList(s).Set(i, C.Object(item))
}

type DOIndexValueItemAttributeType uint16

const (
	DOINDEXVALUEITEMATTRIBUTETYPE_DATA    DOIndexValueItemAttributeType = 0
	DOINDEXVALUEITEMATTRIBUTETYPE_TEXT    DOIndexValueItemAttributeType = 1
	DOINDEXVALUEITEMATTRIBUTETYPE_BOOL    DOIndexValueItemAttributeType = 2
	DOINDEXVALUEITEMATTRIBUTETYPE_INT8    DOIndexValueItemAttributeType = 3
	DOINDEXVALUEITEMATTRIBUTETYPE_INT16   DOIndexValueItemAttributeType = 4
	DOINDEXVALUEITEMATTRIBUTETYPE_INT32   DOIndexValueItemAttributeType = 5
	DOINDEXVALUEITEMATTRIBUTETYPE_INT64   DOIndexValueItemAttributeType = 6
	DOINDEXVALUEITEMATTRIBUTETYPE_UINT8   DOIndexValueItemAttributeType = 7
	DOINDEXVALUEITEMATTRIBUTETYPE_UINT16  DOIndexValueItemAttributeType = 8
	DOINDEXVALUEITEMATTRIBUTETYPE_UINT32  DOIndexValueItemAttributeType = 9
	DOINDEXVALUEITEMATTRIBUTETYPE_UINT64  DOIndexValueItemAttributeType = 10
	DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT32 DOIndexValueItemAttributeType = 11
	DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT64 DOIndexValueItemAttributeType = 12
)

func (c DOIndexValueItemAttributeType) String() string {
	switch c {
	case DOINDEXVALUEITEMATTRIBUTETYPE_DATA:
		return "data"
	case DOINDEXVALUEITEMATTRIBUTETYPE_TEXT:
		return "text"
	case DOINDEXVALUEITEMATTRIBUTETYPE_BOOL:
		return "bool"
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT8:
		return "int8"
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT16:
		return "int16"
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT32:
		return "int32"
	case DOINDEXVALUEITEMATTRIBUTETYPE_INT64:
		return "int64"
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT8:
		return "uint8"
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT16:
		return "uint16"
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT32:
		return "uint32"
	case DOINDEXVALUEITEMATTRIBUTETYPE_UINT64:
		return "uint64"
	case DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT32:
		return "float32"
	case DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT64:
		return "float64"
	default:
		return ""
	}
}

func DOIndexValueItemAttributeTypeFromString(c string) DOIndexValueItemAttributeType {
	switch c {
	case "data":
		return DOINDEXVALUEITEMATTRIBUTETYPE_DATA
	case "text":
		return DOINDEXVALUEITEMATTRIBUTETYPE_TEXT
	case "bool":
		return DOINDEXVALUEITEMATTRIBUTETYPE_BOOL
	case "int8":
		return DOINDEXVALUEITEMATTRIBUTETYPE_INT8
	case "int16":
		return DOINDEXVALUEITEMATTRIBUTETYPE_INT16
	case "int32":
		return DOINDEXVALUEITEMATTRIBUTETYPE_INT32
	case "int64":
		return DOINDEXVALUEITEMATTRIBUTETYPE_INT64
	case "uint8":
		return DOINDEXVALUEITEMATTRIBUTETYPE_UINT8
	case "uint16":
		return DOINDEXVALUEITEMATTRIBUTETYPE_UINT16
	case "uint32":
		return DOINDEXVALUEITEMATTRIBUTETYPE_UINT32
	case "uint64":
		return DOINDEXVALUEITEMATTRIBUTETYPE_UINT64
	case "float32":
		return DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT32
	case "float64":
		return DOINDEXVALUEITEMATTRIBUTETYPE_FLOAT64
	default:
		return 0
	}
}

type DOIndexValueItemAttributeType_List C.PointerList

func NewDOIndexValueItemAttributeTypeList(s *C.Segment, sz int) DOIndexValueItemAttributeType_List {
	return DOIndexValueItemAttributeType_List(s.NewUInt16List(sz))
}
func (s DOIndexValueItemAttributeType_List) Len() int { return C.UInt16List(s).Len() }
func (s DOIndexValueItemAttributeType_List) At(i int) DOIndexValueItemAttributeType {
	return DOIndexValueItemAttributeType(C.UInt16List(s).At(i))
}
func (s DOIndexValueItemAttributeType_List) ToArray() []DOIndexValueItemAttributeType {
	n := s.Len()
	a := make([]DOIndexValueItemAttributeType, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s DOIndexValueItemAttributeType_List) Set(i int, item DOIndexValueItemAttributeType) {
	C.UInt16List(s).Set(i, uint16(item))
}
func (s DOIndexValueItemAttributeType) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DOIndexValueItemAttributeType) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s DOIndexValueItemAttributeType) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	_, err = b.WriteString(s.String())
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DOIndexValueItemAttributeType) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type DOIndexValue C.Struct

func NewDOIndexValue(s *C.Segment) DOIndexValue      { return DOIndexValue(s.NewStruct(0, 2)) }
func NewRootDOIndexValue(s *C.Segment) DOIndexValue  { return DOIndexValue(s.NewRootStruct(0, 2)) }
func AutoNewDOIndexValue(s *C.Segment) DOIndexValue  { return DOIndexValue(s.NewStructAR(0, 2)) }
func ReadRootDOIndexValue(s *C.Segment) DOIndexValue { return DOIndexValue(s.Root(0).ToStruct()) }
func (s DOIndexValue) Metadata() lib_gc_capnp_common.Metadata {
	return lib_gc_capnp_common.Metadata(C.Struct(s).GetObject(0).ToStruct())
}
func (s DOIndexValue) SetMetadata(v lib_gc_capnp_common.Metadata) {
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s DOIndexValue) Items() DOIndexValueItem_List {
	return DOIndexValueItem_List(C.Struct(s).GetObject(1))
}
func (s DOIndexValue) SetItems(v DOIndexValueItem_List) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s DOIndexValue) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"metadata\":")
	if err != nil {
		return err
	}
	{
		s := s.Metadata()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"items\":")
	if err != nil {
		return err
	}
	{
		s := s.Items()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteJSON(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DOIndexValue) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s DOIndexValue) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("metadata = ")
	if err != nil {
		return err
	}
	{
		s := s.Metadata()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("items = ")
	if err != nil {
		return err
	}
	{
		s := s.Items()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteCapLit(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DOIndexValue) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type DOIndexValue_List C.PointerList

func NewDOIndexValueList(s *C.Segment, sz int) DOIndexValue_List {
	return DOIndexValue_List(s.NewCompositeList(0, 2, sz))
}
func (s DOIndexValue_List) Len() int { return C.PointerList(s).Len() }
func (s DOIndexValue_List) At(i int) DOIndexValue {
	return DOIndexValue(C.PointerList(s).At(i).ToStruct())
}
func (s DOIndexValue_List) ToArray() []DOIndexValue {
	n := s.Len()
	a := make([]DOIndexValue, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s DOIndexValue_List) Set(i int, item DOIndexValue) { C.PointerList(s).Set(i, C.Object(item)) }
