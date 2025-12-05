package pubpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	const (
		x        = 2
		y        = 3
		expected = 5
	)
	actual := Add(x, y)
	assert.Equal(t, expected, actual)
}

func TestMinus(t *testing.T) {
	const (
		x        = 2
		y        = 3
		expected = -1
	)
	actual := Minus(x, y)
	assert.Equal(t, expected, actual)
}

func TestMultiply(t *testing.T) {
	const (
		x        = 2
		y        = 3
		expected = 6
	)
	actual := Multiply(x, y)
	assert.Equal(t, expected, actual)
}

func TestDivide(t *testing.T) {
	const (
		x        = 2
		y        = 3
		expected = 0
	)
	actual := Divide(x, y)
	assert.Equal(t, expected, actual)
}
