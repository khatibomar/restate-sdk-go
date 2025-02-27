package mocks

import (
	"github.com/restatedev/sdk-go/internal/restatecontext"
	"testing"

	mock "github.com/stretchr/testify/mock"

	options "github.com/restatedev/sdk-go/internal/options"

	rand "github.com/restatedev/sdk-go/internal/rand"

	slog "log/slog"

	time "time"
)

// MockContext is a mock type for the Context type
type MockContext struct {
	t *testing.T
	mock.Mock
}

type MockContext_Expecter struct {
	parent *MockContext
	mock   *mock.Mock
}

func (_m *MockContext) EXPECT() *MockContext_Expecter {
	return &MockContext_Expecter{parent: _m, mock: &_m.Mock}
}

// After provides a mock function with given fields: _a0
func (_m *MockContext) After(_a0 time.Duration) restatecontext.AfterFuture {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for After")
	}

	var r0 restatecontext.AfterFuture
	if rf, ok := ret.Get(0).(func(time.Duration) restatecontext.AfterFuture); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restatecontext.AfterFuture)
		}
	}

	return r0
}

// MockContext_After_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'After'
type MockContext_After_Call struct {
	*mock.Call
}

// After is a helper method to define mock.On call
//   - _a0 time.Duration
func (_e *MockContext_Expecter) After(_a0 interface{}) *MockContext_After_Call {
	return &MockContext_After_Call{Call: _e.mock.On("After", _a0)}
}

func (_c *MockContext_After_Call) Run(run func(_a0 time.Duration)) *MockContext_After_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockContext_After_Call) Return(_a0 restatecontext.AfterFuture) *MockContext_After_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_After_Call) RunAndReturn(run func(time.Duration) restatecontext.AfterFuture) *MockContext_After_Call {
	_c.Call.Return(run)
	return _c
}

// Awakeable provides a mock function with given fields: _a0
func (_m *MockContext) Awakeable(_a0 ...options.AwakeableOption) restatecontext.AwakeableFuture {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Awakeable")
	}

	var r0 restatecontext.AwakeableFuture
	if rf, ok := ret.Get(0).(func(...options.AwakeableOption) restatecontext.AwakeableFuture); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restatecontext.AwakeableFuture)
		}
	}

	return r0
}

// MockContext_Awakeable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Awakeable'
type MockContext_Awakeable_Call struct {
	*mock.Call
}

// Awakeable is a helper method to define mock.On call
//   - _a0 ...options.AwakeableOption
func (_e *MockContext_Expecter) Awakeable(_a0 ...interface{}) *MockContext_Awakeable_Call {
	return &MockContext_Awakeable_Call{Call: _e.mock.On("Awakeable",
		append([]interface{}{}, _a0...)...)}
}

func (_c *MockContext_Awakeable_Call) Run(run func(_a0 ...options.AwakeableOption)) *MockContext_Awakeable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.AwakeableOption, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(options.AwakeableOption)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Awakeable_Call) Return(_a0 restatecontext.AwakeableFuture) *MockContext_Awakeable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Awakeable_Call) RunAndReturn(run func(...options.AwakeableOption) restatecontext.AwakeableFuture) *MockContext_Awakeable_Call {
	_c.Call.Return(run)
	return _c
}

// Clear provides a mock function with given fields: key
func (_m *MockContext) Clear(key string) {
	_m.Called(key)
}

// MockContext_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type MockContext_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
//   - key string
func (_e *MockContext_Expecter) Clear(key interface{}) *MockContext_Clear_Call {
	return &MockContext_Clear_Call{Call: _e.mock.On("Clear", key)}
}

func (_c *MockContext_Clear_Call) Run(run func(key string)) *MockContext_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockContext_Clear_Call) Return() *MockContext_Clear_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_Clear_Call) RunAndReturn(run func(string)) *MockContext_Clear_Call {
	_c.Run(run)
	return _c
}

// ClearAll provides a mock function with no fields
func (_m *MockContext) ClearAll() {
	_m.Called()
}

// MockContext_ClearAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearAll'
type MockContext_ClearAll_Call struct {
	*mock.Call
}

// ClearAll is a helper method to define mock.On call
func (_e *MockContext_Expecter) ClearAll() *MockContext_ClearAll_Call {
	return &MockContext_ClearAll_Call{Call: _e.mock.On("ClearAll")}
}

