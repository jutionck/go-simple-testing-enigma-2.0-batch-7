package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// pembuatan calculator mock -> untuk mock service
type CalCulatorMock struct {
	mock.Mock
}

func (c *CalCulatorMock) Add() (*float64, error) {
	args := c.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float64), nil
}

func (c *CalCulatorMock) Sub() (*float64, error) {
	args := c.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float64), nil
}

// pembuatan struct calculator suite untuk kebutuhan pembuatan test case
type CalculatorTestSuite struct {
	suite.Suite
	calculator Calculator
	calMock *CalCulatorMock
}

// SetupTest() => untuk melakuakn setup test case
func (suite *CalculatorTestSuite) SetupTest() {
	suite.calMock = new(CalCulatorMock)
	suite.calculator = Calculator{
		Num1: 7,
		Num2: 1,
	}
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Success() {
	expected := 8.0
	suite.calMock.On("Add").Return(&expected, nil)
	actual, err := suite.calculator.Add()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, *actual)
}

func (suite *CalculatorTestSuite) TestCalculatorAdd_Fail() {
	expected := 10.0
	suite.calMock.On("Add").Return(&expected, nil)
	_, err := suite.calculator.Add()
	assert.Nil(suite.T(), err)
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Success() {
	expected := 6.0
	suite.calMock.On("Sub").Return(&expected, nil)
	actual, err := suite.calculator.Sub()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, *actual)
}

func (suite *CalculatorTestSuite) TestCalculatorSub_Fail() {
	expected := 10.0
	suite.calMock.On("Sub").Return(&expected, nil)
	_, err := suite.calculator.Sub()
	assert.Nil(suite.T(), err)
}

// untuk running test case yang dibuat (diatas)
func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
