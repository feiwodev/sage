package user

import (
	"encoding/json"
	"fmt"
	"log"

	"sage/server/internal/models"
	"sage/server/pkg/cache"
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
//  用户登录缓存，保持session
// ------------------------------------------------------

const expire = 604800
const Key = "sage.user.%s"

func AddSession(key string, user *models.User)  {
	cache.Set(fmt.Sprintf(Key, key), user, int(expire))
}

func GetSession(key string) *models.User {
	user := new(models.User)
	userCache, err := cache.Get(fmt.Sprintf(Key, key))
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(userCache), user)
	if err != nil {
		log.Printf("缓存序列化失败 --> %s", err.Error())
	}
	return user
}

func DelSession(key string)  {
	cache.Del(key)
}