func (_c *MockContext_ClearAll_Call) Run(run func()) *MockContext_ClearAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_ClearAll_Call) Return() *MockContext_ClearAll_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_ClearAll_Call) RunAndReturn(run func()) *MockContext_ClearAll_Call {
	_c.Run(run)
	return _c
}

// Deadline provides a mock function with no fields
func (_m *MockContext) Deadline() (time.Time, bool) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Deadline")
	}

	var r0 time.Time
	var r1 bool
	if rf, ok := ret.Get(0).(func() (time.Time, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockContext_Deadline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deadline'
type MockContext_Deadline_Call struct {
	*mock.Call
}

// Deadline is a helper method to define mock.On call
func (_e *MockContext_Expecter) Deadline() *MockContext_Deadline_Call {
	return &MockContext_Deadline_Call{Call: _e.mock.On("Deadline")}
}

func (_c *MockContext_Deadline_Call) Run(run func()) *MockContext_Deadline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Deadline_Call) Return(deadline time.Time, ok bool) *MockContext_Deadline_Call {
	_c.Call.Return(deadline, ok)
	return _c
}

func (_c *MockContext_Deadline_Call) RunAndReturn(run func() (time.Time, bool)) *MockContext_Deadline_Call {
	_c.Call.Return(run)
	return _c
}

// Done provides a mock function with no fields
func (_m *MockContext) Done() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Done")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// MockContext_Done_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Done'
type MockContext_Done_Call struct {
	*mock.Call
}

// Done is a helper method to define mock.On call
func (_e *MockContext_Expecter) Done() *MockContext_Done_Call {
	return &MockContext_Done_Call{Call: _e.mock.On("Done")}
}

func (_c *MockContext_Done_Call) Run(run func()) *MockContext_Done_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Done_Call) Return(_a0 <-chan struct{}) *MockContext_Done_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Done_Call) RunAndReturn(run func() <-chan struct{}) *MockContext_Done_Call {
	_c.Call.Return(run)
	return _c
}

// Err provides a mock function with no fields
func (_m *MockContext) Err() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Err")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_Err_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Err'
type MockContext_Err_Call struct {
	*mock.Call
}

// Err is a helper method to define mock.On call
func (_e *MockContext_Expecter) Err() *MockContext_Err_Call {
	return &MockContext_Err_Call{Call: _e.mock.On("Err")}
}

func (_c *MockContext_Err_Call) Run(run func()) *MockContext_Err_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Err_Call) Return(_a0 error) *MockContext_Err_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Err_Call) RunAndReturn(run func() error) *MockContext_Err_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key, output, _a2
func (_m *MockContext) Get(key string, output any, _a2 ...options.GetOption) (bool, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, output)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, any, ...options.GetOption) (bool, error)); ok {
		return rf(key, output, _a2...)
	}
	if rf, ok := ret.Get(0).(func(string, any, ...options.GetOption) bool); ok {
		r0 = rf(key, output, _a2...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, any, ...options.GetOption) error); ok {
		r1 = rf(key, output, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockContext_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockContext_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
//   - output any
//   - _a2 ...options.GetOption
func (_e *MockContext_Expecter) Get(key interface{}, output interface{}, _a2 ...interface{}) *MockContext_Get_Call {
	return &MockContext_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{key, output}, _a2...)...)}
}

func (_c *MockContext_Get_Call) Run(run func(key string, output any, _a2 ...options.GetOption)) *MockContext_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.GetOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(options.GetOption)
			}
		}
		run(args[0].(string), args[1].(any), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Get_Call) Return(_a0 bool, _a1 error) *MockContext_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockContext_Get_Call) RunAndReturn(run func(string, any, ...options.GetOption) (bool, error)) *MockContext_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Key provides a mock function with no fields
func (_m *MockContext) Key() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Key")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockContext_Key_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Key'
type MockContext_Key_Call struct {
	*mock.Call
}

// Key is a helper method to define mock.On call
func (_e *MockContext_Expecter) Key() *MockContext_Key_Call {
	return &MockContext_Key_Call{Call: _e.mock.On("Key")}
}

