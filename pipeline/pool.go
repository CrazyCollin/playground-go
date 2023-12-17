package pipeline

import "context"

type Pool interface {
	Submit(ctx context.Context, f func())
}
