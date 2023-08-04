package repository

type CalcualtorRepository interface{
	Add() (*float64, error)
	Sub() (*float64, error)
}

type calculatorRepository struct {
	Num1 float64
	Num2 float64
}

func (c *calculatorRepository) Add() (*float64, error) {
	result := c.Num1 + c.Num2
	return &result, nil
}

func (c *calculatorRepository) Sub() (*float64, error) {
	result := c.Num1 - c.Num2
	return &result, nil
}

func NewCalculatorRepository() CalcualtorRepository {
	return &calculatorRepository{}
}