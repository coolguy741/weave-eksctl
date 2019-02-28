// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import cloudtrail "github.com/aws/aws-sdk-go/service/cloudtrail"

import context "context"
import mock "github.com/stretchr/testify/mock"
import request "github.com/aws/aws-sdk-go/aws/request"

// CloudTrailAPI is an autogenerated mock type for the CloudTrailAPI type
type CloudTrailAPI struct {
	mock.Mock
}

// AddTags provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) AddTags(_a0 *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.AddTagsOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.AddTagsInput) *cloudtrail.AddTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.AddTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.AddTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddTagsRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) AddTagsRequest(_a0 *cloudtrail.AddTagsInput) (*request.Request, *cloudtrail.AddTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.AddTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.AddTagsOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.AddTagsInput) *cloudtrail.AddTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.AddTagsOutput)
		}
	}

	return r0, r1
}

// AddTagsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) AddTagsWithContext(_a0 context.Context, _a1 *cloudtrail.AddTagsInput, _a2 ...request.Option) (*cloudtrail.AddTagsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.AddTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.AddTagsInput, ...request.Option) *cloudtrail.AddTagsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.AddTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.AddTagsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTrail provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) CreateTrail(_a0 *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.CreateTrailOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.CreateTrailInput) *cloudtrail.CreateTrailOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.CreateTrailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.CreateTrailInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTrailRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) CreateTrailRequest(_a0 *cloudtrail.CreateTrailInput) (*request.Request, *cloudtrail.CreateTrailOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.CreateTrailInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.CreateTrailOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.CreateTrailInput) *cloudtrail.CreateTrailOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.CreateTrailOutput)
		}
	}

	return r0, r1
}

// CreateTrailWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) CreateTrailWithContext(_a0 context.Context, _a1 *cloudtrail.CreateTrailInput, _a2 ...request.Option) (*cloudtrail.CreateTrailOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.CreateTrailOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.CreateTrailInput, ...request.Option) *cloudtrail.CreateTrailOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.CreateTrailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.CreateTrailInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTrail provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) DeleteTrail(_a0 *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.DeleteTrailOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.DeleteTrailInput) *cloudtrail.DeleteTrailOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.DeleteTrailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.DeleteTrailInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTrailRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) DeleteTrailRequest(_a0 *cloudtrail.DeleteTrailInput) (*request.Request, *cloudtrail.DeleteTrailOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.DeleteTrailInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.DeleteTrailOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.DeleteTrailInput) *cloudtrail.DeleteTrailOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.DeleteTrailOutput)
		}
	}

	return r0, r1
}

// DeleteTrailWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) DeleteTrailWithContext(_a0 context.Context, _a1 *cloudtrail.DeleteTrailInput, _a2 ...request.Option) (*cloudtrail.DeleteTrailOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.DeleteTrailOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.DeleteTrailInput, ...request.Option) *cloudtrail.DeleteTrailOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.DeleteTrailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.DeleteTrailInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTrails provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) DescribeTrails(_a0 *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.DescribeTrailsOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.DescribeTrailsInput) *cloudtrail.DescribeTrailsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.DescribeTrailsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.DescribeTrailsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTrailsRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) DescribeTrailsRequest(_a0 *cloudtrail.DescribeTrailsInput) (*request.Request, *cloudtrail.DescribeTrailsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.DescribeTrailsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.DescribeTrailsOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.DescribeTrailsInput) *cloudtrail.DescribeTrailsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.DescribeTrailsOutput)
		}
	}

	return r0, r1
}

// DescribeTrailsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) DescribeTrailsWithContext(_a0 context.Context, _a1 *cloudtrail.DescribeTrailsInput, _a2 ...request.Option) (*cloudtrail.DescribeTrailsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.DescribeTrailsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.DescribeTrailsInput, ...request.Option) *cloudtrail.DescribeTrailsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.DescribeTrailsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.DescribeTrailsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventSelectors provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) GetEventSelectors(_a0 *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.GetEventSelectorsOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.GetEventSelectorsInput) *cloudtrail.GetEventSelectorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.GetEventSelectorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.GetEventSelectorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventSelectorsRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) GetEventSelectorsRequest(_a0 *cloudtrail.GetEventSelectorsInput) (*request.Request, *cloudtrail.GetEventSelectorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.GetEventSelectorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.GetEventSelectorsOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.GetEventSelectorsInput) *cloudtrail.GetEventSelectorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.GetEventSelectorsOutput)
		}
	}

	return r0, r1
}

// GetEventSelectorsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) GetEventSelectorsWithContext(_a0 context.Context, _a1 *cloudtrail.GetEventSelectorsInput, _a2 ...request.Option) (*cloudtrail.GetEventSelectorsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.GetEventSelectorsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.GetEventSelectorsInput, ...request.Option) *cloudtrail.GetEventSelectorsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.GetEventSelectorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.GetEventSelectorsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrailStatus provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) GetTrailStatus(_a0 *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.GetTrailStatusOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.GetTrailStatusInput) *cloudtrail.GetTrailStatusOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.GetTrailStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.GetTrailStatusInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrailStatusRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) GetTrailStatusRequest(_a0 *cloudtrail.GetTrailStatusInput) (*request.Request, *cloudtrail.GetTrailStatusOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.GetTrailStatusInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.GetTrailStatusOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.GetTrailStatusInput) *cloudtrail.GetTrailStatusOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.GetTrailStatusOutput)
		}
	}

	return r0, r1
}

// GetTrailStatusWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) GetTrailStatusWithContext(_a0 context.Context, _a1 *cloudtrail.GetTrailStatusInput, _a2 ...request.Option) (*cloudtrail.GetTrailStatusOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.GetTrailStatusOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.GetTrailStatusInput, ...request.Option) *cloudtrail.GetTrailStatusOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.GetTrailStatusOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.GetTrailStatusInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPublicKeys provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) ListPublicKeys(_a0 *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.ListPublicKeysOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.ListPublicKeysInput) *cloudtrail.ListPublicKeysOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.ListPublicKeysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.ListPublicKeysInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPublicKeysRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) ListPublicKeysRequest(_a0 *cloudtrail.ListPublicKeysInput) (*request.Request, *cloudtrail.ListPublicKeysOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.ListPublicKeysInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.ListPublicKeysOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.ListPublicKeysInput) *cloudtrail.ListPublicKeysOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.ListPublicKeysOutput)
		}
	}

	return r0, r1
}

// ListPublicKeysWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) ListPublicKeysWithContext(_a0 context.Context, _a1 *cloudtrail.ListPublicKeysInput, _a2 ...request.Option) (*cloudtrail.ListPublicKeysOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.ListPublicKeysOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.ListPublicKeysInput, ...request.Option) *cloudtrail.ListPublicKeysOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.ListPublicKeysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.ListPublicKeysInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTags provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) ListTags(_a0 *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.ListTagsOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.ListTagsInput) *cloudtrail.ListTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.ListTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.ListTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) ListTagsRequest(_a0 *cloudtrail.ListTagsInput) (*request.Request, *cloudtrail.ListTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.ListTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.ListTagsOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.ListTagsInput) *cloudtrail.ListTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.ListTagsOutput)
		}
	}

	return r0, r1
}

// ListTagsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) ListTagsWithContext(_a0 context.Context, _a1 *cloudtrail.ListTagsInput, _a2 ...request.Option) (*cloudtrail.ListTagsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.ListTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.ListTagsInput, ...request.Option) *cloudtrail.ListTagsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.ListTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.ListTagsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupEvents provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) LookupEvents(_a0 *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.LookupEventsOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.LookupEventsInput) *cloudtrail.LookupEventsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.LookupEventsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.LookupEventsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupEventsPages provides a mock function with given fields: _a0, _a1
func (_m *CloudTrailAPI) LookupEventsPages(_a0 *cloudtrail.LookupEventsInput, _a1 func(*cloudtrail.LookupEventsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cloudtrail.LookupEventsInput, func(*cloudtrail.LookupEventsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LookupEventsPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *CloudTrailAPI) LookupEventsPagesWithContext(_a0 context.Context, _a1 *cloudtrail.LookupEventsInput, _a2 func(*cloudtrail.LookupEventsOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.LookupEventsInput, func(*cloudtrail.LookupEventsOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LookupEventsRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) LookupEventsRequest(_a0 *cloudtrail.LookupEventsInput) (*request.Request, *cloudtrail.LookupEventsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.LookupEventsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.LookupEventsOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.LookupEventsInput) *cloudtrail.LookupEventsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.LookupEventsOutput)
		}
	}

	return r0, r1
}

// LookupEventsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) LookupEventsWithContext(_a0 context.Context, _a1 *cloudtrail.LookupEventsInput, _a2 ...request.Option) (*cloudtrail.LookupEventsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.LookupEventsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.LookupEventsInput, ...request.Option) *cloudtrail.LookupEventsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.LookupEventsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.LookupEventsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutEventSelectors provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) PutEventSelectors(_a0 *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.PutEventSelectorsOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.PutEventSelectorsInput) *cloudtrail.PutEventSelectorsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.PutEventSelectorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.PutEventSelectorsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutEventSelectorsRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) PutEventSelectorsRequest(_a0 *cloudtrail.PutEventSelectorsInput) (*request.Request, *cloudtrail.PutEventSelectorsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.PutEventSelectorsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.PutEventSelectorsOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.PutEventSelectorsInput) *cloudtrail.PutEventSelectorsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.PutEventSelectorsOutput)
		}
	}

	return r0, r1
}

// PutEventSelectorsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) PutEventSelectorsWithContext(_a0 context.Context, _a1 *cloudtrail.PutEventSelectorsInput, _a2 ...request.Option) (*cloudtrail.PutEventSelectorsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.PutEventSelectorsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.PutEventSelectorsInput, ...request.Option) *cloudtrail.PutEventSelectorsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.PutEventSelectorsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.PutEventSelectorsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTags provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) RemoveTags(_a0 *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.RemoveTagsOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.RemoveTagsInput) *cloudtrail.RemoveTagsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.RemoveTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.RemoveTagsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTagsRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) RemoveTagsRequest(_a0 *cloudtrail.RemoveTagsInput) (*request.Request, *cloudtrail.RemoveTagsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.RemoveTagsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.RemoveTagsOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.RemoveTagsInput) *cloudtrail.RemoveTagsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.RemoveTagsOutput)
		}
	}

	return r0, r1
}

// RemoveTagsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) RemoveTagsWithContext(_a0 context.Context, _a1 *cloudtrail.RemoveTagsInput, _a2 ...request.Option) (*cloudtrail.RemoveTagsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.RemoveTagsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.RemoveTagsInput, ...request.Option) *cloudtrail.RemoveTagsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.RemoveTagsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.RemoveTagsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartLogging provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) StartLogging(_a0 *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.StartLoggingOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.StartLoggingInput) *cloudtrail.StartLoggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.StartLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.StartLoggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartLoggingRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) StartLoggingRequest(_a0 *cloudtrail.StartLoggingInput) (*request.Request, *cloudtrail.StartLoggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.StartLoggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.StartLoggingOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.StartLoggingInput) *cloudtrail.StartLoggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.StartLoggingOutput)
		}
	}

	return r0, r1
}

// StartLoggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) StartLoggingWithContext(_a0 context.Context, _a1 *cloudtrail.StartLoggingInput, _a2 ...request.Option) (*cloudtrail.StartLoggingOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.StartLoggingOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.StartLoggingInput, ...request.Option) *cloudtrail.StartLoggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.StartLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.StartLoggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopLogging provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) StopLogging(_a0 *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.StopLoggingOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.StopLoggingInput) *cloudtrail.StopLoggingOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.StopLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.StopLoggingInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopLoggingRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) StopLoggingRequest(_a0 *cloudtrail.StopLoggingInput) (*request.Request, *cloudtrail.StopLoggingOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.StopLoggingInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.StopLoggingOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.StopLoggingInput) *cloudtrail.StopLoggingOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.StopLoggingOutput)
		}
	}

	return r0, r1
}

// StopLoggingWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) StopLoggingWithContext(_a0 context.Context, _a1 *cloudtrail.StopLoggingInput, _a2 ...request.Option) (*cloudtrail.StopLoggingOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.StopLoggingOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.StopLoggingInput, ...request.Option) *cloudtrail.StopLoggingOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.StopLoggingOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.StopLoggingInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTrail provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) UpdateTrail(_a0 *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error) {
	ret := _m.Called(_a0)

	var r0 *cloudtrail.UpdateTrailOutput
	if rf, ok := ret.Get(0).(func(*cloudtrail.UpdateTrailInput) *cloudtrail.UpdateTrailOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.UpdateTrailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cloudtrail.UpdateTrailInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTrailRequest provides a mock function with given fields: _a0
func (_m *CloudTrailAPI) UpdateTrailRequest(_a0 *cloudtrail.UpdateTrailInput) (*request.Request, *cloudtrail.UpdateTrailOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*cloudtrail.UpdateTrailInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *cloudtrail.UpdateTrailOutput
	if rf, ok := ret.Get(1).(func(*cloudtrail.UpdateTrailInput) *cloudtrail.UpdateTrailOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*cloudtrail.UpdateTrailOutput)
		}
	}

	return r0, r1
}

// UpdateTrailWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudTrailAPI) UpdateTrailWithContext(_a0 context.Context, _a1 *cloudtrail.UpdateTrailInput, _a2 ...request.Option) (*cloudtrail.UpdateTrailOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *cloudtrail.UpdateTrailOutput
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrail.UpdateTrailInput, ...request.Option) *cloudtrail.UpdateTrailOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudtrail.UpdateTrailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrail.UpdateTrailInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
