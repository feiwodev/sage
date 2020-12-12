package router

import (
	"github.com/gin-gonic/gin"
	"sage/server/internal/app/optLog"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/12
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  日志路由
// ------------------------------------------------------

func logRouter(v1 *gin.RouterGroup)  {
	logV1 := v1.Group("/log")
	{
		logV1.GET("/list", optLog.NewHandler().List)
		logV1.DELETE("", optLog.NewHandler().Delete)
		logV1.DELETE("/clear", optLog.NewHandler().Clear)
	}
}