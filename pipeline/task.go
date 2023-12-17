package pipeline

import "context"

type Task interface {
	Execute(ctx context.Context) (interface{}, error)
}
