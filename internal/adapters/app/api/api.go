package api

import (
	"github.com/imjowend/hex-arch-golang/internal/ports"
)

type Adapter struct {
	arith ports.AritmethicPort
}

func NewAdapter(arith ports.AritmethicPort) *Adapter {
	return &Adapter{
		arith: arith,
	}
}

func (apiAdapter Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiAdapter Adapter) GetSubstraction(a, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiAdapter Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiAdapter Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
