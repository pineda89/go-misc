@0xbe95e4d345006c26;

using Go = import "../../../../lib_gc_protobuf_protocols/capnp/go.capnp";
using Common = import "../../../../lib_gc_protobuf_protocols/capnp/lib_gc_capnp_common/capnp_common.capnp";
$Go.package("lib_idx_value");
$Go.import("lib_idx_value/doIndexValue"); # import it self.


struct DOIndexValueItem{
	
	metadata @0:Common.Metadata;		# CO metadata
	
	#value @0 :Data;
	
	valueType @1 :AttributeType;

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
									
	value :union{
		vData @2: Data;
		vBool @3 :Bool;
		vInt8 @4: Int8;
		vInt16 @5: Int16;
		vInt32 @6: Int32;
		vInt64 @7: Int64;
		vUint8 @8: UInt8;
		vUint16 @9: UInt16;
		vUint32 @10: UInt32;
		vUint64 @11: UInt64;
		vFloat32 @12: Float32;
		vFloat64 @13: Float64;
		vText @14: Text;
	}

    isNull @15 :Bool;
}

struct DOIndexValue {
	metadata @0:Common.Metadata;		# CO metadata
	items @1: List(DOIndexValueItem);
}


