package interfaces

import (
	"context"

	"github.com/jlau-ice/collect/internal/types"
)

type UserService interface {
	Register(ctx context.Context, req types.User) (*types.User, error)
}
