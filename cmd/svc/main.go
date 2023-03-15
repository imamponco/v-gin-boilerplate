package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/imamponco/v-gin-boilerplate/docs"
	"github.com/imamponco/v-gin-boilerplate/src/pkg/vtype"
	"github.com/imamponco/v-gin-boilerplate/src/svc"
	"github.com/imamponco/v-gin-boilerplate/src/svc/contract"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/nbs-go/nlogger"
	"os"
	"time"
)

var log = nlogger.Get()

// @title           Boilerplate - Service Using Gin Framework
// @version         1.0
// @description     RBoilerplate - Service Using Gin Framework.

// @contact.name   Imam Taufiq Ponco Utomo
// @contact.url
// @contact.email  imamtaufiqponco@gmail.co.id

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @schemes http

func main() {
	// Start cmd
	startedAt := time.Now()

	// Config load
	config := new(contract.Config)
	err := envconfig.Process("", config)
	if err != nil {
		panic(err)
	}

	// Init cmd contract
	app, err := svc.NewAPI(config)
	if err != nil {
		panic(err)
	}

	// Check if migration option is set
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = bootMigration(workDir, config)
	if err != nil {
		panic(err)
	}

	// Booting cmd
	err = app.Boot()
	if err != nil {
		panic(err)
	}

	// Init router
	router, err := app.InitRouter(workDir)
	if err != nil {
		panic(err)
	}

	if config.ServerSecure {
		ServeTLS(router, config, startedAt)
	} else {
		// Serve application
		log.Debugf("Boot time: %s", time.Since(startedAt))
		port := fmt.Sprintf(":%s", vtype.ParseStringFallback(config.Port, "8000"))
		err = router.Run(port)
		if err != nil {
			log.Errorf("%s", err.Error())
			log.Fatal("failed to serve cmd non tls.")
			os.Exit(2)
		}
	}
}

func ServeTLS(router *gin.Engine, config *contract.Config, startedAt time.Time) {
	// Serve application secure
	log.Debugf("Boot time: %s", time.Since(startedAt))
	port := fmt.Sprintf(":%s", vtype.ParseStringFallback(config.Port, "8000"))
	err := router.RunTLS(port, fmt.Sprintf("./configs/%s", config.ServerCert), fmt.Sprintf("./configs/%s", config.ServerCertKey))
	if err != nil {
		log.Errorf("%s", err.Error())
		log.Fatal("failed to serve cmd tls.")
		os.Exit(2)
	}
}
