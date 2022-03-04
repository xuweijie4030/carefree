package main

import (
	"carefree/app/daemon"
	"carefree/app/middleware"
	"carefree/router"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/dbdao/gormdb"
	"github.com/carefreex-io/logger"
	"github.com/carefreex-io/xhttp"
)

func main() {
	config.InitConfig()

	logger.InitLogger()

	if err := gormdb.InitDB(); err != nil {
		logger.Fatalf("mysql.InitDB failed: err=%v", err)
	}

	daemon.RunStartBeforeFn()

	http := xhttp.NewXHttp()

	middleware.Register(http.Engine)

	router.Register(http.Engine)

	http.Start()

	daemon.RunStoppedFn()
}
