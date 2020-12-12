package resp

// ------------------------------------------------------
// Created by fei wo at 2020/11/8
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  异常code
// ------------------------------------------------------

// sys error
const (
	SysError = 500
	Success = 200
)

// db error
const (
	ParamsErrorCode = iota + 1000
	DeleteErrorCode
	AddErrorCode
	UpdateErrorCode
	QueryErrorCode
	FindRecordExists
)

// business error
const (
	UserNotFoundCode = iota + 2000
	LoginErrorCode
	ScenicNotFoundCode
	ServeNotFoundCode
	ConfNotFoundCode
)