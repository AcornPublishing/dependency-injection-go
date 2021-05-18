// Code generated by mockery v1.0.0

// @generated

package list

import (
	"context"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/modules/data"
	"github.com/stretchr/testify/mock"
)

// mockMyLoader is an autogenerated mock type for the myLoader type
type mockMyLoader struct {
	mock.Mock
}

// LoadAll provides a mock function with given fields: ctx
func (_m *mockMyLoader) LoadAll(ctx context.Context) ([]*data.Person, error) {
	ret := _m.Called(ctx)

	var r0 []*data.Person
	if rf, ok := ret.Get(0).(func(context.Context) []*data.Person); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
