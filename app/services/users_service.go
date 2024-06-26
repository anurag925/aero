package services

import (
	"context"

	"github.com/anurag925/aero/app/models"
)

type UsersService interface {
	FindUserById(ctx context.Context, id int64) (models.User, error)
}
