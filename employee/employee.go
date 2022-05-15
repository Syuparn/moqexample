package employee

import "context"

//go:generate moq -out employee_moq_test.go . EmployeeRepository

type EmployeeRepository interface {
	Get(ctx context.Context, id EmployeeID) (*Employee, error)
}

type EmployeeID uint
type EmployeeName string

type Employee struct {
	ID   EmployeeID
	Name EmployeeName
}
