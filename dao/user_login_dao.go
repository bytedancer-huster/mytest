package dao

import (
	"github.com/jinzhu/gorm"
	"mytest/model"
)

type UserLoginDao struct {
	db *gorm.DB
}

func NewUserLoginTable() *UserLoginDao {
	dao := &UserLoginDao{db: userDB}
	dao.db = dao.db.Model(model.UserLogin{})
	return dao
}

func (u *UserLoginDao) WithUserName(userName string) *UserLoginDao {
	u.db = u.db.Where("user_name=?", userName)
	return u
}

func (u *UserLoginDao) WithPassword(password string) *UserLoginDao {
	u.db = u.db.Where("password=?", password)
	return u
}

func (u *UserLoginDao) All() ([]*model.UserLogin, error) {
	var res []*model.UserLogin
	err := u.db.Find(&res).Error
	return res, err
}

func (u *UserLoginDao) Create(user *model.UserLogin) error {
	return u.db.Create(user).Error
}

func (u *UserLoginDao) Update(update map[string]interface{}) error {
	return u.db.Update(update).Error
}
