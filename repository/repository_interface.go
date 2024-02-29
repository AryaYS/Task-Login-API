package repository

import (
	"apilogin/model"
	"context"

	"gorm.io/gorm"
)

type RepoApp interface {
	SelectByName(ctx context.Context, db *gorm.DB, username string) model.User
	CreateUser(ctx context.Context, db *gorm.DB, req model.User) error
	GetListWorkerByJob(ctx context.Context, db *gorm.DB, id_role int) model.Role
}
