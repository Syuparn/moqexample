package workflow

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syuparn/moqexample/employee"
)

func TestCreateEmployeeUsecase(t *testing.T) {
	tests := []struct {
		name                 string
		in                   *CreateEmployeeInputData
		mockRegisterEmployee func(context.Context, *employee.Employee) (*Workflow, error)
		mockGetWorkflow      func(context.Context, WorkflowID) (*Workflow, error)
		expected             *CreateEmployeeOutputData
	}{
		{
			"create employee Taro",
			&CreateEmployeeInputData{
				Name: "Taro",
			},
			func(_ context.Context, _ *employee.Employee) (*Workflow, error) {
				return &Workflow{
					ID:       WorkflowID(5678),
					Progress: 0,
				}, nil
			},
			func() func(context.Context, WorkflowID) (*Workflow, error) {
				i := 0
				return func(_ context.Context, _ WorkflowID) (*Workflow, error) {
					wf := []*Workflow{
						{ID: 5678, Progress: 20},
						{ID: 5678, Progress: 40},
						{ID: 5678, Progress: 60},
						{ID: 5678, Progress: 80},
						{ID: 5678, Progress: 100},
					}[i]
					i++
					return wf, nil
				}
			}(),
			&CreateEmployeeOutputData{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wfSvc := &EmployeeWorkflowServiceMock{
				RegisterFunc: tt.mockRegisterEmployee,
			}
			wfRepo := &WorkflowRepositoryMock{
				GetFunc: tt.mockGetWorkflow,
			}

			u := &CreateEmployeeUsecase{
				EmployeeWorkflowService: wfSvc,
				WorkflowRepository:      wfRepo,
			}

			actual, err := u.Exec(context.TODO(), tt.in)

			assert.Equal(t, tt.expected, actual)
			assert.Nil(t, err)
		})
	}
}
