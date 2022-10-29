package server

import (
	"context"
	"database/sql"
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

	log, err := setUpLogger(cfg.AppConfig.LogFile)
	if err != nil {
		return nil, err
	}

	log.Infof("Loaded config: %+v", cfg)

	db, err := sql.Open("postgres", cfg.DBconn)
	if err != nil {
		panic(err)
	}

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
		Addr:    s.Cfg.AppConfig.Port,
	}

	s.Logger.Infof("Running on: %s", s.Cfg.AppConfig.Port)

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
