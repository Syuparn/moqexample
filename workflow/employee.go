package workflow

import (
	"context"

	"github.com/syuparn/moqexample/employee"
)

//go:generate moq -fmt goimports -out employee_moq_test.go . EmployeeWorkflowService

type EmployeeWorkflowService interface {
	Register(ctx context.Context, employee *employee.Employee) (*Workflow, error)
}
