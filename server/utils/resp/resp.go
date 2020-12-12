package resp

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
//  返回结果处理
// ------------------------------------------------------

type Response struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
	Cxt *gin.Context `json:"-"`
}

func Build(cxt *gin.Context, err error) *Response {
	if err != nil {
		switch err.(type) {
		case *Error:
			e := err.(*Error)
			return &Response{
				Code: e.Code,
				Msg:  e.Msg,
				Data: nil,
				Cxt:  cxt,
			}
		default:
			return &Response{
				Code: SysError,
				Msg:  err.Error(),
				Data: nil,
				Cxt:  cxt,
			}
		}

	}
	return &Response{
		Code: Success,
		Msg:  "success",
		Data: nil,
		Cxt: cxt,
	}
}

func (r *Response) SetData(data interface{}) *Response{
	r.Data = data
	return r
}

func (r *Response) SetPageData(total int64, data interface{}) *Response{
	index, _ := strconv.Atoi(r.Cxt.DefaultQuery("index", "1"))
	size, _ := strconv.Atoi(r.Cxt.DefaultQuery("size", "15"))
	r.Data = buildPage(index, size, total, data)
	return r
}

func (r *Response) OK() {
	r.Cxt.JSON(http.StatusOK, r)
}