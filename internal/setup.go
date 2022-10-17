package internal

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func SetupRouter(conf *Config) *gin.Engine {
	r := gin.Default()
	initRouters(r, conf)
	return r
}

func initRouters(r *gin.Engine, conf *Config) {
	for _, rt := range conf.Routers {
		switch strings.ToLower(rt.Method) {
		case "get":
			initGet(r, &rt)
		}
	}
}

func initGet(r *gin.Engine, rt *Router) {
	r.GET(rt.Path, func(c *gin.Context) {
		c.Data(rt.StatusCode, rt.ContentType, []byte(rt.Body))
	})
}
