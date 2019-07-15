package access

import (
	"fmt"
	"github.com/go-redis/redis"
	rd "github.com/songxuexian/gogomicro/basic/redis"
	"sync"
)

var (
	s *service
	r *redis.Client
	m sync.RWMutex
)

type service struct {
}

type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)
}



// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	r = rd.GetRedis()

	s = &service{}
}