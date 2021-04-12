package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	IsAlive = true
)

func L7Check(c *gin.Context) {
	if IsAlive {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusNotFound)
	}

}

func Disable() {
	IsAlive = false
}
