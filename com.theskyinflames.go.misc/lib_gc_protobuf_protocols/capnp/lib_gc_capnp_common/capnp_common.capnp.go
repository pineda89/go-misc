package lib_gc_capnp_common

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
	"math"
)

type Metadata C.Struct

func NewMetadata(s *C.Segment) Metadata      { return Metadata(s.NewStruct(16, 0)) }
func NewRootMetadata(s *C.Segment) Metadata  { return Metadata(s.NewRootStruct(16, 0)) }
func AutoNewMetadata(s *C.Segment) Metadata  { return Metadata(s.NewStructAR(16, 0)) }
func ReadRootMetadata(s *C.Segment) Metadata { return Metadata(s.Root(0).ToStruct()) }
func (s Metadata) Version() uint64           { return C.Struct(s).Get64(0) }
func (s Metadata) SetVersion(v uint64)       { C.Struct(s).Set64(0, v) }
func (s Metadata) Status() int16             { return int16(C.Struct(s).Get16(8)) }
func (s Metadata) SetStatus(v int16)         { C.Struct(s).Set16(8, uint16(v)) }
func (s Metadata) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"version\":")
	if err != nil {
		return err
	}
	{
		s := s.Version()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"status\":")
	if err != nil {
		return err
	}
	{
		s := s.Status()
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
func (s Metadata) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Metadata) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("version = ")
	if err != nil {
		return err
	}
	{
		s := s.Version()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("status = ")
	if err != nil {
		return err
	}
	{
		s := s.Status()
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
func (s Metadata) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Metadata_List C.PointerList

func NewMetadataList(s *C.Segment, sz int) Metadata_List {
	return Metadata_List(s.NewCompositeList(16, 0, sz))
}
func (s Metadata_List) Len() int          { return C.PointerList(s).Len() }
func (s Metadata_List) At(i int) Metadata { return Metadata(C.PointerList(s).At(i).ToStruct()) }
func (s Metadata_List) ToArray() []Metadata {
	n := s.Len()
	a := make([]Metadata, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Metadata_List) Set(i int, item Metadata) { C.PointerList(s).Set(i, C.Object(item)) }

type NullableAttributeText C.Struct

func NewNullableAttributeText(s *C.Segment) NullableAttributeText {
	return NullableAttributeText(s.NewStruct(8, 1))
}
func NewRootNullableAttributeText(s *C.Segment) NullableAttributeText {
	return NullableAttributeText(s.NewRootStruct(8, 1))
}
func AutoNewNullableAttributeText(s *C.Segment) NullableAttributeText {
	return NullableAttributeText(s.NewStructAR(8, 1))
}
func ReadRootNullableAttributeText(s *C.Segment) NullableAttributeText {
	return NullableAttributeText(s.Root(0).ToStruct())
}
func (s NullableAttributeText) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeText) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeText) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeText) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeText) Value() string                { return C.Struct(s).GetObject(0).ToText() }
func (s NullableAttributeText) SetValue(v string)            { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s NullableAttributeText) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeText) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeText) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeText) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeText_List C.PointerList

func NewNullableAttributeTextList(s *C.Segment, sz int) NullableAttributeText_List {
	return NullableAttributeText_List(s.NewCompositeList(8, 1, sz))
}
func (s NullableAttributeText_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeText_List) At(i int) NullableAttributeText {
	return NullableAttributeText(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeText_List) ToArray() []NullableAttributeText {
	n := s.Len()
	a := make([]NullableAttributeText, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeText_List) Set(i int, item NullableAttributeText) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeData C.Struct

func NewNullableAttributeData(s *C.Segment) NullableAttributeData {
	return NullableAttributeData(s.NewStruct(8, 1))
}
func NewRootNullableAttributeData(s *C.Segment) NullableAttributeData {
	return NullableAttributeData(s.NewRootStruct(8, 1))
}
func AutoNewNullableAttributeData(s *C.Segment) NullableAttributeData {
	return NullableAttributeData(s.NewStructAR(8, 1))
}
func ReadRootNullableAttributeData(s *C.Segment) NullableAttributeData {
	return NullableAttributeData(s.Root(0).ToStruct())
}
func (s NullableAttributeData) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeData) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeData) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeData) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeData) Value() []byte                { return C.Struct(s).GetObject(0).ToData() }
func (s NullableAttributeData) SetValue(v []byte)            { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s NullableAttributeData) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeData) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeData) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeData) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeData_List C.PointerList

