package handler

import (
	"context"

	user "github.com/songxuexian/gogomicro/user-srv/proto/user"
)

type User struct{}

func (e *User) QueryUserByName(context.Context, *user.Request, *user.Response) error {
	panic("implement me")
}
