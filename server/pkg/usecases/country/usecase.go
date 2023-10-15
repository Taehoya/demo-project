package usecase

import (
	"context"

	"github.com/Taehoya/pocket-mate/pkg/dto"
)

type CountryUsecase struct {
	repository CountryRepository
}

func NewCountryUseCase(repository CountryRepository) *CountryUsecase {
	return &CountryUsecase{
		repository: repository,
	}
}

func (u *CountryUsecase) GetCountries(ctx context.Context) ([]*dto.CountryResponse, error) {
	countries, err := u.repository.GetCountries(ctx)

	if err != nil {
		return nil, err
	}

	countryResp := dto.CreateCountryResponseList(countries)

	return countryResp, nil
}
