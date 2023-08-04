package services

import "go-simple-testing/repository"

type CalcualtorService interface {
	CalculatorAdd() (*float64, error)
	CalculatorSub() (*float64, error)
}

type calculatorService struct {
	repo repository.CalcualtorRepository
}

func (c *calculatorService) CalculatorAdd() (*float64, error) {
	return c.repo.Add()
}

func (c *calculatorService) CalculatorSub() (*float64, error) {
	return c.repo.Sub()
}

func NewCalculatorService(repo repository.CalcualtorRepository) CalcualtorService {
	return &calculatorService{repo: repo}
}
