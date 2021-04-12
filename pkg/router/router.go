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

	mn := router.Group("/monitor")
	mn.GET("/l7check", controller.L7Check)

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

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Logger.Errorf("router stop err: %v", err)
	}
}