func NewNullableAttributeDataList(s *C.Segment, sz int) NullableAttributeData_List {
	return NullableAttributeData_List(s.NewCompositeList(8, 1, sz))
}
func (s NullableAttributeData_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeData_List) At(i int) NullableAttributeData {
	return NullableAttributeData(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeData_List) ToArray() []NullableAttributeData {
	n := s.Len()
	a := make([]NullableAttributeData, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeData_List) Set(i int, item NullableAttributeData) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeBool C.Struct

func NewNullableAttributeBool(s *C.Segment) NullableAttributeBool {
	return NullableAttributeBool(s.NewStruct(8, 0))
}
func NewRootNullableAttributeBool(s *C.Segment) NullableAttributeBool {
	return NullableAttributeBool(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeBool(s *C.Segment) NullableAttributeBool {
	return NullableAttributeBool(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeBool(s *C.Segment) NullableAttributeBool {
	return NullableAttributeBool(s.Root(0).ToStruct())
}
func (s NullableAttributeBool) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeBool) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeBool) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeBool) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeBool) Value() bool                  { return C.Struct(s).Get1(1) }
func (s NullableAttributeBool) SetValue(v bool)              { C.Struct(s).Set1(1, v) }
func (s NullableAttributeBool) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeBool) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeBool) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeBool) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeBool_List C.PointerList

func NewNullableAttributeBoolList(s *C.Segment, sz int) NullableAttributeBool_List {
	return NullableAttributeBool_List(s.NewUInt32List(sz))
}
func (s NullableAttributeBool_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeBool_List) At(i int) NullableAttributeBool {
	return NullableAttributeBool(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeBool_List) ToArray() []NullableAttributeBool {
	n := s.Len()
	a := make([]NullableAttributeBool, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeBool_List) Set(i int, item NullableAttributeBool) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeInt8 C.Struct

func NewNullableAttributeInt8(s *C.Segment) NullableAttributeInt8 {
	return NullableAttributeInt8(s.NewStruct(8, 0))
}
func NewRootNullableAttributeInt8(s *C.Segment) NullableAttributeInt8 {
	return NullableAttributeInt8(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeInt8(s *C.Segment) NullableAttributeInt8 {
	return NullableAttributeInt8(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeInt8(s *C.Segment) NullableAttributeInt8 {
	return NullableAttributeInt8(s.Root(0).ToStruct())
}
func (s NullableAttributeInt8) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeInt8) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeInt8) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeInt8) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeInt8) Value() int8                  { return int8(C.Struct(s).Get8(1)) }
func (s NullableAttributeInt8) SetValue(v int8)              { C.Struct(s).Set8(1, uint8(v)) }
func (s NullableAttributeInt8) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeInt8) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeInt8) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeInt8) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeInt8_List C.PointerList

