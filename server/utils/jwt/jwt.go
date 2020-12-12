package jwt

import (
	"log"
	"sync"

	"github.com/dgrijalva/jwt-go"
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
//  jwtToken token util
// ------------------------------------------------------

const SigningKey = "sage"

type IJwt interface {
	NewToken(id uint, name string) (string, error)
	Parser(token string) (uint, string)
}

var (
	j IJwt
	jwtLock sync.Mutex
)

func New() IJwt {
	if j == nil {
		jwtLock.Lock()
		defer jwtLock.Unlock()
		if j == nil {
			j = &jwtToken{}
		}
	}
	return j
}

type jwtToken struct {

}

func (j *jwtToken) NewToken(id uint, name string) (string, error) {
	mapClaims := jwt.MapClaims{
		"id":   id,
		"name": name,
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	token, err := claims.SignedString([]byte(SigningKey))
	if err != nil {
		log.Printf("生成Token失败 --> %s", err.Error())
		return "", err
	}
	return token, nil
}

func (j *jwtToken) Parser(token string) (uint, string) {
	mapClaims := make(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(token, mapClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		log.Printf("解析Token失败 --> %s", err.Error())
	}
	id := mapClaims["id"].(float64)
	name := mapClaims["name"].(string)
	return uint(id), name
}
