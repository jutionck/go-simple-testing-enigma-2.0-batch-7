package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorAdd_Success(t *testing.T) {
	cal := &Calculator{
		Num1: 6,
		Num2: 1,
	}
	var expected float64 = 7
	actual, err := cal.Add()
	assert.Nil(t, err)
	assert.Equal(t, expected, *actual)
}

func TestCalculatorAdd_Fail(t *testing.T) {
	cal := &Calculator{
		Num1: 6,
		Num2: 1,
	}
	var expected float64 = 5
	actual, err := cal.Add()
	assert.NoError(t, err)
	assert.NotEqual(t, expected, *actual)
}

func TestCalculatorSub_Success(t *testing.T) {
	cal := &Calculator{
		Num1: 10,
		Num2: 5,
	}
	var expected float64 = 5
	actual, err := cal.Sub()
	assert.Nil(t, err)
	assert.Equal(t, expected, *actual)
}

func TestCalculatorSub_Fail(t *testing.T) {
	cal := &Calculator{
		Num1: 6,
		Num2: 1,
	}

	var expected float64 = 10
	actual, err := cal.Sub()
	assert.NoError(t, err)
	assert.NotEqual(t, expected, *actual)
}