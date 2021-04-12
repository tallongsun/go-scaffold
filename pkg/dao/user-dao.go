package dao

import (
	"github.com/tallongsun/go-scaffold/pkg/lib/db"
	"github.com/tallongsun/go-scaffold/pkg/model"
)

func GetUserById(id int64) (*model.User, error) {
	user := model.User{}
	res := db.Engine.First(&user, id)
	return &user, res.Error
}

func CreateUser(u *model.User) (int64, error) {
	res := db.Engine.Create(u)
	return u.Id, res.Error
}

func UpdateUser(u *model.User) (int64, error) {
	res := db.Engine.Model(u).Update(u)
	return res.RowsAffected, res.Error
}

func DeleteUserById(id int64) (int64, error) {
	user := model.User{Id: id}
	res := db.Engine.Delete(&user)
	return res.RowsAffected, res.Error
}

func FindUsers() ([]*model.User, error) {
	var users []*model.User
	res := db.Engine.Find(&users)
	return users, res.Error
}
