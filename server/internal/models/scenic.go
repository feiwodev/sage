package models

import (
	"sage/server/internal/db"
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
//  景区
// ------------------------------------------------------

type Scenic struct {
	db.Model
	// 景区名称
	Name string `json:"name"`
	// 景区编码
	Code string `json:"code"`
	// 是否开启自动更新
	AutoUpdate bool `json:"autoUpdate"`
}

type ScenicForm struct {
	Id uint `json:"id"`
	// 景区名称
	Name string `json:"name" valid:"required~景区名称不能为空"`
	// 景区编码
	Code string `json:"code" valid:"required~景区编码不能为空"`
	// 是否开启自动更新
	AutoUpdate bool `json:"autoUpdate"`
}

type ScenicInfo struct {
	db.Model
	// 景区名称
	Name string `json:"name"`
	// 景区编码
	Code string `json:"code"`
	// 是否开启自动更新
	AutoUpdate bool `json:"autoUpdate"`
	// 景区服务列表
	Serves []*Serve `json:"serves"`
}