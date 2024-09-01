package types

import (
	"testing"
)

func TestBool(t *testing.T) {
	b := NewBool().WithValue(true)
	defer b.Clean()

	var values []int

	i := 0

	for v := range b.Each() {
		values = append(values, i)

		i++

		v.Set(i != 5)
	}

	if len(values) != 5 {
		t.Errorf("Expected 5 values, got %d", len(values))
	}

	if values[0] != 0 {
		t.Errorf("Expected 0, got %d", values[0])
	}

	if values[1] != 1 {
		t.Errorf("Expected 1, got %d", values[1])
	}

	if values[2] != 2 {
		t.Errorf("Expected 2, got %d", values[2])
	}

	if values[3] != 3 {
		t.Errorf("Expected 3, got %d", values[3])
	}

	if values[4] != 4 {
		t.Errorf("Expected 4, got %d", values[4])
	}
}
