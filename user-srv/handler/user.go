package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"

	us "github.com/songxuexian/gogomicro/user-srv/models/user"
	user "github.com/songxuexian/gogomicro/user-srv/proto/user"
)

type Service struct{}

var (
	userService us.Service
)

func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] init handler error")
		return
	}

}
func (e *Service) QueryUserByName(ctx context.Context, req *user.Request, rsp *user.Response) error {
	qUser, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &user.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return err
	}

	rsp.User = qUser
	rsp.Success = true

	return nil
}
