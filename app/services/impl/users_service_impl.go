package impl

import (
	"context"

	"github.com/anurag925/aero/app/models"
	"github.com/anurag925/aero/app/repositories"
	"github.com/anurag925/aero/app/services"
)

type usersServiceImpl struct {
	usersRepo repositories.UsersRepository
}

var _ services.UsersService = (*usersServiceImpl)(nil)

func NewUsersService(usersRepo repositories.UsersRepository) *usersServiceImpl {
	return &usersServiceImpl{
		usersRepo: usersRepo,
	}
}

func (u *usersServiceImpl) FindUserById(ctx context.Context, id int64) (models.User, error) {
	return u.usersRepo.FindUserById(ctx, id)
}
