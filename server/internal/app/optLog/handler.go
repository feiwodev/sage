package optLog

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"sage/server/utils/resp"
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
//  日志api
// ------------------------------------------------------

type IHandler interface {
	List(cxt *gin.Context)
	Delete(cxt *gin.Context)
	Clear(cxt *gin.Context)
}

var  (
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

func (h *handler) List(cxt *gin.Context) {
	index, err := strconv.Atoi(cxt.DefaultQuery("index", "1"))
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode,"index 参数错误")).OK()
		return
	}
	size, err := strconv.Atoi(cxt.DefaultQuery("size", "15"))
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "size 参数错误")).OK()
		return
	}

	list, total:= h.s.List(index, size)
	resp.Build(cxt, nil).SetPageData(total, list).OK()
	return
}

func (h *handler) Delete(cxt *gin.Context) {
	var ids []uint
	err := cxt.ShouldBindJSON(&ids)
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "删除参数错误")).OK()
		return
	}

	err = h.s.Delete(ids)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}

	resp.Build(cxt, nil).OK()
	return
}

func (h *handler) Clear(cxt *gin.Context) {
	err := h.s.Clear()
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).OK()
	return
}
