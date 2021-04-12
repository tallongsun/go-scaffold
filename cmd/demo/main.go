package main

import (
	"flag"
	"fmt"
	"github.com/tallongsun/go-scaffold/pkg/controller"
	"github.com/tallongsun/go-scaffold/pkg/lib/config"
	"github.com/tallongsun/go-scaffold/pkg/lib/log"
	"github.com/tallongsun/go-scaffold/pkg/router"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	flag.Parse()
	fmt.Println("application is initializing")

	config.Init()
	fmt.Println("config initialized", "RUN_MODE", config.Config.Get("mode"))

	log.Init()
	fmt.Println("log initialized")
}

func main() {
	router.Start()
	fmt.Println("router is running")

	signalHandler()
	fmt.Println("application started")
}

func signalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Logger.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//graceful shutdown
			controller.Disable()
			router.Stop()
			log.Logger.Info("application stopped")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
