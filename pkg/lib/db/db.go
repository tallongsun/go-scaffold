package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tallongsun/go-scaffold/pkg/lib/config"
	"github.com/tallongsun/go-scaffold/pkg/model"
	"time"
)

var Engine *gorm.DB

func Init() {
	var err error
	Engine, err = gorm.Open("mysql", config.Config.GetString("mysql.datasource"))
	if err != nil {
		panic(err)
	}

	Engine.DB().SetMaxIdleConns(5)
	Engine.DB().SetMaxOpenConns(100)
	Engine.DB().SetConnMaxLifetime(time.Hour)

	mode := config.Config.GetString("mode")
	if mode == "alpha" {
		Engine.AutoMigrate(&model.User{})
	}

}

func Stop() {
	if Engine != nil {
		Engine.Close()
	}
}
