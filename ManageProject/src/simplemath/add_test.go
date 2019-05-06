package simplemath

import "testing"

func TestAdd(t *testing.T)  {
	r := Add(1, 2)
	if r != 2 {
		t.Log("Add(1,2) error, got %d, expected 3.", r)
	}
}
