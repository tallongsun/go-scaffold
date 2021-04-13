package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tallongsun/go-scaffold/pkg/dao"
	"github.com/tallongsun/go-scaffold/pkg/dto"
	"github.com/tallongsun/go-scaffold/pkg/lib/log"
	"github.com/tallongsun/go-scaffold/pkg/model"
	"net/http"
	"strconv"
	"time"
)

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Logger.Errorf("parse user id %s err %v ", id, err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user, err := dao.GetUserById(id)
	if err != nil {
		log.Logger.Errorf("get user id %s err %v ", id, err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	userRes := &dto.UserRes{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusOK, userRes)
}

func CreateUser(c *gin.Context) {
	userCreateReq := &dto.UserCreateReq{}
	err := c.ShouldBindJSON(userCreateReq)
	if err != nil {
		log.Logger.Errorf("parse user err %v ", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user := model.User{
		Name:      userCreateReq.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := dao.CreateUser(&user)
	if err != nil {
		log.Logger.Errorf("create user err %v ", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Logger.Errorf("parse user id %s err %v ", id, err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userUpdateReq := &dto.UserUpdateReq{}
	err = c.ShouldBindJSON(userUpdateReq)
	if err != nil {
		log.Logger.Errorf("parse user err %v ", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user := model.User{
		Id:        id,
		UpdatedAt: time.Now(),
	}
	if userUpdateReq.Name != "" {
		user.Name = userUpdateReq.Name
	}
	_, err = dao.UpdateUser(&user)
	if err != nil {
		log.Logger.Errorf("update user err %v ", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Logger.Errorf("parse user id %s err %v ", id, err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	_, err = dao.DeleteUserById(id)
	if err != nil {
		log.Logger.Errorf("delete user err %v ", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func ListUsers(c *gin.Context) {
	users, err := dao.FindUsers()
	if err != nil {
		log.Logger.Errorf("list user err %v ", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	userResList := make([]*dto.UserRes, len(users))
	for i, u := range users {
		userResList[i] = &dto.UserRes{
			Id:        u.Id,
			Name:      u.Name,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}
	}
	c.JSON(http.StatusOK, userResList)
}
