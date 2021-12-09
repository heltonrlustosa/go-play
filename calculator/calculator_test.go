package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Press(t *testing.T) {
	c := &Calculator{}

	c.Press(10)

	assert.Equal(t, 10, c.Result())
}

func TestCalculator_Add(t *testing.T) {
	c := &Calculator{}

	c.Add(3)

	assert.Equal(t, 3, c.Result())
}

func TestCalculator_Sub(t *testing.T) {
	c := &Calculator{}
	c.Press(10)

	c.Sub(3)

	assert.Equal(t, 7, c.Result())
}

func TestCalculator_MultiplyBy(t *testing.T) {
	c := &Calculator{}
	c.Press(10)

	c.MultiplyBy(10)

	assert.Equal(t, 100, c.Result())
}

func TestCalculator_Clear(t *testing.T) {
	c := &Calculator{}
	c.Press(10)

	c.Clear()

	assert.Equal(t, 0, c.Result())
}

func TestCalculator_Result(t *testing.T) {
	c := &Calculator{}
	c.Press(10)

	assert.Equal(t, c.Result(), 10)
}
