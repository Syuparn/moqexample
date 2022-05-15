package employee

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployeeUsecase(t *testing.T) {
	tests := []struct {
		name            string
		in              *GetEmployeeInputData
		mockGetEmployee func(context.Context, EmployeeID) (*Employee, error)
		expected        *GetEmployeeOutputData
	}{
		{
			"get employee Taro",
			&GetEmployeeInputData{
				EmployeeID: 1234,
			},
			func(_ context.Context, id EmployeeID) (*Employee, error) {
				return &Employee{
					ID:   EmployeeID(1234),
					Name: EmployeeName("Taro"),
				}, nil
			},
			&GetEmployeeOutputData{
				&Employee{
					ID:   EmployeeID(1234),
					Name: EmployeeName("Taro"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &EmployeeRepositoryMock{
				GetFunc: tt.mockGetEmployee,
			}
			u := &GetEmployeeUsecase{
				EmployeeRepository: repo,
			}

			actual, err := u.Exec(context.TODO(), tt.in)

			assert.Equal(t, tt.expected, actual)
			assert.Nil(t, err)
		})
	}
}

func TestGetEmployeeUsecaseError(t *testing.T) {
	tests := []struct {
		name            string
		in              *GetEmployeeInputData
		mockGetEmployee func(context.Context, EmployeeID) (*Employee, error)
	}{
		{
			"employee not found",
			&GetEmployeeInputData{
				EmployeeID: 404,
			},
			func(_ context.Context, _ EmployeeID) (*Employee, error) {
				return nil, errors.New("employee Not Found")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &EmployeeRepositoryMock{
				GetFunc: tt.mockGetEmployee,
			}
			u := &GetEmployeeUsecase{
				EmployeeRepository: repo,
			}

			_, err := u.Exec(context.TODO(), tt.in)
			assert.Error(t, err)
		})
	}
}
