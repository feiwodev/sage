package optLog

import (
	"log"
	"sync"

	"gorm.io/gorm"
	"sage/server/internal/db"
	"sage/server/internal/models"
	"sage/server/utils/resp"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/11
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  日志服务
// ------------------------------------------------------

type IService interface {
	Add(form *models.OptLogForm) (*models.OptLog, error)
	Delete(ids []uint) error
	List(index int, size int) ([]*models.OptLog, int64)
	Clear() error
}

var (
	s IService
	serviceLock sync.Mutex
)
func NewService() IService {
	if s == nil {
		serviceLock.Lock()
		defer serviceLock.Unlock()
		if s == nil {
			s = &service{database: db.Client().GetDB()}
		}
	}
	return s
}

type service struct {
	database *gorm.DB
}

func (s *service) Add(form *models.OptLogForm) (*models.OptLog, error) {
	o := &models.OptLog{
		UId:      form.UId,
		UserName: form.UserName,
		Action:   form.Action,
		Url:      form.Url,
	}

	err := s.database.Create(o).Error
	if err != nil {
		return nil, resp.NewError(resp.AddErrorCode, "添加日志失败")
	}

	return o, nil
}


func (s *service) Delete(ids []uint) error {
	err := s.database.Where("id in ?", ids).Delete(models.OptLog{}).Error
	if err != nil {
		log.Printf("批量删除日志失败 --> %s", err.Error())
		return resp.NewError(resp.DeleteErrorCode, "批量删除日志失败")
	}
	return nil
}

func (s *service) List(index int, size int) ([]*models.OptLog, int64) {
	logs := make([]*models.OptLog, 15)

	err := s.database.Offset((index - 1) * index).Limit(size).Order("created_at DESC").Find(&logs).Error
	if err != nil {
		log.Printf("查询日志失败 --> %s", err.Error())
	}

	var total int64
	s.database.Model(models.OptLog{}).Count(&total)

	return logs, total
}

func (s *service) Clear() error {
	err := s.database.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(models.OptLog{}).Error
	if err != nil{
		log.Printf("清空日志失败 --> %s", err.Error())
		return resp.NewError(resp.DeleteErrorCode, "清空日志失败")
	}
	return nil
}

