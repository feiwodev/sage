package migrate

import (
	"log"

	"gorm.io/gorm"
	"sage/server/internal/db"
	"sage/server/internal/models"
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
//  创建数据库表
// ------------------------------------------------------

type Migrate interface {
	Migrate()
}

var m Migrate
func NewMigrate() Migrate {
	if m == nil {
		m =&migrate{}
	}
	return m
}

type migrate struct {
	db *gorm.DB
}

func (m *migrate) Migrate() {
	database := db.Client().GetDB()
	database.Set("gorm:table_options", "ENGINE=InnoDB")
	m.db = database

	m.autoCreateTable()
}

func (m *migrate) autoCreateTable() {
	var tables []interface{}
	tables = append(tables, &models.Scenic{})
	tables = append(tables, &models.Serve{})
	tables = append(tables, &models.Conf{})
	tables = append(tables, &models.User{})
	tables = append(tables, &models.OptLog{})
	err := m.db.AutoMigrate(tables...)
	if err != nil {
		log.Panicf("Migrate Table error --> %v", err)
	}
	log.Println("Migrate tables success")
}


