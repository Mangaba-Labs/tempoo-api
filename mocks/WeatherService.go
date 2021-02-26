// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather/model"
	mock "github.com/stretchr/testify/mock"
)

// WeatherService is an autogenerated mock type for the WeatherService type
type WeatherService struct {
	mock.Mock
}

// GetCurrentWeather provides a mock function with given fields: coords
func (_m *WeatherService) GetCurrentWeather(coords *model.Request) (*model.Response, error) {
	ret := _m.Called(coords)

	var r0 *model.Response
	if rf, ok := ret.Get(0).(func(*model.Request) *model.Response); ok {
		r0 = rf(coords)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Request) error); ok {
		r1 = rf(coords)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
