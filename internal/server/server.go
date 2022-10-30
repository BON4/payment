package server

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BON4/payment/config"
	"github.com/BON4/payment/internal/middleware"
	tx_handlers "github.com/BON4/payment/internal/transaction/delivery/http"
	"github.com/BON4/payment/internal/transaction/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/BON4/payment/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setUpLogger(fileName string) (*logrus.Logger, error) {
	// instantiation
	logger := logrus.New()

	if len(fileName) == 0 {
		logger.Out = os.Stdout
	} else {
		//Write to file
		src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		//Set output
		logger.Out = src
	}

	//Set log level
	logger.SetLevel(logrus.DebugLevel)

	//Format log
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger, nil
}

func runDBMigration(migrationURL string, dbSource string) error {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		return err
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func waiForDb(waitSec int, driver, connectionString string) (*sql.DB, error) {
	var err error
	var db *sql.DB
	for ; waitSec > 0; waitSec-- {
		db, err = sql.Open(driver, connectionString)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
		continue
	}

	if db == nil {

		return nil, errors.New("Cant connect to DB")
	}

	return db, err
}

type Server struct {
	g      *gin.Engine
	Logger *logrus.Logger
	DB     *sql.DB
	Cfg    config.ServerConfig
}

func SetupHandlers(s *Server) {
	mdw := middleware.NewMiddleware(s.Logger.WithField("middleware", "middleware"))
	s.g.Use(mdw.CORS())

	tuc := usecase.NewTxUsecase(s.DB)

	tx_group := s.g.Group("payments")

	tx_handlers.NewTransactionHandler(tx_group, tuc, s.Logger.WithField("payments_handler", "payments_handler"))
}

func NewServer(configPath string) (*Server, error) {
	g := gin.Default()

	cfg, err := config.LoadServerConfig(configPath)
	if err != nil {
		return nil, err
	}

	log, err := setUpLogger(cfg.LogFile)
	if err != nil {
		return nil, err
	}

	log.Infof("Loaded config: %+v", cfg)

	db, err := waiForDb(60, cfg.DBDriver, cfg.DBconn)
	if err != nil {
		return nil, err
	}

	if err := runDBMigration(cfg.MirationURL, cfg.DBconn); err != nil {
		return nil, err
	}

	log.Infof("Database migrated")

	//Swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &Server{
		g:      g,
		Logger: log,
		DB:     db,
		Cfg:    cfg,
	}, nil
}

func (s *Server) Run() error {
	srv := &http.Server{
		Handler: s.g,
		Addr:    s.Cfg.Port,
	}

	s.Logger.Infof("Running on: %s", s.Cfg.Port)

	//Swagger
	//s.g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	SetupHandlers(s)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Errorf("listen: %s\n", err)
			return
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.Logger.Info("Shutdown Server ...")

	s.DB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Error("Server Shutdown Err:", err)
		return err
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		s.Logger.Info("timeout of 5 seconds.")
	}
	s.Logger.Info("Server exiting")
	return nil
}
