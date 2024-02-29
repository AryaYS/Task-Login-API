package service

import (
	"apilogin/model/response"
	"context"
)

type ServiceInterface interface {
	LoginService(ctx context.Context, req response.User_req) (response.Response_user, string)
	RegisterService(ctx context.Context, req response.Create_req)
	AllWorkerByJobService(ctx context.Context, req int) response.Role
}
