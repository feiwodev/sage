package serve

import (
	"errors"
	"log"
	"sync"

	"gorm.io/gorm"
	"sage/server/internal/db"
	"sage/server/internal/models"
	"sage/server/utils/resp"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/10
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  运维服务
// ------------------------------------------------------

type IService interface {
	Add(form *models.ServeForm) (*models.Serve, error)
	Update(form *models.ServeForm) (*models.Serve, error)
	Delete(id uint) error
	List(scenicId uint) []*models.Serve
}

var (
	s           IService
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

func (s *service) Add(form *models.ServeForm) (*models.Serve, error) {
	serve := &models.Serve{
		ScenicId:     form.ScenicId,
		Name:         form.Name,
		Code:         form.Code,
		FileUrl:      form.FileUrl,
		VersionCode:  form.VersionCode,
		Type:         form.Type,
		SoftWareType: form.SoftWareType,
	}

	err := s.database.Create(serve).Error
	if err != nil {
		log.Printf("添加服务失败 --> %s", err.Error())
		return nil, resp.NewError(resp.AddErrorCode, "添加服务失败")
	}
	return serve, nil
}

func (s *service) Update(form *models.ServeForm) (*models.Serve, error) {
	serve := new(models.Serve)
	err := s.database.Model(models.Serve{}).Where("Id = ?", form.Id).First(serve).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.NewError(resp.ServeNotFoundCode, "服务不存在")
		}
		return nil, err
	}

	serve.Name = form.Name
	serve.Code = form.Code
	serve.ScenicId = form.ScenicId
	serve.FileUrl = form.FileUrl
	serve.VersionCode = form.VersionCode
	serve.Type = form.Type
	serve.SoftWareType = form.SoftWareType

	err = s.database.Save(serve).Error
	if err != nil {
		return nil, resp.NewError(resp.UpdateErrorCode, "服务更新失败")
	}
	return serve, nil
}

func (s *service) Delete(id uint) error {
	serve := new(models.Serve)
	err := s.database.Model(models.Serve{}).Where("Id = ?", id).First(serve).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.NewError(resp.ServeNotFoundCode, "服务不存在")
		}
		return err
	}

	err = s.database.Where("Id = ?", id).Delete(serve).Error
	if err != nil {
		return resp.NewError(resp.DeleteErrorCode, "删除服务失败")
	}
	return nil
}

func (s *service) List(scenicId uint) []*models.Serve {
	serves := make([]*models.Serve, 15)

	err := s.database.Where("scenic_id = ?", scenicId).Find(&serves).Error
	if err != nil {
		log.Printf("查询服务列表失败 --> %s", err.Error())
	}
	var scenicIds []uint
	for _, s := range serves {
		scenicIds = append(scenicIds, s.ScenicId)
	}

	scenicList := make([]*models.Scenic, 15)
	err = s.database.Where("Id in ?", scenicIds).Find(&scenicList).Error
	if err != nil {
		log.Printf("查询景点服务列表失败 --> %s", err.Error())
	}

	for _, s := range serves {
		for _, sce := range scenicList {
			if s.ScenicId == sce.ID {
				s.ScenicName = sce.Name
			}
		}
	}

	return serves
}
