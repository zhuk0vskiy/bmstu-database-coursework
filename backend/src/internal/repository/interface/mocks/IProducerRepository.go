// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	dto "app/src/internal/model/dto"
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "app/src/internal/model"
)

// IProducerRepository is an autogenerated mock type for the IProducerRepository type
type IProducerRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, request
func (_m *IProducerRepository) Add(ctx context.Context, request *dto.AddProducerRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.AddProducerRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, request
func (_m *IProducerRepository) Delete(ctx context.Context, request *dto.DeleteProducerRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.DeleteProducerRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, request
func (_m *IProducerRepository) Get(ctx context.Context, request *dto.GetProducerRequest) (*model.Producer, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Producer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetProducerRequest) (*model.Producer, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetProducerRequest) *model.Producer); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Producer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetProducerRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByStudio provides a mock function with given fields: ctx, request
func (_m *IProducerRepository) GetByStudio(ctx context.Context, request *dto.GetProducerByStudioRequest) ([]*model.Producer, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetByStudio")
	}

	var r0 []*model.Producer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetProducerByStudioRequest) ([]*model.Producer, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetProducerByStudioRequest) []*model.Producer); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Producer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetProducerByStudioRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, request
func (_m *IProducerRepository) Update(ctx context.Context, request *dto.UpdateProducerRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.UpdateProducerRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIProducerRepository creates a new instance of IProducerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProducerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProducerRepository {
	mock := &IProducerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}