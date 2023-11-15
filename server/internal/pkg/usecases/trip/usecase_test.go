package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Taehoya/pocket-mate/internal/pkg/dto"
	"github.com/Taehoya/pocket-mate/internal/pkg/entities"
	mocks "github.com/Taehoya/pocket-mate/internal/pkg/mocks/trip"
	"github.com/stretchr/testify/assert"
)

func TestRegisterTrip(t *testing.T) {
	t.Run("successfully save trip", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		userId := 1

		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		repository.Mock.On("SaveTrip", ctx, name, userId, budget, countryId, description, startDateTime, endDateTime).Return(nil)
		err := usecase.RegisterTrip(ctx, userId, dto)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("failed to save trip", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		userId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		repository.Mock.On("SaveTrip", ctx, name, userId, budget, countryId, description, startDateTime, endDateTime).Return(fmt.Errorf("error"))
		err := usecase.RegisterTrip(ctx, userId, dto)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}

func TestGetTrips(t *testing.T) {
	t.Run("successfully get trips", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		userId := 1

		repository.Mock.On("GetTrip", ctx, userId).Return([]*entities.Trip{}, nil)
		_, err := usecase.GetTrips(ctx, userId)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("failed to get trip", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		userId := 1

		repository.Mock.On("GetTrip", ctx, userId).Return(nil, fmt.Errorf("error"))
		_, err := usecase.GetTrips(ctx, userId)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}

func TestDeleteTrip(t *testing.T) {
	t.Run("successfully delete trip", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		tripId := 1

		repository.Mock.On("DeleteTrip", ctx, tripId).Return(nil)
		err := usecase.DeleteTrip(ctx, tripId)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("failed to delete trip", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		tripId := 1

		repository.Mock.On("DeleteTrip", ctx, tripId).Return(fmt.Errorf("error"))
		err := usecase.DeleteTrip(ctx, tripId)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}

func TestUpdateTrip(t *testing.T) {
	t.Run("successfully update trip", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		tripId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		repository.Mock.On("UpdateTrip", ctx, tripId, name, budget, countryId, description, startDateTime, endDateTime).Return(nil)
		err := usecase.UpdateTrip(ctx, tripId, dto)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("failed to update trip", func(t *testing.T) {
		repository := mocks.NewTripRepositoryMock()
		usecase := NewTripUseCase(repository)

		ctx := context.TODO()
		tripId := 1
		name := "test-name"
		budget := 1000.0
		countryId := 1
		description := "test-description"
		startDateTime := time.Now()
		endDateTime := time.Now()

		dto := dto.TripRequestDTO{
			Name:          name,
			Budget:        budget,
			CountryId:     countryId,
			Description:   description,
			StartDateTime: startDateTime,
			EndDateTime:   endDateTime,
		}

		repository.Mock.On("UpdateTrip", ctx, tripId, name, budget, countryId, description, startDateTime, endDateTime).Return(fmt.Errorf("error"))
		err := usecase.UpdateTrip(ctx, tripId, dto)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
