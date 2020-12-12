package router

import (
	"github.com/gin-gonic/gin"
	"sage/server/internal/app/scenic"
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
//  景区router
// ------------------------------------------------------

func scenicHandler(scenicV1 *gin.RouterGroup)  {
	scenicV1 = scenicV1.Group("/scenic")
	{
		scenicV1.GET("/list", scenic.NewHandler().List)
		scenicV1.POST("", scenic.NewHandler().Create)
		scenicV1.PUT("", scenic.NewHandler().Update)
		scenicV1.DELETE("/:id", scenic.NewHandler().Delete)
	}
}