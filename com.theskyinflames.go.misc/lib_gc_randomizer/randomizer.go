package lib_gc_randomizer

import (
	M_RAND "math/rand"
	"time"

	"bytes"
	"crypto/rand"
	"encoding/binary"
)

func init() {
	// Set random seed
	M_RAND.Seed(time.Now().Unix())
}

func GetRandom() int64 {
	return M_RAND.Int63()
}

func GetRingRandom(ring_size int32) (int32, error) {

	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}

	var p int32
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &p)

	p = p % ring_size
	if p < 0 {
		p = p * -1
	}
	return p, nil
}
