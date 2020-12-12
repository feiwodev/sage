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
//  登录用户信息
// ------------------------------------------------------

// entity
type User struct {
	db.Model
	// 用户名
	Name string `json:"name"`
	// 密码
	Password string `json:"password"`
	// 是否禁用
	IsUse bool `json:"isUse"`
	// 角色,1为管理员，后续添加皆为普通用户
	Rule uint `json:"rule"`
	// token
	Token string `json:"token" gorm:"-"`
}

// form
type UserForm struct {
	// 用户ID
	ID uint `json:"id"`
	// 用户名
	Name string `json:"name" valid:"required~用户名不能为空"`
	// 密码
	Password string `json:"password" valid:"required~用户密码不能为空"`
	// 是否禁用
	IsUse bool `json:"isUse"`
}

type UserLogin struct {
	// 登录名
	Name string `json:"name" valid:"required~登录名不能为空"`
	// 密码
	Password string `json:"password" valid:"required~登录密码不能为空"`
}