package router

import (
	"github.com/gin-gonic/gin"
	"sage/server/utils/midleware"
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
//  路由注册
// ------------------------------------------------------

func RegisterBackendRouterV1(engine *gin.Engine)  {
	v1 := engine.Group("/v1")
	v1.Use(midleware.AuthAndLogger())
	userHandler(v1)
	scenicHandler(v1)
	confRouter(v1)
	logRouter(v1)
	serveRouter(v1)
}

func RegisterApiV1(engine *gin.Engine)  {
	v1 := engine.Group("/api/v1")
	apiRouters(v1)
}

