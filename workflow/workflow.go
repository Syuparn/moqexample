package workflow

import "context"

//go:generate moq -out workflow_moq_test.go . WorkflowRepository

type WorkflowRepository interface {
	Get(ctx context.Context, id WorkflowID) (*Workflow, error)
}

type WorkflowID uint

type Workflow struct {
	ID       WorkflowID
	Progress uint
}

func (w *Workflow) Finished() bool {
	return w.Progress == 100
}