func (_c *MockContext_Key_Call) Run(run func()) *MockContext_Key_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Key_Call) Return(_a0 string) *MockContext_Key_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Key_Call) RunAndReturn(run func() string) *MockContext_Key_Call {
	_c.Call.Return(run)
	return _c
}

// Keys provides a mock function with no fields
func (_m *MockContext) Keys() ([]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Keys")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockContext_Keys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Keys'
type MockContext_Keys_Call struct {
	*mock.Call
}

// Keys is a helper method to define mock.On call
func (_e *MockContext_Expecter) Keys() *MockContext_Keys_Call {
	return &MockContext_Keys_Call{Call: _e.mock.On("Keys")}
}

func (_c *MockContext_Keys_Call) Run(run func()) *MockContext_Keys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Keys_Call) Return(_a0 []string, _a1 error) *MockContext_Keys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockContext_Keys_Call) RunAndReturn(run func() ([]string, error)) *MockContext_Keys_Call {
	_c.Call.Return(run)
	return _c
}

// Log provides a mock function with no fields
func (_m *MockContext) Log() *slog.Logger {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Log")
	}

	var r0 *slog.Logger
	if rf, ok := ret.Get(0).(func() *slog.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slog.Logger)
		}
	}

	return r0
}

// MockContext_Log_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Log'
type MockContext_Log_Call struct {
	*mock.Call
}

// Log is a helper method to define mock.On call
func (_e *MockContext_Expecter) Log() *MockContext_Log_Call {
	return &MockContext_Log_Call{Call: _e.mock.On("Log")}
}

func (_c *MockContext_Log_Call) Run(run func()) *MockContext_Log_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Log_Call) Return(_a0 *slog.Logger) *MockContext_Log_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Log_Call) RunAndReturn(run func() *slog.Logger) *MockContext_Log_Call {
	_c.Call.Return(run)
	return _c
}

// Object provides a mock function with given fields: service, key, method, _a3
func (_m *MockContext) Object(service string, key string, method string, _a3 ...options.ClientOption) restatecontext.Client {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, service, key, method)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Object")
	}

	var r0 restatecontext.Client
	if rf, ok := ret.Get(0).(func(string, string, string, ...options.ClientOption) restatecontext.Client); ok {
		r0 = rf(service, key, method, _a3...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restatecontext.Client)
		}
	}

	return r0
}

// MockContext_Object_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Object'
type MockContext_Object_Call struct {
	*mock.Call
}

// Object is a helper method to define mock.On call
//   - service string
//   - key string
//   - method string
//   - _a3 ...options.ClientOption
func (_e *MockContext_Expecter) Object(service interface{}, key interface{}, method interface{}, _a3 ...interface{}) *MockContext_Object_Call {
	return &MockContext_Object_Call{Call: _e.mock.On("Object",
		append([]interface{}{service, key, method}, _a3...)...)}
}

func (_c *MockContext_Object_Call) Run(run func(service string, key string, method string, _a3 ...options.ClientOption)) *MockContext_Object_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.ClientOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(options.ClientOption)
			}
		}
		run(args[0].(string), args[1].(string), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Object_Call) Return(_a0 restatecontext.Client) *MockContext_Object_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Object_Call) RunAndReturn(run func(string, string, string, ...options.ClientOption) restatecontext.Client) *MockContext_Object_Call {
	_c.Call.Return(run)
	return _c
}

// Promise provides a mock function with given fields: name, _a1
func (_m *MockContext) Promise(name string, _a1 ...options.PromiseOption) restatecontext.DurablePromise {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Promise")
	}

	var r0 restatecontext.DurablePromise
	if rf, ok := ret.Get(0).(func(string, ...options.PromiseOption) restatecontext.DurablePromise); ok {
		r0 = rf(name, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restatecontext.DurablePromise)
		}
	}

	return r0
}

// MockContext_Promise_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Promise'
type MockContext_Promise_Call struct {
	*mock.Call
}

// Promise is a helper method to define mock.On call
//   - name string
//   - _a1 ...options.PromiseOption
func (_e *MockContext_Expecter) Promise(name interface{}, _a1 ...interface{}) *MockContext_Promise_Call {
	return &MockContext_Promise_Call{Call: _e.mock.On("Promise",
		append([]interface{}{name}, _a1...)...)}
}

