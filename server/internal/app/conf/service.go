package conf

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
// Created by fei wo at 2020/11/11
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  配置服务
// ------------------------------------------------------

type IService interface {
	Add(form *models.ConfForm) (*models.Conf, error)
	Update(form *models.ConfForm) (*models.Conf, error)
	Delete(id uint) error
	List(scenicId, serviceId uint) []*models.Conf
}

var (
	s IService
	confLock sync.Mutex
)

func NewService() IService {
	if s == nil {
		confLock.Lock()
		defer confLock.Unlock()
		if s == nil {
			s = &service{database: db.Client().GetDB()}
		}
	}
	return s
}

type service struct {
	database *gorm.DB
}

func (s *service) Add(form *models.ConfForm) (*models.Conf, error) {
	conf := &models.Conf{
		ScenicId:  form.ScenicId,
		ServiceId: form.ServiceId,
		FileName:  form.FileName,
		FilePath:  form.FilePath,
		ConfInfo:  form.ConfInfo,
	}

	err := s.database.Create(conf).Error
	if err != nil {
		return nil, resp.NewError(resp.AddErrorCode, "添加配置信息失败")
	}

	return conf, nil
}

func (s *service) Update(form *models.ConfForm) (*models.Conf, error) {
	c := new(models.Conf)
	err := s.database.Model(models.Conf{}).Where("id = ?", form.Id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, resp.NewError(resp.ConfNotFoundCode,"配置信息不存在")
		}
	}

	c.ID = form.Id
	c.ServiceId = form.ServiceId
	c.ScenicId = form.ScenicId
	c.IsUse = form.IsUse
	c.FileName = form.FileName
	c.FilePath = form.FilePath
	c.ConfInfo = form.ConfInfo

	err = s.database.Save(c).Error
	if err != nil {
		log.Printf("更新配置信息失败 --> %s", err.Error())
		return nil, resp.NewError(resp.UpdateErrorCode, "更新配置信息失败")
	}
	return c, nil
}

func (s *service) Delete(id uint) error {
	var count int64 = 0
	s.database.Model(models.Conf{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return resp.NewError(resp.ConfNotFoundCode, "配置信息不存在")
	}
	conf := new(models.Conf)
	err := s.database.Where("id = ?", id).Delete(conf).Error
	if err != nil {
		return resp.NewError(resp.DeleteErrorCode, "删除配置信息失败")
	}
	return nil
}

func (s *service) List(scenicId, serviceId uint) []*models.Conf {
	confList := make([]*models.Conf, 15)
	where := s.database.Model(models.Conf{}).Where("scenic_id = ?", scenicId)
	if serviceId != 0 {
		where.Where("service_id = ?", serviceId)
	}
	where.Find(&confList)

	// set scenic name and service name
	if len(confList) > 0 {
		var scenicIds []uint
		for _, val := range confList {
			scenicIds = append(scenicIds, val.ScenicId)
		}
		scenicList := make([]*models.Scenic,len(confList))
		err := s.database.Where("id in ?", scenicIds).Find(&scenicList).Error
		if err != nil {
			log.Printf("查询景区失败 --> %s", err.Error())
		}
		for _, conf := range confList {
			for _, sc := range scenicList {
				if conf.ScenicId == sc.ID {
					conf.ScenicName = sc.Name
				}
			}
		}

		var serviceIds []uint
		for _, val := range confList {
			serviceIds = append(serviceIds, val.ServiceId)
		}
		serveList := make([]*models.Serve, len(confList))
		err = s.database.Where("id in ?", serviceIds).Find(&serveList).Error
		if err != nil {
			log.Printf("查询服务失败 --> %s", err.Error())
		}

		for _, conf := range confList {
			for _, s := range serveList {
				if conf.ServiceId == s.ID {
					conf.ServiceName = s.Name
					conf.ServiceType = s.Type
				}
			}
		}
	}
	return confList
}


