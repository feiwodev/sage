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
//  配置文件信息
// ------------------------------------------------------

type Conf struct {
	db.Model
	// 景区ID
	ScenicId uint `json:"scenicId"`
	// 景区名称
	ScenicName string `gorm:"-" json:"scenicName"`
	// 服务ID
	ServiceId uint `json:"serviceId"`
	// 服务名称
	ServiceName string `gorm:"-" json:"serviceName"`
	// 服务类型
	ServiceType Type `gorm:"-" json:"serviceType"`
	// 配置文件名称
	FileName string `json:"fileName"`
	// 配置文件路径
	FilePath string `json:"filePath"`
	// 是否启用
	IsUse bool `json:"isUse"`
	// 配置信息
	ConfInfo string `json:"confInfo"`
}

type ConfForm struct {
	Id uint
	// 景区ID
	ScenicId uint `json:"scenicId"`
	// 服务ID
	ServiceId uint `json:"serviceId"`
	// 配置文件名称
	FileName string `json:"fileName" valid:"required~配置文件名称不能为空"`
	// 配置文件路径
	FilePath string `json:"filePath" valid:"required~配置文件路径不能为空"`
	// 是否启用
	IsUse bool `json:"isUse"`
	// 配置信息
	ConfInfo string `json:"confInfo" valid:"required~配置内容不能为空"`
}