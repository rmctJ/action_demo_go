package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Expected %d return %d", 1 + 2, result)
	}
}

func TestPi(t *testing.T) {
	result := Pi()
	if result != math.Pi {
		t.Errorf("Expected %f return %f", math.Pi, result)
	}
}