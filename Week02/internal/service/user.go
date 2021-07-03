package service

import (
	"github.com/pkg/errors"
	"week02/internal/model"
)

type GetUserListRequest struct {
	Name string
}

func (srv *Service) GetUserList(name string, pageOffset, pageSize int) ([]*model.User, error) {
	users, err := srv.dao.GetUserList(name, pageOffset, pageSize)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("can not find users!")
	}
	return users, nil
}

func (srv *Service) GetUserDetail(name string) (*model.User, error) {
	return srv.dao.GetUserDetail(name)
}
