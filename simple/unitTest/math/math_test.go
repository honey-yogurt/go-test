package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// expected/actual
func TestAbs(t *testing.T) {
	actual := Abs(-1)
	expected := float64(1)
	if actual != expected {
		t.Errorf("Abs(-1) = %f; expected = 1", actual)
	}
}

// want/got
func TestMax(t *testing.T) {
	got := Max(-1, -1.5)
	want := float64(-1)
	assert.Equal(t, want, got)
}

func TestMaxMultiple(t *testing.T) {
	max := []struct {
		x, y     float64
		expected float64
	}{
		{-1, 2, 2},
		{-0.5, -1, -0.5},
		{2, 3, 3},
		{0, 0, 0},
	}

	for _, v := range max {
		actual := Max(v.x, v.y)
		assert.Equal(t, v.expected, actual)
	}
}
