package inventory

import (
	"fmt"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

type Service interface {
	Sell(bookId, userId int64) (id int64, err error)
	Confirm(id int64, state int) (err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService not init")
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
