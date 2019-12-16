package lib_gc_capnp_helpers

import (
	C "github.com/glycerine/go-capnproto"

	"bytes"
)

func HSerializeAsCapNProto(capnp_segment *C.Segment) (*[]byte, error) {
	var _b bytes.Buffer

	if _, err := capnp_segment.WriteTo(&_b); err == nil {
		__b := _b.Bytes()
		return &__b, nil
	} else {
		return nil, err
	}
}
