package models

import (
	"sage/server/internal/db"
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
//  操作日志记录
// ------------------------------------------------------


type OptLog struct {
	db.Model
	// 操作用户ID
	UId uint `json:"userId"`
	// 用户名称
	UserName string `json:"userName"`
	// 动作
	Action string `json:"action"`
	// URL
	Url string `json:"url"`
}

type OptLogForm struct {
	// 操作用户ID
	UId uint
	// 用户名称
	UserName string
	// 动作
	Action string
	// URL
	Url string
}