package cache

import (
	"encoding/json"
	"log"

	"github.com/coocood/freecache"
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
//  缓存工具
// ------------------------------------------------------

var mCache *freecache.Cache
func init()  {
	mCache = freecache.NewCache(10 * 1024 * 1024)
}

func Set(key string, data interface{}, expire int)  {
	switch data.(type) {
	case struct{}:
		bytes, _ := json.Marshal(data)
		err := mCache.Set([]byte(key), bytes, expire)
		if err != nil {
			log.Printf("缓存失败 --> %s", err.Error())
		}
		break
	case string:
		err := mCache.Set([]byte(key), []byte(data.(string)), 3000)
		if err != nil {
			log.Printf("缓存失败 --> %s", err.Error())
		}
	case interface{}:
		bytes, _ := json.Marshal(data)
		err := mCache.Set([]byte(key), bytes, expire)
		if err != nil {
			log.Printf("缓存失败 --> %s", err.Error())
		}

	}
}

func Get(key string) (string, error) {
	bytes, err := mCache.Get([]byte(key))
	if err != nil {
		log.Printf("获取缓存失败 --> %s", err.Error())
		return "", err
	}

	return string(bytes), nil
}

func Del(key string)  {
	del := mCache.Del([]byte(key))
	if !del {
		log.Printf("删除缓存失败")
	}
}
