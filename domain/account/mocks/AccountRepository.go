// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/sangianpatrick/devoria-article-service/domain/account/entity"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *AccountRepository) FindByEmail(ctx context.Context, email string) (entity.Account, error) {
	ret := _m.Called(ctx, email)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Account); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, ID
func (_m *AccountRepository) FindByID(ctx context.Context, ID int64) (entity.Account, error) {
	ret := _m.Called(ctx, ID)

	var r0 entity.Account
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Account); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *AccountRepository) Save(ctx context.Context, _a1 entity.Account) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, entity.Account) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Account) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ID, updatedAccount
func (_m *AccountRepository) Update(ctx context.Context, ID int64, updatedAccount entity.Account) error {
	ret := _m.Called(ctx, ID, updatedAccount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, entity.Account) error); ok {
		r0 = rf(ctx, ID, updatedAccount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
