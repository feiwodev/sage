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
//  服务管理
// ------------------------------------------------------

type Type uint
const (
	Basic Type = iota + 1
	Ext
)

type SoftType uint

const (
	Java SoftType = iota + 1
	NodeJs
)

type Serve struct {
	db.Model
	// 景区ID
	ScenicId uint `json:"scenicId"`
	// 景区名称
	ScenicName string `gorm:"-" json:"scenicName"`
	// 服务名称
	Name string `json:"name"`
	// 服务编码
	Code string `json:"code"`
	// 服务文件地址
	FileUrl string `json:"fileUrl"`
	// 版本号
	VersionCode string `json:"versionCode"`
	// 服务类型
	Type Type `json:"type"`
	// 软件类型
	SoftWareType SoftType `json:"softwareType"`
}

type ServeForm struct {
	Id uint `json:"Id"`
	// 景区ID
	ScenicId uint `json:"scenicId" valid:"required~请选择服务所属景区"`
	// 服务名称
	Name string `json:"name" valid:"required~请输入服务名称"`
	// 服务编码
	Code string `json:"code" valid:"required~请输入服务编码"`
	// 服务文件地址
	FileUrl string `json:"fileUrl" valid:"required~请输入服务文件地址"`
	// 版本号
	VersionCode string `json:"versionCode" valid:"required~请输入版本号"`
	// 服务类型
	Type Type `json:"type" valid:"required~请选择服务类型"`
	// 软件类型
	SoftWareType SoftType `json:"softwareType" valid:"required~请选择软件类型"`
}