package scenic

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
//  景区api
// ------------------------------------------------------

type IHandler interface {
	Create(cxt *gin.Context)
	Update(cxt *gin.Context)
	Delete(cxt *gin.Context)
	List(cxt *gin.Context)
	FindScenicByCode(cxt *gin.Context)
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

func (h *handler) Create(cxt *gin.Context) {
	f := new(models.ScenicForm)
	err := valid.Valid(cxt, f)
	if err != nil {
		return
	}

	scenic, err := h.s.Add(f)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}

	resp.Build(cxt, nil).SetData(scenic).OK()
	return
}

func (h *handler) Update(cxt *gin.Context) {
	f := new(models.ScenicForm)
	err := valid.Valid(cxt, f)
	if err != nil {
		return
	}

	if f.Id == 0 {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "景区ID不能为空")).OK()
		return
	}
	scenic, err := h.s.Update(f)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}

	resp.Build(cxt, nil).SetData(scenic).OK()
	return
}

func (h *handler) Delete(cxt *gin.Context) {
	param := cxt.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "删除ID不能为空")).OK()
		return
	}
	err = h.s.Delete(uint(id))
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).OK()
}

func (h *handler) List(cxt *gin.Context) {
	index, _ := strconv.Atoi(cxt.DefaultQuery("index", "1"))
	size, _ := strconv.Atoi(cxt.DefaultQuery("size", "15"))
	scenicList, total := h.s.List(index, size)
	resp.Build(cxt, nil).SetPageData(total, scenicList).OK()
}

func (h *handler) FindScenicByCode(cxt *gin.Context)  {
	code := cxt.Param("code")
	if len(code) <= 0 {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, "参数code缺失"))
		return
	}

	info, err := h.s.FindScenicByCode(code)
	if err != nil {
		resp.Build(cxt, err).OK()
		return
	}
	resp.Build(cxt, nil).SetData(info).OK()
	return
}