func NewNullableAttributeInt8List(s *C.Segment, sz int) NullableAttributeInt8_List {
	return NullableAttributeInt8_List(s.NewUInt32List(sz))
}
func (s NullableAttributeInt8_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeInt8_List) At(i int) NullableAttributeInt8 {
	return NullableAttributeInt8(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeInt8_List) ToArray() []NullableAttributeInt8 {
	n := s.Len()
	a := make([]NullableAttributeInt8, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeInt8_List) Set(i int, item NullableAttributeInt8) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeInt16 C.Struct

func NewNullableAttributeInt16(s *C.Segment) NullableAttributeInt16 {
	return NullableAttributeInt16(s.NewStruct(8, 0))
}
func NewRootNullableAttributeInt16(s *C.Segment) NullableAttributeInt16 {
	return NullableAttributeInt16(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeInt16(s *C.Segment) NullableAttributeInt16 {
	return NullableAttributeInt16(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeInt16(s *C.Segment) NullableAttributeInt16 {
	return NullableAttributeInt16(s.Root(0).ToStruct())
}
func (s NullableAttributeInt16) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeInt16) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeInt16) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeInt16) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeInt16) Value() int16                 { return int16(C.Struct(s).Get16(4)) }
func (s NullableAttributeInt16) SetValue(v int16)             { C.Struct(s).Set16(4, uint16(v)) }
func (s NullableAttributeInt16) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeInt16) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeInt16) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeInt16) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeInt16_List C.PointerList

func NewNullableAttributeInt16List(s *C.Segment, sz int) NullableAttributeInt16_List {
	return NullableAttributeInt16_List(s.NewUInt64List(sz))
}
func (s NullableAttributeInt16_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeInt16_List) At(i int) NullableAttributeInt16 {
	return NullableAttributeInt16(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeInt16_List) ToArray() []NullableAttributeInt16 {
	n := s.Len()
	a := make([]NullableAttributeInt16, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeInt16_List) Set(i int, item NullableAttributeInt16) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeInt32 C.Struct

func NewNullableAttributeInt32(s *C.Segment) NullableAttributeInt32 {
	return NullableAttributeInt32(s.NewStruct(8, 0))
}
func NewRootNullableAttributeInt32(s *C.Segment) NullableAttributeInt32 {
	return NullableAttributeInt32(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeInt32(s *C.Segment) NullableAttributeInt32 {
	return NullableAttributeInt32(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeInt32(s *C.Segment) NullableAttributeInt32 {
	return NullableAttributeInt32(s.Root(0).ToStruct())
}
func (s NullableAttributeInt32) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeInt32) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeInt32) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeInt32) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeInt32) Value() int32                 { return int32(C.Struct(s).Get32(4)) }
func (s NullableAttributeInt32) SetValue(v int32)             { C.Struct(s).Set32(4, uint32(v)) }
func (s NullableAttributeInt32) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeInt32) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeInt32) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeInt32) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeInt32_List C.PointerList

func NewNullableAttributeInt32List(s *C.Segment, sz int) NullableAttributeInt32_List {
	return NullableAttributeInt32_List(s.NewUInt64List(sz))
}
func (s NullableAttributeInt32_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeInt32_List) At(i int) NullableAttributeInt32 {
	return NullableAttributeInt32(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeInt32_List) ToArray() []NullableAttributeInt32 {
	n := s.Len()
	a := make([]NullableAttributeInt32, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeInt32_List) Set(i int, item NullableAttributeInt32) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeInt64 C.Struct

func NewNullableAttributeInt64(s *C.Segment) NullableAttributeInt64 {
	return NullableAttributeInt64(s.NewStruct(16, 0))
}
func NewRootNullableAttributeInt64(s *C.Segment) NullableAttributeInt64 {
	return NullableAttributeInt64(s.NewRootStruct(16, 0))
}
func AutoNewNullableAttributeInt64(s *C.Segment) NullableAttributeInt64 {
	return NullableAttributeInt64(s.NewStructAR(16, 0))
}
func ReadRootNullableAttributeInt64(s *C.Segment) NullableAttributeInt64 {
	return NullableAttributeInt64(s.Root(0).ToStruct())
}
func (s NullableAttributeInt64) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeInt64) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeInt64) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeInt64) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeInt64) Value() int64                 { return int64(C.Struct(s).Get64(8)) }
func (s NullableAttributeInt64) SetValue(v int64)             { C.Struct(s).Set64(8, uint64(v)) }
func (s NullableAttributeInt64) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeInt64) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeInt64) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeInt64) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeInt64_List C.PointerList

