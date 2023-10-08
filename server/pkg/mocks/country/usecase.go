package mocks

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/dto"
	"github.com/stretchr/testify/mock"
)

type CountryUseCaseMock struct {
	mock.Mock
}

func NewCountryUseCase() *CountryUseCaseMock {
	return new(CountryUseCaseMock)
}

func (m *CountryUseCaseMock) GetCountries(ctx context.Context) ([]*dto.CountryResponse, error) {
	ret := m.Called(ctx)

	r0 := ret.Get(0).([]*dto.CountryResponse)

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}
	return r0, r1
}
