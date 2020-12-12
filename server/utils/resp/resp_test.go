package resp

import (
	"reflect"
	"testing"
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
//  test
// ------------------------------------------------------

func TestBuild(t *testing.T) {
	err := NewError(1, "error")
	t.Log(reflect.TypeOf(err))
	t.Logf("%T", err)
	t.Log(interface{}(err).(*Error))
}