func NewNullableAttributeInt64List(s *C.Segment, sz int) NullableAttributeInt64_List {
	return NullableAttributeInt64_List(s.NewCompositeList(16, 0, sz))
}
func (s NullableAttributeInt64_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeInt64_List) At(i int) NullableAttributeInt64 {
	return NullableAttributeInt64(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeInt64_List) ToArray() []NullableAttributeInt64 {
	n := s.Len()
	a := make([]NullableAttributeInt64, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeInt64_List) Set(i int, item NullableAttributeInt64) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeUInt8 C.Struct

func NewNullableAttributeUInt8(s *C.Segment) NullableAttributeUInt8 {
	return NullableAttributeUInt8(s.NewStruct(8, 0))
}
func NewRootNullableAttributeUInt8(s *C.Segment) NullableAttributeUInt8 {
	return NullableAttributeUInt8(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeUInt8(s *C.Segment) NullableAttributeUInt8 {
	return NullableAttributeUInt8(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeUInt8(s *C.Segment) NullableAttributeUInt8 {
	return NullableAttributeUInt8(s.Root(0).ToStruct())
}
func (s NullableAttributeUInt8) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeUInt8) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeUInt8) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeUInt8) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeUInt8) Value() uint8                 { return C.Struct(s).Get8(1) }
func (s NullableAttributeUInt8) SetValue(v uint8)             { C.Struct(s).Set8(1, v) }
func (s NullableAttributeUInt8) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeUInt8) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeUInt8) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeUInt8) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeUInt8_List C.PointerList

func NewNullableAttributeUInt8List(s *C.Segment, sz int) NullableAttributeUInt8_List {
	return NullableAttributeUInt8_List(s.NewUInt32List(sz))
}
func (s NullableAttributeUInt8_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeUInt8_List) At(i int) NullableAttributeUInt8 {
	return NullableAttributeUInt8(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeUInt8_List) ToArray() []NullableAttributeUInt8 {
	n := s.Len()
	a := make([]NullableAttributeUInt8, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeUInt8_List) Set(i int, item NullableAttributeUInt8) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeUInt16 C.Struct

func NewNullableAttributeUInt16(s *C.Segment) NullableAttributeUInt16 {
	return NullableAttributeUInt16(s.NewStruct(8, 0))
}
func NewRootNullableAttributeUInt16(s *C.Segment) NullableAttributeUInt16 {
	return NullableAttributeUInt16(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeUInt16(s *C.Segment) NullableAttributeUInt16 {
	return NullableAttributeUInt16(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeUInt16(s *C.Segment) NullableAttributeUInt16 {
	return NullableAttributeUInt16(s.Root(0).ToStruct())
}
func (s NullableAttributeUInt16) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeUInt16) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeUInt16) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeUInt16) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeUInt16) Value() uint16                { return C.Struct(s).Get16(4) }
func (s NullableAttributeUInt16) SetValue(v uint16)            { C.Struct(s).Set16(4, v) }
func (s NullableAttributeUInt16) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeUInt16) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeUInt16) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeUInt16) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeUInt16_List C.PointerList

