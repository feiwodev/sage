package midleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sage/server/internal/app/optLog"
	"sage/server/internal/app/user"
	"sage/server/internal/models"
	"sage/server/utils/jwt"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/22
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  日志中间件
// ------------------------------------------------------

func AuthAndLogger() gin.HandlerFunc {
	iJwt := jwt.New()
	logService := optLog.NewService()
	return func(cxt *gin.Context) {
		token := cxt.GetHeader("token")

		if len(token) > 0 {
			id, name := iJwt.Parser(token)
			session := user.GetSession(strconv.Itoa(int(id)))
			if session == nil{
				cxt.AbortWithStatusJSON(http.StatusOK,gin.H{
					"code": 401,
					"msg": "Token已过期",
				})
			}
			form := &models.OptLogForm{
				UId:      id,
				UserName: name,
				Action:   cxt.Request.Method,
				Url:      cxt.Request.URL.String(),
			}
			_, err := logService.Add(form)
			if err != nil {
				log.Printf("写入日子失败 --> %s", err.Error())
			}
		}
	}
}