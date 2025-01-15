package usecase

import "context"

type Transaction interface {
	RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
