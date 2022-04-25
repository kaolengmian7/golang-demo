package defensive_program

import (
	"log"
	"math"
	"testing"
)

func TestUInt(t *testing.T) {
	a := uint64(math.MaxUint64)
	b := uint64(1)
	log.Println(a + b)

	if math.MaxUint64-a < b {
		panic("")
	} else {
		log.Println(a + b)
	}
}

func TestStr(t *testing.T) {
	log.Println(3677554 * 16 / 1024 / 1024)
}
