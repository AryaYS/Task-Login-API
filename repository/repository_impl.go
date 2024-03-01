package repository

import (
	"apilogin/exception"
	"apilogin/model"
	"context"

	"gorm.io/gorm"
)

type RepoImpl struct {
}

// DeleteUser implements RepoApp.
func (*RepoImpl) DeleteUser(ctx context.Context, db *gorm.DB, user model.User) error {
	obj := user
	err := db.WithContext(ctx).Delete(&obj)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdatePassword implements RepoApp.
func (*RepoImpl) UpdatePassword(ctx context.Context, db *gorm.DB, user model.User) error {
	md := user
	err := db.Model(&md).WithContext(ctx).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetListWorkerByJob implements RepoApp.
func (*RepoImpl) GetListWorkerByJob(ctx context.Context, db *gorm.DB, id_role int) model.Role {
	var res model.Role
	err := db.WithContext(ctx).Model(&model.Role{}).Preload("User").Where("role_id = ?", id_role).Find(&res).Error
	if err != nil {
		panic(err)
	}
	return res
}

// CreateUser implements RepoApp.
func (*RepoImpl) CreateUser(ctx context.Context, db *gorm.DB, req model.User) error {
	crt := req
	err := db.WithContext(ctx).Create(&crt).Error
	if err != nil {
		return err
	}
	return nil
}

// SelectByName implements RepoApp.
func (*RepoImpl) SelectByName(ctx context.Context, db *gorm.DB, username string) model.User {
	var res model.User
	err := db.WithContext(ctx).Where("user_name = ?", username).Take(&res).Error
	if err != nil {
		panic(exception.NotFoundErrorF(err.Error()))
	}
	return res
}

func NewRepo() RepoApp {
	return &RepoImpl{}
}
