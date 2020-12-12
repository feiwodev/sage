package valid

import (
	"github.com/asaskevich/govalidator"
	_ "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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
//  参数验证
// ------------------------------------------------------

func Valid(cxt *gin.Context, form interface{}) error {
	err := cxt.ShouldBindJSON(form)
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, err.Error())).OK()
	}
	_, err = govalidator.ValidateStruct(form)
	if err != nil {
		resp.Build(cxt, resp.NewError(resp.ParamsErrorCode, err.Error())).OK()
	}
	return err
}