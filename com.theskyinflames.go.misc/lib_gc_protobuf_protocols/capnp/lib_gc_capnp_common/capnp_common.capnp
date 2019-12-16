@0x8f81b5083ef1d2dc;

using Go = import "../go.capnp";

$Go.package("lib_gc_capnp_common");
$Go.import("beanstalkapp.com/serhstourism/mdh-common/com.serhs.mdh.common/lib_gc_protobuf_protocols/capnp/lib_gc_capnp_common"); # import it self.

struct Metadata {
	version @0: UInt64;	# version of the CO into cache
	status @1: Int16;  	# status of the CO into cache
}

struct NullableAttributeText {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Text;					# attribute value
}

struct NullableAttributeData {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Data;					# attribute value
}

struct NullableAttributeBool {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Bool;					# attribute value
}

struct NullableAttributeInt8 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Int8;					# attribute value
}

struct NullableAttributeInt16 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Int16;				# attribute value
}

struct NullableAttributeInt32 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Int32;				# attribute value
}

struct NullableAttributeInt64 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Int64;				# attribute value
}


struct NullableAttributeUInt8 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: UInt8;				# attribute value
}

struct NullableAttributeUInt16 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: UInt16;				# attribute value
}

struct NullableAttributeUInt32 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: UInt32;				# attribute value
}

struct NullableAttributeUInt64 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: UInt64;				# attribute value
}


struct NullableAttributeFloat32 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Float32;				# attribute value
}

struct NullableAttributeFloat64 {
	isNull @0: Bool;				# is null ?
	valueType @1 :AttributeType;	# attribute type
	value @2: Float64;				# attribute value
}


# Attribute types
enum AttributeType {
	data @0;
	text @1;
	bool @2;
	int8 @3;
	int16 @4;
	int32 @5;
	int64 @6;
	uint8 @7;
	uint16 @8;
	uint32 @9;
	uint64 @10;
	float32 @11;
	float64 @12;
}

