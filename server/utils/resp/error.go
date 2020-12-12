package resp

import (
	"encoding/json"
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
//  返回异常结构体
// ------------------------------------------------------

type Error struct {
	Code int
	Msg string
}

func (e Error) Error() string {
	bytes, _ := json.Marshal(e)
	return string(bytes)
}

func NewError(code int, msg string) error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func NewSysError() error {
	return &Error{
		Code: SysError,
		Msg:  "系统内部错误",
	}
}

func NewParamsError(msg string) error {
	return &Error{
		Code: ParamsErrorCode,
		Msg:  msg,
	}
}

