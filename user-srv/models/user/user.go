package user

import (
	"fmt"
	proto "github.com/songxuexian/gogomicro/user-srv/proto/user"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

type Service interface {
	QueryUserByName(userName string) (ret *proto.User, err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetServie] GetSerivice not inited")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}
	s = &service{}
}
