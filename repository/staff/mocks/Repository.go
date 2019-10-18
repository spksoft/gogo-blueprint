// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import goerror "github.com/devit-tel/goerror"
import mock "github.com/stretchr/testify/mock"

import staff "github.com/devit-tel/gogo-blueprint/domain/staff"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, staffId
func (_m *Repository) Get(ctx context.Context, staffId string) (*staff.Staff, goerror.Error) {
	ret := _m.Called(ctx, staffId)

	var r0 *staff.Staff
	if rf, ok := ret.Get(0).(func(context.Context, string) *staff.Staff); ok {
		r0 = rf(ctx, staffId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*staff.Staff)
		}
	}

	var r1 goerror.Error
	if rf, ok := ret.Get(1).(func(context.Context, string) goerror.Error); ok {
		r1 = rf(ctx, staffId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(goerror.Error)
		}
	}

	return r0, r1
}

// GetStaffsByCompany provides a mock function with given fields: ctx, companyId, offset, limit
func (_m *Repository) GetStaffsByCompany(ctx context.Context, companyId string, offset int64, limit int64) ([]*staff.Staff, goerror.Error) {
	ret := _m.Called(ctx, companyId, offset, limit)

	var r0 []*staff.Staff
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) []*staff.Staff); ok {
		r0 = rf(ctx, companyId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*staff.Staff)
		}
	}

	var r1 goerror.Error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) goerror.Error); ok {
		r1 = rf(ctx, companyId, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(goerror.Error)
		}
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *Repository) Save(ctx context.Context, _a1 *staff.Staff) goerror.Error {
	ret := _m.Called(ctx, _a1)

	var r0 goerror.Error
	if rf, ok := ret.Get(0).(func(context.Context, *staff.Staff) goerror.Error); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(goerror.Error)
		}
	}

	return r0
}