func (_c *MockContext_Promise_Call) Run(run func(name string, _a1 ...options.PromiseOption)) *MockContext_Promise_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.PromiseOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(options.PromiseOption)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Promise_Call) Return(_a0 restatecontext.DurablePromise) *MockContext_Promise_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Promise_Call) RunAndReturn(run func(string, ...options.PromiseOption) restatecontext.DurablePromise) *MockContext_Promise_Call {
	_c.Call.Return(run)
	return _c
}

// Rand provides a mock function with no fields
func (_m *MockContext) Rand() rand.Rand {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rand")
	}

	var r0 rand.Rand
	if rf, ok := ret.Get(0).(func() rand.Rand); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rand.Rand)
		}
	}

	return r0
}

// MockContext_Rand_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rand'
type MockContext_Rand_Call struct {
	*mock.Call
}

// Rand is a helper method to define mock.On call
func (_e *MockContext_Expecter) Rand() *MockContext_Rand_Call {
	return &MockContext_Rand_Call{Call: _e.mock.On("Rand")}
}

func (_c *MockContext_Rand_Call) Run(run func()) *MockContext_Rand_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Rand_Call) Return(_a0 rand.Rand) *MockContext_Rand_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Rand_Call) RunAndReturn(run func() rand.Rand) *MockContext_Rand_Call {
	_c.Call.Return(run)
	return _c
}

// RejectAwakeable provides a mock function with given fields: id, reason
func (_m *MockContext) RejectAwakeable(id string, reason error) {
	_m.Called(id, reason)
}

// MockContext_RejectAwakeable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RejectAwakeable'
type MockContext_RejectAwakeable_Call struct {
	*mock.Call
}

// RejectAwakeable is a helper method to define mock.On call
//   - id string
//   - reason error
func (_e *MockContext_Expecter) RejectAwakeable(id interface{}, reason interface{}) *MockContext_RejectAwakeable_Call {
	return &MockContext_RejectAwakeable_Call{Call: _e.mock.On("RejectAwakeable", id, reason)}
}

func (_c *MockContext_RejectAwakeable_Call) Run(run func(id string, reason error)) *MockContext_RejectAwakeable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(error))
	})
	return _c
}

func (_c *MockContext_RejectAwakeable_Call) Return() *MockContext_RejectAwakeable_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_RejectAwakeable_Call) RunAndReturn(run func(string, error)) *MockContext_RejectAwakeable_Call {
	_c.Run(run)
	return _c
}

// Request provides a mock function with no fields
func (_m *MockContext) Request() *restatecontext.Request {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Request")
	}

	var r0 *restatecontext.Request
	if rf, ok := ret.Get(0).(func() *restatecontext.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*restatecontext.Request)
		}
	}

	return r0
}

// MockContext_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type MockContext_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
func (_e *MockContext_Expecter) Request() *MockContext_Request_Call {
	return &MockContext_Request_Call{Call: _e.mock.On("Request")}
}

func (_c *MockContext_Request_Call) Run(run func()) *MockContext_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Request_Call) Return(_a0 *restatecontext.Request) *MockContext_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Request_Call) RunAndReturn(run func() *restatecontext.Request) *MockContext_Request_Call {
	_c.Call.Return(run)
	return _c
}

// ResolveAwakeable provides a mock function with given fields: id, value, _a2
func (_m *MockContext) ResolveAwakeable(id string, value any, _a2 ...options.ResolveAwakeableOption) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, value)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockContext_ResolveAwakeable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResolveAwakeable'
type MockContext_ResolveAwakeable_Call struct {
	*mock.Call
}

// ResolveAwakeable is a helper method to define mock.On call
//   - id string
//   - value any
//   - _a2 ...options.ResolveAwakeableOption
func (_e *MockContext_Expecter) ResolveAwakeable(id interface{}, value interface{}, _a2 ...interface{}) *MockContext_ResolveAwakeable_Call {
	return &MockContext_ResolveAwakeable_Call{Call: _e.mock.On("ResolveAwakeable",
		append([]interface{}{id, value}, _a2...)...)}
}

