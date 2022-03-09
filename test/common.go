package test

import (
	"carefree/app/daemon"
	"carefree/app/middleware"
	"carefree/app/validation"
	"carefree/router"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/dbdao/gormdb"
	"github.com/carefreex-io/logger"
	"github.com/carefreex-io/xhttp"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	config.DefaultCustomOptions.Path = "../conf/conf.yaml"
	config.InitConfig()

	logger.InitLogger()

	if err := gormdb.InitDB(); err != nil {
		logger.Fatalf("mysql.InitDB failed: err=%v", err)
	}

	daemon.RunStartBeforeFn()

	http := xhttp.NewXHttp()

	validation.RegisterValidation()

	middleware.Register(http.Engine)

	router.Register(http.Engine)

	return http.Engine
}
