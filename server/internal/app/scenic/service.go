package scenic

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
//  景区服务
// ------------------------------------------------------

type IService interface {
	Add(form *models.ScenicForm) (*models.Scenic, error)
	Update(form *models.ScenicForm) (*models.Scenic, error)
	Delete(id uint) error
	List(index int, size int) ([]*models.Scenic, int64)
	FindScenicByCode(code string) (*models.ScenicInfo, error)
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

func (s *service) Add(form *models.ScenicForm) (*models.Scenic, error) {
	scenic := new(models.Scenic)
	scenic.Name = form.Name
	scenic.Code = form.Code
	scenic.AutoUpdate = form.AutoUpdate

	var codeCount int64 = 0
	err := s.database.Model(models.Scenic{}).Where("code = ?", form.Code).Count(&codeCount).Error
	if err != nil {
		log.Printf("查询景区code失败")
		return nil, resp.NewError(resp.QueryErrorCode, "添加景区失败")
	}
	if codeCount > 0{
		return nil, resp.NewError(resp.FindRecordExists, "景区code已存在")
	}

	err = s.database.Create(scenic).Error
	if err != nil {
		log.Printf("添加景区失败 --> %s", err.Error())
		return nil, resp.NewError(resp.AddErrorCode, "添加景区失败")
	}
	return scenic, nil
}

func (s *service) Update(form *models.ScenicForm) (*models.Scenic, error) {
	scenic := new(models.Scenic)
	err := s.database.Model(models.Scenic{}).Where("id = ?", form.Id).First(scenic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.NewError(resp.ScenicNotFoundCode, "景区不存在")
		}
		return nil, err
	}
	scenic.Name = form.Name
	scenic.Code = form.Code
	scenic.AutoUpdate = form.AutoUpdate

	err = s.database.Save(scenic).Error
	if err != nil {
		log.Printf("更新景区失败 --> %s", err.Error())
		return nil, resp.NewError(resp.UpdateErrorCode, "更新景区失败")
	}
	return scenic, nil
}

func (s *service) Delete(id uint) error {
	scenic := new(models.Scenic)
	err := s.database.Model(models.Scenic{}).Where("id = ?", id).First(scenic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.NewError(resp.ScenicNotFoundCode, "景区不存在")
		}
		return err
	}
	err = s.database.Where("id = ?", id).Delete(scenic).Error
	if err != nil {
		return resp.NewError(resp.DeleteErrorCode, "删除景区失败")
	}
	return nil
}

func (s *service) List(index int, size int) ([]*models.Scenic, int64) {
	scenicList := make([]*models.Scenic,15)

	var count int64 = 0
	err := s.database.Model(models.Scenic{}).Count(&count).Error
	if err != nil {
		log.Printf("查询景区列表失败 --> %s", err.Error())
	}
	err = s.database.Offset((index - 1) * size).Limit(size).Find(&scenicList).Error
	if err != nil {
		log.Printf("查询景区列表失败 --> %s", err.Error())
	}
	return scenicList, count
}

func (s *service) FindScenicByCode(code string) (*models.ScenicInfo, error)  {
	info := new(models.ScenicInfo)

	scenic := new(models.Scenic)
	err := s.database.Model(models.Scenic{}).Where("code = ?", code).First(scenic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.NewError(resp.QueryErrorCode, "景区不存在")
		}
		return nil, err
	}

	serves := make([]*models.Serve, 0, 15)
	err = s.database.Model(models.Serve{}).Where("scenic_id = ?", scenic.ID).Find(&serves).Error
	if err != nil {
		return nil, resp.NewError(resp.QueryErrorCode, "景区服务不存在")
	}

	info.Serves = serves
	info.ID = scenic.ID
	info.Name = scenic.Name
	info.Code = scenic.Code
	info.AutoUpdate = scenic.AutoUpdate
	info.CreatedAt = scenic.CreatedAt

	return info, nil
}