func NewNullableAttributeUInt16List(s *C.Segment, sz int) NullableAttributeUInt16_List {
	return NullableAttributeUInt16_List(s.NewUInt64List(sz))
}
func (s NullableAttributeUInt16_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeUInt16_List) At(i int) NullableAttributeUInt16 {
	return NullableAttributeUInt16(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeUInt16_List) ToArray() []NullableAttributeUInt16 {
	n := s.Len()
	a := make([]NullableAttributeUInt16, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeUInt16_List) Set(i int, item NullableAttributeUInt16) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeUInt32 C.Struct

func NewNullableAttributeUInt32(s *C.Segment) NullableAttributeUInt32 {
	return NullableAttributeUInt32(s.NewStruct(8, 0))
}
func NewRootNullableAttributeUInt32(s *C.Segment) NullableAttributeUInt32 {
	return NullableAttributeUInt32(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeUInt32(s *C.Segment) NullableAttributeUInt32 {
	return NullableAttributeUInt32(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeUInt32(s *C.Segment) NullableAttributeUInt32 {
	return NullableAttributeUInt32(s.Root(0).ToStruct())
}
func (s NullableAttributeUInt32) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeUInt32) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeUInt32) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeUInt32) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeUInt32) Value() uint32                { return C.Struct(s).Get32(4) }
func (s NullableAttributeUInt32) SetValue(v uint32)            { C.Struct(s).Set32(4, v) }
func (s NullableAttributeUInt32) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeUInt32) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeUInt32) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeUInt32) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeUInt32_List C.PointerList

func NewNullableAttributeUInt32List(s *C.Segment, sz int) NullableAttributeUInt32_List {
	return NullableAttributeUInt32_List(s.NewUInt64List(sz))
}
func (s NullableAttributeUInt32_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeUInt32_List) At(i int) NullableAttributeUInt32 {
	return NullableAttributeUInt32(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeUInt32_List) ToArray() []NullableAttributeUInt32 {
	n := s.Len()
	a := make([]NullableAttributeUInt32, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeUInt32_List) Set(i int, item NullableAttributeUInt32) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeUInt64 C.Struct

func NewNullableAttributeUInt64(s *C.Segment) NullableAttributeUInt64 {
	return NullableAttributeUInt64(s.NewStruct(16, 0))
}
func NewRootNullableAttributeUInt64(s *C.Segment) NullableAttributeUInt64 {
	return NullableAttributeUInt64(s.NewRootStruct(16, 0))
}
func AutoNewNullableAttributeUInt64(s *C.Segment) NullableAttributeUInt64 {
	return NullableAttributeUInt64(s.NewStructAR(16, 0))
}
func ReadRootNullableAttributeUInt64(s *C.Segment) NullableAttributeUInt64 {
	return NullableAttributeUInt64(s.Root(0).ToStruct())
}
func (s NullableAttributeUInt64) IsNull() bool                 { return C.Struct(s).Get1(0) }
func (s NullableAttributeUInt64) SetIsNull(v bool)             { C.Struct(s).Set1(0, v) }
func (s NullableAttributeUInt64) ValueType() AttributeType     { return AttributeType(C.Struct(s).Get16(2)) }
func (s NullableAttributeUInt64) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeUInt64) Value() uint64                { return C.Struct(s).Get64(8) }
func (s NullableAttributeUInt64) SetValue(v uint64)            { C.Struct(s).Set64(8, v) }
func (s NullableAttributeUInt64) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeUInt64) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeUInt64) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeUInt64) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeUInt64_List C.PointerList

func NewNullableAttributeUInt64List(s *C.Segment, sz int) NullableAttributeUInt64_List {
	return NullableAttributeUInt64_List(s.NewCompositeList(16, 0, sz))
}
func (s NullableAttributeUInt64_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeUInt64_List) At(i int) NullableAttributeUInt64 {
	return NullableAttributeUInt64(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeUInt64_List) ToArray() []NullableAttributeUInt64 {
	n := s.Len()
	a := make([]NullableAttributeUInt64, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeUInt64_List) Set(i int, item NullableAttributeUInt64) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeFloat32 C.Struct

func NewNullableAttributeFloat32(s *C.Segment) NullableAttributeFloat32 {
	return NullableAttributeFloat32(s.NewStruct(8, 0))
}
func NewRootNullableAttributeFloat32(s *C.Segment) NullableAttributeFloat32 {
	return NullableAttributeFloat32(s.NewRootStruct(8, 0))
}
func AutoNewNullableAttributeFloat32(s *C.Segment) NullableAttributeFloat32 {
	return NullableAttributeFloat32(s.NewStructAR(8, 0))
}
func ReadRootNullableAttributeFloat32(s *C.Segment) NullableAttributeFloat32 {
	return NullableAttributeFloat32(s.Root(0).ToStruct())
}
func (s NullableAttributeFloat32) IsNull() bool     { return C.Struct(s).Get1(0) }
func (s NullableAttributeFloat32) SetIsNull(v bool) { C.Struct(s).Set1(0, v) }
func (s NullableAttributeFloat32) ValueType() AttributeType {
	return AttributeType(C.Struct(s).Get16(2))
}
func (s NullableAttributeFloat32) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeFloat32) Value() float32               { return math.Float32frombits(C.Struct(s).Get32(4)) }
func (s NullableAttributeFloat32) SetValue(v float32)           { C.Struct(s).Set32(4, math.Float32bits(v)) }
func (s NullableAttributeFloat32) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeFloat32) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeFloat32) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeFloat32) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeFloat32_List C.PointerList

