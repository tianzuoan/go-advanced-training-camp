package dao

import (
	"github.com/pkg/errors"
	"week02/internal/model"
)

func (d *Dao) GetUserList(name string, pageOffset, pageSize int) ([]*model.User, error) {
	user := model.User{UserName: name}
	users, err := user.List(d.db, pageOffset, pageSize)
	if err != nil {
		return nil, errors.Wrap(err, "dao get users failed!")
	}
	return users, nil
}

func (d *Dao) GetUserDetail(name string) (*model.User, error) {
	user := model.User{UserName: name}
	u, err := user.Detail(d.db, name)
	if err != nil {
		//如果err为sql.ErrNoRows,包装后往上拋
		return nil, errors.Wrap(err, "dao get user detail failed!")
	}
	return u, nil
}
