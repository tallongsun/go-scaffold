package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tallongsun/go-scaffold/pkg/controller"
	"github.com/tallongsun/go-scaffold/pkg/lib/config"
	"github.com/tallongsun/go-scaffold/pkg/lib/log"
	"net/http"
	"time"
)

var httpServer *http.Server

func Start() {
	gin.SetMode(config.Config.GetString("ginMode"))
	router := gin.Default()

	setupRoute(router)

	httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Config.GetString("port")),
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Logger.Errorf("router serve err: %v", err)
		}
	}()
}

func setupRoute(router *gin.Engine) {
	mn := router.Group("/monitor")
	mn.GET("/l7check", controller.L7Check)

	apiG := router.Group("/api")
	apiG.GET("/users", controller.ListUsers)
	apiG.GET("/users/:id", controller.GetUser)
	apiG.POST("/users", controller.CreateUser)
	apiG.PUT("/users/:id", controller.UpdateUser)
	apiG.DELETE("/users/:id", controller.DeleteUser)
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Logger.Errorf("router stop err: %v", err)
	}
}