func NewNullableAttributeFloat32List(s *C.Segment, sz int) NullableAttributeFloat32_List {
	return NullableAttributeFloat32_List(s.NewUInt64List(sz))
}
func (s NullableAttributeFloat32_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeFloat32_List) At(i int) NullableAttributeFloat32 {
	return NullableAttributeFloat32(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeFloat32_List) ToArray() []NullableAttributeFloat32 {
	n := s.Len()
	a := make([]NullableAttributeFloat32, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeFloat32_List) Set(i int, item NullableAttributeFloat32) {
	C.PointerList(s).Set(i, C.Object(item))
}

type NullableAttributeFloat64 C.Struct

func NewNullableAttributeFloat64(s *C.Segment) NullableAttributeFloat64 {
	return NullableAttributeFloat64(s.NewStruct(16, 0))
}
func NewRootNullableAttributeFloat64(s *C.Segment) NullableAttributeFloat64 {
	return NullableAttributeFloat64(s.NewRootStruct(16, 0))
}
func AutoNewNullableAttributeFloat64(s *C.Segment) NullableAttributeFloat64 {
	return NullableAttributeFloat64(s.NewStructAR(16, 0))
}
func ReadRootNullableAttributeFloat64(s *C.Segment) NullableAttributeFloat64 {
	return NullableAttributeFloat64(s.Root(0).ToStruct())
}
func (s NullableAttributeFloat64) IsNull() bool     { return C.Struct(s).Get1(0) }
func (s NullableAttributeFloat64) SetIsNull(v bool) { C.Struct(s).Set1(0, v) }
func (s NullableAttributeFloat64) ValueType() AttributeType {
	return AttributeType(C.Struct(s).Get16(2))
}
func (s NullableAttributeFloat64) SetValueType(v AttributeType) { C.Struct(s).Set16(2, uint16(v)) }
func (s NullableAttributeFloat64) Value() float64               { return math.Float64frombits(C.Struct(s).Get64(8)) }
func (s NullableAttributeFloat64) SetValue(v float64)           { C.Struct(s).Set64(8, math.Float64bits(v)) }
func (s NullableAttributeFloat64) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
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
func (s NullableAttributeFloat64) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s NullableAttributeFloat64) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
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
func (s NullableAttributeFloat64) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type NullableAttributeFloat64_List C.PointerList

func NewNullableAttributeFloat64List(s *C.Segment, sz int) NullableAttributeFloat64_List {
	return NullableAttributeFloat64_List(s.NewCompositeList(16, 0, sz))
}
func (s NullableAttributeFloat64_List) Len() int { return C.PointerList(s).Len() }
func (s NullableAttributeFloat64_List) At(i int) NullableAttributeFloat64 {
	return NullableAttributeFloat64(C.PointerList(s).At(i).ToStruct())
}
func (s NullableAttributeFloat64_List) ToArray() []NullableAttributeFloat64 {
	n := s.Len()
	a := make([]NullableAttributeFloat64, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NullableAttributeFloat64_List) Set(i int, item NullableAttributeFloat64) {
	C.PointerList(s).Set(i, C.Object(item))
}

type AttributeType uint16

const (
	ATTRIBUTETYPE_DATA    AttributeType = 0
	ATTRIBUTETYPE_TEXT    AttributeType = 1
	ATTRIBUTETYPE_BOOL    AttributeType = 2
	ATTRIBUTETYPE_INT8    AttributeType = 3
	ATTRIBUTETYPE_INT16   AttributeType = 4
	ATTRIBUTETYPE_INT32   AttributeType = 5
	ATTRIBUTETYPE_INT64   AttributeType = 6
	ATTRIBUTETYPE_UINT8   AttributeType = 7
	ATTRIBUTETYPE_UINT16  AttributeType = 8
	ATTRIBUTETYPE_UINT32  AttributeType = 9
	ATTRIBUTETYPE_UINT64  AttributeType = 10
	ATTRIBUTETYPE_FLOAT32 AttributeType = 11
	ATTRIBUTETYPE_FLOAT64 AttributeType = 12
)

func (c AttributeType) String() string {
	switch c {
	case ATTRIBUTETYPE_DATA:
		return "data"
	case ATTRIBUTETYPE_TEXT:
		return "text"
	case ATTRIBUTETYPE_BOOL:
		return "bool"
	case ATTRIBUTETYPE_INT8:
		return "int8"
	case ATTRIBUTETYPE_INT16:
		return "int16"
	case ATTRIBUTETYPE_INT32:
		return "int32"
	case ATTRIBUTETYPE_INT64:
		return "int64"
	case ATTRIBUTETYPE_UINT8:
		return "uint8"
	case ATTRIBUTETYPE_UINT16:
		return "uint16"
	case ATTRIBUTETYPE_UINT32:
		return "uint32"
	case ATTRIBUTETYPE_UINT64:
		return "uint64"
	case ATTRIBUTETYPE_FLOAT32:
		return "float32"
	case ATTRIBUTETYPE_FLOAT64:
		return "float64"
	default:
		return ""
	}
}

func AttributeTypeFromString(c string) AttributeType {
	switch c {
	case "data":
		return ATTRIBUTETYPE_DATA
	case "text":
		return ATTRIBUTETYPE_TEXT
	case "bool":
		return ATTRIBUTETYPE_BOOL
	case "int8":
		return ATTRIBUTETYPE_INT8
	case "int16":
		return ATTRIBUTETYPE_INT16
	case "int32":
		return ATTRIBUTETYPE_INT32
	case "int64":
		return ATTRIBUTETYPE_INT64
	case "uint8":
		return ATTRIBUTETYPE_UINT8
	case "uint16":
		return ATTRIBUTETYPE_UINT16
	case "uint32":
		return ATTRIBUTETYPE_UINT32
	case "uint64":
		return ATTRIBUTETYPE_UINT64
	case "float32":
		return ATTRIBUTETYPE_FLOAT32
	case "float64":
		return ATTRIBUTETYPE_FLOAT64
	default:
		return 0
	}
}

type AttributeType_List C.PointerList

func NewAttributeTypeList(s *C.Segment, sz int) AttributeType_List {
	return AttributeType_List(s.NewUInt16List(sz))
}
func (s AttributeType_List) Len() int               { return C.UInt16List(s).Len() }
func (s AttributeType_List) At(i int) AttributeType { return AttributeType(C.UInt16List(s).At(i)) }
func (s AttributeType_List) ToArray() []AttributeType {
	n := s.Len()
	a := make([]AttributeType, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s AttributeType_List) Set(i int, item AttributeType) { C.UInt16List(s).Set(i, uint16(item)) }
func (s AttributeType) WriteJSON(w io.Writer) error {
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
func (s AttributeType) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s AttributeType) WriteCapLit(w io.Writer) error {
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
func (s AttributeType) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}
