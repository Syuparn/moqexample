package employee

import (
	"context"
	"fmt"
)

type GetEmployeeInputData struct {
	EmployeeID uint
}

type GetEmployeeOutputData struct {
	Employee *Employee
}

type GetEmployeeUsecase struct {
	EmployeeRepository EmployeeRepository
}

func (u *GetEmployeeUsecase) Exec(
	ctx context.Context,
	in *GetEmployeeInputData,
) (*GetEmployeeOutputData, error) {
	id := EmployeeID(in.EmployeeID)
	employee, err := u.EmployeeRepository.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee: %w", err)
	}

	return &GetEmployeeOutputData{
		Employee: employee,
	}, nil
}
