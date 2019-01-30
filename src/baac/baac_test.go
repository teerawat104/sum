package baac_test

import (
	. "baac"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 5
	actual := MyAdd(3, 2)
	if expected != actual {
		t.Error("Nope")
	}
}
