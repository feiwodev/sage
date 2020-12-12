package conf

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"sage/server/internal/models"
	"sage/server/utils/resp"
	"sage/server/utils/valid"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/11
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  配置api
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
	f := new(models.ConfForm)
	err := valid.Valid(cxt, f)
	if err != nil {
		return
	}

	conf, err := h.s.Add(f)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}

	resp.Build(cxt, nil).SetData(conf).OK()
	return
}

func (h *handler) Update(cxt *gin.Context) {
	f := new(models.ConfForm)
	err := valid.Valid(cxt, f)
	if err != nil {
		return
	}

	conf, err := h.s.Update(f)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}

	resp.Build(cxt, nil).SetData(conf).OK()
	return
}

func (h *handler) Delete(cxt *gin.Context) {
	id, err := strconv.Atoi(cxt.Param("id"))
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "ID参数错误")).OK()
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
	scenicId, err := strconv.Atoi(cxt.Query("scenicId"));
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "参数错误")).OK()
		return
	}
	serviceId, err := strconv.Atoi(cxt.Query("serviceId"))
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "参数错误")).OK()
		return
	}
	confList := h.s.List(uint(scenicId), uint(serviceId))
	resp.Build(cxt, nil).SetData(confList).OK()
	return
}



