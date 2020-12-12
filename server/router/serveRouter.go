package router

import (
	"github.com/gin-gonic/gin"
	"sage/server/internal/app/serve"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/17
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  景区服务路由
// ------------------------------------------------------

func serveRouter(v1 *gin.RouterGroup)  {
	serveV1 := v1.Group("/serve")
	{
		serveV1.POST("", serve.NewHandler().Add)
		serveV1.GET("/list/:scenicId", serve.NewHandler().List)
		serveV1.PUT("", serve.NewHandler().Update)
		serveV1.DELETE("/:id", serve.NewHandler().Delete)
	}
}