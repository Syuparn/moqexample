// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package workflow

import (
	"context"
	"github.com/syuparn/moqexample/employee"
	"sync"
)

// Ensure, that EmployeeWorkflowServiceMock does implement EmployeeWorkflowService.
// If this is not the case, regenerate this file with moq.
var _ EmployeeWorkflowService = &EmployeeWorkflowServiceMock{}

// EmployeeWorkflowServiceMock is a mock implementation of EmployeeWorkflowService.
//
// 	func TestSomethingThatUsesEmployeeWorkflowService(t *testing.T) {
//
// 		// make and configure a mocked EmployeeWorkflowService
// 		mockedEmployeeWorkflowService := &EmployeeWorkflowServiceMock{
// 			RegisterFunc: func(ctx context.Context, employeeMoqParam *employee.Employee) (*Workflow, error) {
// 				panic("mock out the Register method")
// 			},
// 		}
//
// 		// use mockedEmployeeWorkflowService in code that requires EmployeeWorkflowService
// 		// and then make assertions.
//
// 	}
type EmployeeWorkflowServiceMock struct {
	// RegisterFunc mocks the Register method.
	RegisterFunc func(ctx context.Context, employeeMoqParam *employee.Employee) (*Workflow, error)

	// calls tracks calls to the methods.
	calls struct {
		// Register holds details about calls to the Register method.
		Register []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EmployeeMoqParam is the employeeMoqParam argument value.
			EmployeeMoqParam *employee.Employee
		}
	}
	lockRegister sync.RWMutex
}

// Register calls RegisterFunc.
func (mock *EmployeeWorkflowServiceMock) Register(ctx context.Context, employeeMoqParam *employee.Employee) (*Workflow, error) {
	if mock.RegisterFunc == nil {
		panic("EmployeeWorkflowServiceMock.RegisterFunc: method is nil but EmployeeWorkflowService.Register was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		EmployeeMoqParam *employee.Employee
	}{
		Ctx:              ctx,
		EmployeeMoqParam: employeeMoqParam,
	}
	mock.lockRegister.Lock()
	mock.calls.Register = append(mock.calls.Register, callInfo)
	mock.lockRegister.Unlock()
	return mock.RegisterFunc(ctx, employeeMoqParam)
}

// RegisterCalls gets all the calls that were made to Register.
// Check the length with:
//     len(mockedEmployeeWorkflowService.RegisterCalls())
func (mock *EmployeeWorkflowServiceMock) RegisterCalls() []struct {
	Ctx              context.Context
	EmployeeMoqParam *employee.Employee
} {
	var calls []struct {
		Ctx              context.Context
		EmployeeMoqParam *employee.Employee
	}
	mock.lockRegister.RLock()
	calls = mock.calls.Register
	mock.lockRegister.RUnlock()
	return calls
}