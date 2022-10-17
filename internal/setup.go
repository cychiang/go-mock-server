package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter(conf *Config) *gin.Engine {
	r := gin.Default()
	initRouters(r, conf)
	return r
}

func initRouters(r *gin.Engine, conf *Config) {
	for _, router := range conf.Routers {
		switch strings.ToLower(router.Method) {
		case "get":
			r.GET(router.Path, func(c *gin.Context) {
				c.Data(http.StatusOK, "application/json", []byte(router.Body))
			})
		}
	}
}