func (_c *MockContext_ResolveAwakeable_Call) Run(run func(id string, value any, _a2 ...options.ResolveAwakeableOption)) *MockContext_ResolveAwakeable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.ResolveAwakeableOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(options.ResolveAwakeableOption)
			}
		}
		run(args[0].(string), args[1].(any), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_ResolveAwakeable_Call) Return() *MockContext_ResolveAwakeable_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_ResolveAwakeable_Call) RunAndReturn(run func(string, any, ...options.ResolveAwakeableOption)) *MockContext_ResolveAwakeable_Call {
	_c.Run(run)
	return _c
}

// Run provides a mock function with given fields: fn, output, _a2
func (_m *MockContext) Run(fn func(restatecontext.RunContext) (any, error), output any, _a2 ...options.RunOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fn, output)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(restatecontext.RunContext) (any, error), any, ...options.RunOption) error); ok {
		r0 = rf(fn, output, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockContext_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - fn func(restatecontext.RunContext)(any , error)
//   - output any
//   - _a2 ...options.RunOption
func (_e *MockContext_Expecter) Run(fn interface{}, output interface{}, _a2 ...interface{}) *MockContext_Run_Call {
	return &MockContext_Run_Call{Call: _e.mock.On("Run",
		append([]interface{}{fn, output}, _a2...)...)}
}

func (_c *MockContext_Run_Call) Run(run func(fn func(restatecontext.RunContext) (any, error), output any, _a2 ...options.RunOption)) *MockContext_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.RunOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(options.RunOption)
			}
		}
		run(args[0].(func(restatecontext.RunContext) (any, error)), args[1].(any), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Run_Call) Return(_a0 error) *MockContext_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Run_Call) RunAndReturn(run func(func(restatecontext.RunContext) (any, error), any, ...options.RunOption) error) *MockContext_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Select provides a mock function with given fields: futs
func (_m *MockContext) Select(futs ...restatecontext.Selectable) restatecontext.Selector {
	_va := make([]interface{}, len(futs))
	for _i := range futs {
		_va[_i] = futs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Select")
	}

	var r0 restatecontext.Selector
	if rf, ok := ret.Get(0).(func(...restatecontext.Selectable) restatecontext.Selector); ok {
		r0 = rf(futs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restatecontext.Selector)
		}
	}

	return r0
}

// MockContext_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type MockContext_Select_Call struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//   - futs ...restatecontext.Selectable
func (_e *MockContext_Expecter) Select(futs ...interface{}) *MockContext_Select_Call {
	return &MockContext_Select_Call{Call: _e.mock.On("Select",
		append([]interface{}{}, futs...)...)}
}

func (_c *MockContext_Select_Call) Run(run func(futs ...restatecontext.Selectable)) *MockContext_Select_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]restatecontext.Selectable, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(restatecontext.Selectable)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Select_Call) Return(_a0 restatecontext.Selector) *MockContext_Select_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Select_Call) RunAndReturn(run func(...restatecontext.Selectable) restatecontext.Selector) *MockContext_Select_Call {
	_c.Call.Return(run)
	return _c
}

// Service provides a mock function with given fields: service, method, _a2
func (_m *MockContext) Service(service string, method string, _a2 ...options.ClientOption) restatecontext.Client {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, service, method)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Service")
	}

	var r0 restatecontext.Client
	if rf, ok := ret.Get(0).(func(string, string, ...options.ClientOption) restatecontext.Client); ok {
		r0 = rf(service, method, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restatecontext.Client)
		}
	}

	return r0
}

// MockContext_Service_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Service'
type MockContext_Service_Call struct {
	*mock.Call
}

// Service is a helper method to define mock.On call
//   - service string
//   - method string
//   - _a2 ...options.ClientOption
func (_e *MockContext_Expecter) Service(service interface{}, method interface{}, _a2 ...interface{}) *MockContext_Service_Call {
	return &MockContext_Service_Call{Call: _e.mock.On("Service",
		append([]interface{}{service, method}, _a2...)...)}
}

func (_c *MockContext_Service_Call) Run(run func(service string, method string, _a2 ...options.ClientOption)) *MockContext_Service_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.ClientOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(options.ClientOption)
			}
		}
		run(args[0].(string), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Service_Call) Return(_a0 restatecontext.Client) *MockContext_Service_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Service_Call) RunAndReturn(run func(string, string, ...options.ClientOption) restatecontext.Client) *MockContext_Service_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value, _a2
