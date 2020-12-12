package db

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sage/server/utils"
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
//  数据库连接相关
// ------------------------------------------------------

type Conn interface {
	connDb()
	GetDB() *gorm.DB
}

func config() *gorm.Config {
	newLogger := logger.New(log.New(os.Stdout,"\r\n",log.LstdFlags),logger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      logger.Silent,
	})
	return &gorm.Config{
		Logger: newLogger,
	}
}

var c Conn
var once sync.Once
func Client() Conn {
	once.Do(func() {
		c = &conn{}
	})
	return c
}

type conn struct {
	db *gorm.DB
}


func (c *conn) GetDB() *gorm.DB {
	c.connDb()
	return c.db
}

var dbLock sync.Mutex
func (c *conn) connDb() {
	var  err error
	if c.db == nil {
		dbLock.Lock()
		defer dbLock.Unlock()
		if c.db == nil {
			c.db, err = gorm.Open(sqlite.Open(utils.DbFileName), config())
			if err != nil {
				log.Panicf("开发数据库失败 --> %v", err)
			}
			log.Println("连接数据库成功...")
		}
	}
}
