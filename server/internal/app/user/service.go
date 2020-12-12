package user

import (
	"errors"
	"log"
	"strconv"
	"sync"

	"gorm.io/gorm"
	"sage/server/internal/db"
	"sage/server/internal/models"
	"sage/server/utils/encrypt"
	"sage/server/utils/jwt"
	"sage/server/utils/resp"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/8
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  user服务处理
// ------------------------------------------------------

type IUserService interface {
	Create(form *models.UserForm) (*models.User, error)
	Update(form *models.UserForm) (*models.User, error)
	Login(login *models.UserLogin) (*models.User, error)
	Delete(userId uint) error
	List() []*models.User
}

var (
	iService        IUserService
    userServiceLock sync.Mutex
)
func userService() IUserService {
	if iService == nil{
		userServiceLock.Lock()
		defer userServiceLock.Unlock()
		if iService == nil {
			iService = &service{database: db.Client().GetDB()}
		}
	}
	return iService
}

type service struct {
	database *gorm.DB
}

func (s *service) Login(login *models.UserLogin) (*models.User, error) {
	user := new(models.User)
	err := s.database.Where("name = ? and password = ?", login.Name, encrypt.Crypt(login.Password)).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.NewError(resp.UserNotFoundCode, "用户名或密码错误")
		}
		return nil, err
	}

	token, err := jwt.New().NewToken(user.ID, user.Name)
	if err != nil {
		return nil, resp.NewError(resp.LoginErrorCode, err.Error())
	}
	user.Token = token
	AddSession(strconv.Itoa(int(user.ID)), user)

	return user, nil
}

func (s *service) Delete(userId uint) error {
	return s.database.Where("id = ?", userId).Delete(&models.User{}).Error
}

func (s *service) Create(form *models.UserForm) (*models.User, error) {
	u := new(models.User)
	u.Name = form.Name
	u.Password = encrypt.Crypt(form.Password)
	u.IsUse = true
	u.Rule = 2

	err := s.database.Create(u).Error
	if err != nil {
		return nil, resp.NewError(resp.AddErrorCode, "新建用户失败")
	}
	return u, nil
}

func (s *service) Update(form *models.UserForm) (*models.User, error) {
	if form.ID == 0 {
		return nil, resp.NewError(resp.ParamsErrorCode, "用户ID异常")
	}
	user := new(models.User)
	err := s.database.Model(models.User{}).Where("id = ?", form.ID).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.NewError(resp.UserNotFoundCode, "当前用户不存在")
		}
		log.Printf("查询数据库异常 --> %s", err.Error())
		return nil, resp.NewError(resp.QueryErrorCode, "查询数据库失败")
	}
	user.Name = form.Name
	user.Password = encrypt.Crypt(form.Password)
	user.IsUse = form.IsUse

	err = s.database.Save(user).Error
	if err != nil {
		log.Printf("更新用户失败 --> %s", err.Error())
		return nil, resp.NewError(resp.UpdateErrorCode, "更新用户失败")
	}
	return user, nil
}

func (s *service) List() []*models.User {
	users := make([]*models.User, 15)
	err := s.database.Find(&users).Error
	if err != nil {
		log.Printf("查询用户失败 --> %s", err.Error())
	}
	return users
}
