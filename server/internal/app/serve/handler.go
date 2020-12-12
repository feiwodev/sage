package serve

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"sage/server/internal/models"
	"sage/server/utils/resp"
	"sage/server/utils/valid"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/10
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  api
// ------------------------------------------------------

type IHandler interface {
	Add(cxt *gin.Context)
	Update(cxt *gin.Context)
	Delete(cxt *gin.Context)
	List(cxt *gin.Context)
}

var (
	h IHandler
	handlerLock sync.Mutex
)
func NewHandler() IHandler {
	if h == nil {
		handlerLock.Lock()
		defer handlerLock.Unlock()
		if h == nil {
			h = &handler{s: NewService()}
		}
	}
	return h
}

type handler struct {
	s IService
}

func (h *handler) Add(cxt *gin.Context) {
	f := new(models.ServeForm)
	err := valid.Valid(cxt, f)
	if err != nil {
		return
	}

	serve, err := h.s.Add(f)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).SetData(serve).OK()
	return
}

func (h *handler) Update(cxt *gin.Context) {
	f := new(models.ServeForm)
	err := valid.Valid(cxt, f)
	if err != nil {
		return
	}

	serve, err := h.s.Update(f)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).SetData(serve).OK()
	return
}

func (h *handler) Delete(cxt *gin.Context) {
	param := cxt.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "删除ID错误")).OK()
		return
	}
	err = h.s.Delete(uint(id))
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).OK()
	return
}

func (h *handler) List(cxt *gin.Context) {
	scenicId, err := strconv.Atoi(cxt.Param("scenicId"))
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode,"景区ID错误")).OK()
		return
	}
	serves := h.s.List(uint(scenicId))
	resp.Build(cxt, nil).SetData(serves).OK()
	return
}
