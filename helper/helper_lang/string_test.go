package helper_lang

import (
	"math/rand"
	"testing"
)

func TestIsBlank(t *testing.T) {
	t.Log(IsBlank(""))
	t.Log(IsBlank("pleuvoir"))
}

func TestIsAnyBlank(t *testing.T) {
	t.Log(IsAnyBlank("1", ""))
}

func TestIf(t *testing.T) {
	var number int
	a := If(rand.Int() > 1, 12, 13)
	number = a.(int)
	t.Log(number)
}

func TestToFloat64(t *testing.T) {
	val := "3.14"
	if value, err := ToFloat64(val); err == nil {
		t.Log(value)
	}
}

func TestToUint64(t *testing.T) {
	val := "3"
	if value, err := ToUint64(val); err == nil {
		t.Log(value)
	}
}
