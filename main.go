package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanhua/ZanaoAPI/api/public"
	"github.com/jeanhua/ZanaoAPI/config"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetCfonfig()
	r.GET("/ping", public.Ping)
	r.GET("/home", public.HomePage)
	r.GET("/hot", public.Hot)
	r.Run("0.0.0.0:" + cfg.GetString("port"))
}
