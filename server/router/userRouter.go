package router

import (
	"github.com/gin-gonic/gin"
	"sage/server/internal/app/user"
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
//  用户模块路由
// ------------------------------------------------------

func userHandler(v1 *gin.RouterGroup)  {
	userV1 := v1.Group("/user")
	{
		userV1.POST("", user.NewUserHandler().Add)
		userV1.PUT("", user.NewUserHandler().Update)
		userV1.DELETE("/:userId", user.NewUserHandler().Delete)
		userV1.GET("/list", user.NewUserHandler().List)
		userV1.POST("/login", user.NewUserHandler().Login)
	}

}