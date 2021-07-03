package model

import (
	"gorm.io/gorm"
)

type User struct {
	//Model
	UserId   int    `json:"user_id" gorm:"primaryKey"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) (users []*User, err error) {
	db = db.Offset(pageOffset).Limit(pageSize)
	if err = db.Where("user_name = ?", u.UserName).Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

func (u User) Detail(db *gorm.DB, name string) (user *User, err error) {
	if err = db.Where("user_name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return
}
