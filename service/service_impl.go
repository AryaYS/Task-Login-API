package service

import (
	"apilogin/exception"
	"apilogin/helper"
	"apilogin/model"
	"apilogin/model/response"
	"apilogin/repository"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	repo     repository.RepoApp
	tx       *gorm.DB
	validate *validator.Validate
}

// AllWorkerByJobService implements ServiceInterface.
func (s *ServiceImpl) AllWorkerByJobService(ctx context.Context, req int) response.Role {
	res := s.repo.GetListWorkerByJob(ctx, s.tx, req)
	resp := response.Role{
		Role_id:   res.Role_id,
		Role_name: res.Role_name,
	}
	for i := 0; i < len(res.User); i++ {
		ad := response.Response_user{
			User_name: res.User[i].User_name,
			Role_id:   res.User[i].Role_id,
		}
		resp.User = append(resp.User, ad)
	}
	return resp
}

// LoginService implements ServiceInterface.
func (s *ServiceImpl) LoginService(ctx context.Context, req response.User_req) (response.Response_user, string) {
	user := s.repo.SelectByName(ctx, s.tx, req.User_name)
	hel := helper.CheckPass(req.User_pass, user.User_pass)
	if !hel {
		panic(exception.BadRequestF("Wrong Password"))
	}
	res := response.Response_user{
		User_name: user.User_name,
		Role_id:   user.Role_id,
	}
	tokenString := helper.GenerateJWT(res)
	return res, tokenString
}

// RegisterService implements ServiceInterface.
func (s *ServiceImpl) RegisterService(ctx context.Context, req response.Create_req) {
	err := s.validate.Struct(req)
	if err != nil {
		panic(exception.BadRequestF("Username Or Password not good enough"))
	}

	tx := s.tx.Begin()
	hashed, err := helper.HashPassword(req.User_pass)
	if err != nil {
		panic(err)
	}
	inpt := model.User{
		User_name: req.User_name,
		User_pass: hashed,
		Role_id:   req.Role_id,
	}
	err = s.repo.CreateUser(ctx, tx, inpt)
	if err != nil {
		panic(exception.BadRequestF("Username already Exists!"))
	}
	defer helper.CommitOrRollback(tx)
}

func NewService(repos repository.RepoApp, txs *gorm.DB, val *validator.Validate) ServiceInterface {
	return &ServiceImpl{
		repo:     repos,
		tx:       txs,
		validate: val,
	}
}
