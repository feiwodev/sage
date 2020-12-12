package user

import (
	"log"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"sage/server/internal/models"
	"sage/server/utils/resp"
	"sage/server/utils/valid"
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
//  User请求处理
// ------------------------------------------------------

type IHandler interface {
	Login(cxt *gin.Context)
	Add(cxt *gin.Context)
	Update(cxt *gin.Context)
	Delete(cxt *gin.Context)
	List(cxt *gin.Context)
}

var (
	h IHandler
    userHandlerLock sync.Mutex
)
func NewUserHandler() IHandler {
	if h == nil {
		userHandlerLock.Lock()
		defer userHandlerLock.Unlock()
		if h == nil {
			h = &handler{s: userService()}
		}
	}
	return h
}

type handler struct {
	s IUserService
}

func (h *handler) Login(cxt *gin.Context) {
	loginForm := new(models.UserLogin)
	err := valid.Valid(cxt, loginForm)
	if err != nil {
		return
	}

	user, err := h.s.Login(loginForm)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).SetData(user).OK()
	return
}

func (h *handler) Delete(cxt *gin.Context) {
	param := cxt.Param("userId")
	userId, paramsError := strconv.Atoi(param)
	if paramsError != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, paramsError.Error())).OK()
		return
	}
	err := h.s.Delete(uint(userId))
	if err != nil {
		log.Printf("删除用户失败 --> %s", err.Error())
		resp.Build(cxt, resp.NewError(resp.DeleteErrorCode, "删除用户失败")).OK()
		return
	}
	resp.Build(cxt, nil).OK()
	return
}

func (h *handler) List(cxt *gin.Context) {
	list := h.s.List()
	resp.Build(cxt, nil).SetData(list).OK()
	return
}

func (h *handler) Add(cxt *gin.Context) {
	form := new(models.UserForm)
	validError := valid.Valid(cxt, form)
	if validError != nil{
		return
	}

	user, err := h.s.Create(form)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).SetData(user).OK()
	return
}

func (h *handler) Update(cxt *gin.Context) {
	f := new(models.UserForm)
	validError := valid.Valid(cxt, f)
	if validError != nil {
		return
	}
	user, err := h.s.Update(f)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}

	resp.Build(cxt, nil).SetData(user).OK()
	return
}
