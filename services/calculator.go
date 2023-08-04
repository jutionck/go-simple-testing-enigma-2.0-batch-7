package services

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (c *Calculator) Add() (*float64, error) {
	result := c.Num1 + c.Num2
	return &result, nil
}

func (c *Calculator) Sub() (*float64, error) {
	result := c.Num1 - c.Num2
	return &result, nil
}