package sum_test

import (
	. "sum"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 5
	actual := myAdd(3, 2)
	if expected != actual {
		t.Error("Nope")
	}
}
