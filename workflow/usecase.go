package workflow

import (
	"context"
	"fmt"
	"time"

	"github.com/syuparn/moqexample/employee"
)

type CreateEmployeeInputData struct {
	Name string
}

type CreateEmployeeOutputData struct{}

type CreateEmployeeUsecase struct {
	EmployeeWorkflowService EmployeeWorkflowService
	WorkflowRepository      WorkflowRepository
}

func (u *CreateEmployeeUsecase) Exec(
	ctx context.Context,
	in *CreateEmployeeInputData,
) (*CreateEmployeeOutputData, error) {
	employee := &employee.Employee{
		ID:   1234, // TODO: generate unique number
		Name: employee.EmployeeName(in.Name),
	}
	wf, err := u.EmployeeWorkflowService.Register(ctx, employee)
	if err != nil {
		return nil, fmt.Errorf("failed to register employee: %w", err)
	}

	for i := 0; i < 10; i++ {
		updatedWF, err := u.WorkflowRepository.Get(ctx, wf.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get workflow: %w", err)
		}

		if updatedWF.Finished() {
			return &CreateEmployeeOutputData{}, nil
		}

		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("workflow timeout: %w", err)
}
