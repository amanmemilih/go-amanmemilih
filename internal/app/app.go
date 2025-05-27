package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http"
	"github.com/zinct/amanmemilih/internal/wire"
	"github.com/zinct/amanmemilih/pkg/httpserver"
	"github.com/zinct/amanmemilih/pkg/jwt"
	"github.com/zinct/amanmemilih/pkg/logger"
	"github.com/zinct/amanmemilih/pkg/mysql"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)

	// Initialize MYSQL
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.MYSQL.Username, cfg.MYSQL.Password, cfg.MYSQL.Host, cfg.MYSQL.Port, cfg.MYSQL.Database)
	mysql, err := mysql.New(mysqlUrl, mysql.SetMaxIdleConns(cfg.MYSQL.PoolMax), mysql.SetMaxOpenConns(cfg.MYSQL.PoolMax), mysql.SetConnMaxLifetime(time.Duration(cfg.MYSQL.PoolMax)*time.Second))
	if err != nil {
		panic(err)
	}
	defer mysql.Close()

	// Initialize JWT
	jwtManager := jwt.New(cfg.JWT.Secret)

	// Initialize HTTP Server
	httpserver := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.WithNewGinEngine())

	http.RegisterMiddleware(httpserver.Router, cfg, log)
	http.RegisterRoutes(httpserver.Router, http.RouterOption{
		AuthController:        wire.InitializeAuthController(mysql.DB, cfg, log, jwtManager),
		ProvinceController:    wire.InitializeProvinceController(mysql.DB, cfg, log),
		DistrictController:    wire.InitializeDistrictController(mysql.DB, cfg, log),
		SubdistrictController: wire.InitializeSubdistrictController(mysql.DB, cfg, log),
		VillageController:     wire.InitializeVillageController(mysql.DB, cfg, log),
	}, cfg, log, jwtManager)

	httpserver.Start()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Info("internal/app - Run - signal: %s", s.String())
	case err := <-httpserver.Notify():
		log.Error(fmt.Errorf("internal/app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpserver.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("internal/app - Run - httpServer.Shutdown: %w", err))
	}
}
