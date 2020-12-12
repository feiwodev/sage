package main

import (
	"sage/server/router"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/7
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  后台服务入口
// ------------------------------------------------------

func main()  {
	//utils.CreateSqLiteFile().CreateDbFile()
	//migrate.NewMigrate().Migrate()
	router.NewServer().Start()
}