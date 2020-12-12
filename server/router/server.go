package router

import (
	"log"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ------------------------------------------------------
// Created by fei wo at 2020/11/8
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  gin server
// ------------------------------------------------------

type IHttpServer interface {
	Start()
}

var (
	httpServer IHttpServer
	serverLock sync.Mutex
)
func NewServer() IHttpServer {
	if httpServer == nil {
		serverLock.Lock()
		defer serverLock.Unlock()
		if httpServer == nil{
			httpServer = &server{
				engine: gin.Default(),
			}
		}
	}
	return httpServer
}

type server struct {
	engine *gin.Engine
}

func (s *server) registerRouters() {
	RegisterBackendRouterV1(s.engine)
	RegisterApiV1(s.engine)
}

func (s *server) Start() {
	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT","GET", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin","content-type","token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	s.registerRouters()
	log.Fatal(s.engine.Run(":5566"))

}


