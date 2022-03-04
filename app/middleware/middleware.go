package middleware

import (
	"github.com/carefreex-io/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var middlewares = map[string]gin.HandlerFunc{
	"cors": cors.Default(),
	"jwt":  NewJwtMiddleware().Handle(),
	"demo": NewDemoMiddleware().Handle(),
}

func Register(r *gin.Engine) {
	for name, middleware := range middlewares {
		if name == "jwt" && config.GetString("Jwt.Switch") == "on" {
			r.Use(middleware)
			continue
		}
		if name == "cors" && config.GetString("Cors.Switch") == "on" {
			r.Use(middleware)
			continue
		}
		r.Use(middleware)
	}
}
