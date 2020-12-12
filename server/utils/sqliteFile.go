package utils

import (
	"log"
	"os"
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
//  sqLite 操作
// ------------------------------------------------------

//const DbFileName =  "sage.db"
const DbFileName =  "/Users/feiwo/work/sage/sage.db"

type SqLiteFile interface {
	CreateDbFile()
}

var liteDbFile SqLiteFile
func CreateSqLiteFile() SqLiteFile {
	if liteDbFile == nil {
		liteDbFile = &sqLite{}
	}
	return liteDbFile
}

type sqLite struct {
}

func (s *sqLite) CreateDbFile() {
	_, err := os.Stat(DbFileName)
	if os.IsNotExist(err) {
		_, err := os.Create(DbFileName)
		if err != nil {
			log.Panicf("创建数据库文件失败 --> %v", err)
		}
		return
	}
}



