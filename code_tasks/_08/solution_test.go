package _08

import (
	"math/rand"
	"testing"
)

const (
	numRandomNumbers int   = 100
	seed             int64 = 42
)

var randomNumbers []int64

func init() {
	randomNumbers = make([]int64, 0, numRandomNumbers)
	source := rand.NewSource(seed)
	for i := 0; i < numRandomNumbers; i++ {
		randomNumbers = append(randomNumbers, source.Int63())
	}
}

func TestWithPowersOfTwo(t *testing.T) {
	for bit, actualValue := 0, int64(1); bit < 64; bit, actualValue = bit+1, actualValue*2 {
		if value := WithIthBitSet(0, bit); value != actualValue {
			t.Errorf("values are different when %d bit is set: %d != %d", bit, value, actualValue)
		} else if value := WithIthBitUnset(value, bit); value != 0 {
			t.Errorf("value is not zero after unsetting %d bit: %d != 0", bit, value)
		}
	}
}

func TestSetPowersOfTwo(t *testing.T) {
	for bit, actualValue := 0, int64(1); bit < 64; bit, actualValue = bit+1, actualValue*2 {
		value := int64(0)
		if SetIthBit(&value, bit); value != actualValue {
			t.Errorf("values are different when %d bit is set: %d != %d", bit, value, actualValue)
		} else if UnsetIthBit(&value, bit); value != 0 {
			t.Errorf("value is not zero after unsetting %d bit: %d != 0", bit, value)
		}
	}
}

func TestWithNegativeBits(t *testing.T) {
	for bit := -10; bit < 0; bit++ {
		if value := WithIthBitSet(0, bit); value != 0 {
			t.Errorf("value is not zero after setting negative %d bit: %d != 0", bit, value)
		} else if value := WithIthBitUnset(value, bit); value != 0 {
			t.Errorf("value is not zero after unsetting negative %d bit: %d != 0", bit, value)
		}
	}
}

func TestSetNegativeBits(t *testing.T) {
	for bit := -10; bit < 0; bit++ {
		value := int64(0)
		if SetIthBit(&value, bit); value != 0 {
			t.Errorf("value is not zero after setting negative %d bit: %d != 0", bit, value)
		} else if UnsetIthBit(&value, bit); value != 0 {
			t.Errorf("value is not zero after unsetting negative %d bit: %d != 0", bit, value)
		}
	}
}

func TestWith64thAndGreaterBits(t *testing.T) {
	for bit := 64; bit < 128; bit++ {
		if value := WithIthBitSet(0, bit); value != 0 {
			t.Errorf("value isn't zero after setting %d bit beyond the boundary of 64: %d != 0", bit, value)
		} else if value := WithIthBitUnset(value, bit); value != 0 {
			t.Errorf("value isn't zero after unsetting %d bit beyond the boundary of 64: %d != 0", bit, value)
		}
	}
}

func TestSet64thAndGreaterBits(t *testing.T) {
	for bit := 64; bit < 128; bit++ {
		value := int64(0)
		if SetIthBit(&value, bit); value != 0 {
			t.Errorf("value isn't zero after setting %d bit beyond the boundary of 64: %d != 0", bit, value)
		} else if UnsetIthBit(&value, bit); value != 0 {
			t.Errorf("value isn't zero after unsetting %d bit beyond the boundary of 64: %d != 0", bit, value)
		}
	}
}

func TestWithRandomNumbers(t *testing.T) {
	for _, actualValue := range randomNumbers {
		for bit := 0; bit < 64; bit++ {
			withBit := int64(1) << bit
			if actualValue&withBit != 0 {
				if value := WithIthBitUnset(actualValue, bit); actualValue-withBit != value {
					t.Errorf("values are different after subtraction of %d bit: %d - %d != %d <=> %d != %d",
						bit, actualValue, withBit, value, actualValue-withBit, value)
				}
			} else {
				if value := WithIthBitSet(actualValue, bit); actualValue+withBit != value {
					t.Errorf("values are different after addition of %d bit: %d - %d != %d <=> %d != %d",
						bit, actualValue, withBit, value, actualValue-withBit, value)
				}
			}
		}
	}
}
