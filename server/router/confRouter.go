package router

import (
	"github.com/gin-gonic/gin"
	"sage/server/internal/app/conf"
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
//  conf router
// ------------------------------------------------------

func confRouter(v1 *gin.RouterGroup)  {
	confV1 := v1.Group("/conf")
	{
		confV1.GET("/list", conf.NewHandler().List)
		confV1.POST("", conf.NewHandler().Add)
		confV1.PUT("", conf.NewHandler().Update)
		confV1.DELETE("/:id", conf.NewHandler().Delete)
	}
}