func (_m *MockContext) Set(key string, value any, _a2 ...options.SetOption) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, value)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockContext_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockContext_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - value any
//   - _a2 ...options.SetOption
func (_e *MockContext_Expecter) Set(key interface{}, value interface{}, _a2 ...interface{}) *MockContext_Set_Call {
	return &MockContext_Set_Call{Call: _e.mock.On("Set",
		append([]interface{}{key, value}, _a2...)...)}
}

func (_c *MockContext_Set_Call) Run(run func(key string, value any, _a2 ...options.SetOption)) *MockContext_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.SetOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(options.SetOption)
			}
		}
		run(args[0].(string), args[1].(any), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Set_Call) Return() *MockContext_Set_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_Set_Call) RunAndReturn(run func(string, any, ...options.SetOption)) *MockContext_Set_Call {
	_c.Run(run)
	return _c
}

// Sleep provides a mock function with given fields: _a0
func (_m *MockContext) Sleep(_a0 time.Duration) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Sleep")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_Sleep_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sleep'
type MockContext_Sleep_Call struct {
	*mock.Call
}

// Sleep is a helper method to define mock.On call
//   - _a0 time.Duration
func (_e *MockContext_Expecter) Sleep(_a0 interface{}) *MockContext_Sleep_Call {
	return &MockContext_Sleep_Call{Call: _e.mock.On("Sleep", _a0)}
}

func (_c *MockContext_Sleep_Call) Run(run func(_a0 time.Duration)) *MockContext_Sleep_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockContext_Sleep_Call) Return(_a0 error) *MockContext_Sleep_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Sleep_Call) RunAndReturn(run func(time.Duration) error) *MockContext_Sleep_Call {
	_c.Call.Return(run)
	return _c
}

// Value provides a mock function with given fields: key
func (_m *MockContext) Value(key any) any {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Value")
	}

	var r0 any
	if rf, ok := ret.Get(0).(func(any) any); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	return r0
}

// MockContext_Value_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Value'
type MockContext_Value_Call struct {
	*mock.Call
}

// Value is a helper method to define mock.On call
//   - key any
func (_e *MockContext_Expecter) Value(key interface{}) *MockContext_Value_Call {
	return &MockContext_Value_Call{Call: _e.mock.On("Value", key)}
}

func (_c *MockContext_Value_Call) Run(run func(key any)) *MockContext_Value_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(any))
	})
	return _c
}

func (_c *MockContext_Value_Call) Return(_a0 any) *MockContext_Value_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Value_Call) RunAndReturn(run func(any) any) *MockContext_Value_Call {
	_c.Call.Return(run)
	return _c
}

// Workflow provides a mock function with given fields: seservice, workflowID, method, _a3
func (_m *MockContext) Workflow(seservice string, workflowID string, method string, _a3 ...options.ClientOption) restatecontext.Client {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, seservice, workflowID, method)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Workflow")
	}

	var r0 restatecontext.Client
	if rf, ok := ret.Get(0).(func(string, string, string, ...options.ClientOption) restatecontext.Client); ok {
		r0 = rf(seservice, workflowID, method, _a3...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(restatecontext.Client)
		}
	}

	return r0
}

// MockContext_Workflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Workflow'
type MockContext_Workflow_Call struct {
	*mock.Call
}

// Workflow is a helper method to define mock.On call
//   - seservice string
//   - workflowID string
//   - method string
//   - _a3 ...options.ClientOption
func (_e *MockContext_Expecter) Workflow(seservice interface{}, workflowID interface{}, method interface{}, _a3 ...interface{}) *MockContext_Workflow_Call {
	return &MockContext_Workflow_Call{Call: _e.mock.On("Workflow",
		append([]interface{}{seservice, workflowID, method}, _a3...)...)}
}

func (_c *MockContext_Workflow_Call) Run(run func(seservice string, workflowID string, method string, _a3 ...options.ClientOption)) *MockContext_Workflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.ClientOption, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(options.ClientOption)
			}
		}
		run(args[0].(string), args[1].(string), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockContext_Workflow_Call) Return(_a0 restatecontext.Client) *MockContext_Workflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Workflow_Call) RunAndReturn(run func(string, string, string, ...options.ClientOption) restatecontext.Client) *MockContext_Workflow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockContext creates a new instance of MockContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockContext(t *testing.T) *MockContext {
	mock := &MockContext{t: t}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
