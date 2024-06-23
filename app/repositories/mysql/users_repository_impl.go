package mysql

import (
	"context"

	"github.com/anurag925/aero/app/models"
	"github.com/anurag925/aero/app/repositories"
	"github.com/anurag925/aero/core/logger"
	"github.com/uptrace/bun"
)

type usersRepositoryImpl struct {
	db *bun.DB
}

var _ repositories.UsersRepository = (*usersRepositoryImpl)(nil)

func NewUsersRepository(db *bun.DB) repositories.UsersRepository {
	return &usersRepositoryImpl{db}
}

func (u *usersRepositoryImpl) FindUserById(ctx context.Context, id int64) (models.User, error) {
	user := models.User{}
	if err := u.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		logger.Error(ctx, "error while fetching user by id", err)
		return models.User{}, err
	}
	return user, nil
}
