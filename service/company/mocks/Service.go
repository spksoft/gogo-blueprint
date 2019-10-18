// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import company "github.com/devit-tel/gogo-blueprint/service/company"
import context "context"
import domaincompany "github.com/devit-tel/gogo-blueprint/domain/company"
import goerror "github.com/devit-tel/goerror"
import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateCompany provides a mock function with given fields: ctx, input
func (_m *Service) CreateCompany(ctx context.Context, input *company.CreateCompanyInput) (*domaincompany.Company, goerror.Error) {
	ret := _m.Called(ctx, input)

	var r0 *domaincompany.Company
	if rf, ok := ret.Get(0).(func(context.Context, *company.CreateCompanyInput) *domaincompany.Company); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domaincompany.Company)
		}
	}

	var r1 goerror.Error
	if rf, ok := ret.Get(1).(func(context.Context, *company.CreateCompanyInput) goerror.Error); ok {
		r1 = rf(ctx, input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(goerror.Error)
		}
	}

	return r0, r1
}
