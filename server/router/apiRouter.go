package router

import (
	"github.com/gin-gonic/gin"
	"sage/server/internal/app/scenic"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/28
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  
// ------------------------------------------------------

func apiRouters(v1 *gin.RouterGroup)  {
	v1.GET("/scenic/:code",scenic.NewHandler().FindScenicByCode)
}