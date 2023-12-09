package daylight

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Torebekov/shop-back/config"
	"github.com/Torebekov/shop-back/internal/controller"
	core "github.com/Torebekov/shop-back/internal/core"
	"github.com/Torebekov/shop-back/modules/logger"
	"log"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func Start(ctx context.Context) {
	var err error

	defer func() {
		r := recover()
		if r != nil {
			var err error
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("unknown error")
			}
			// sendMeMail(err)
			log.Fatalln(err)
		}
	}()

	cfg := config.Load()
	if cfg == nil {
		os.Exit(1)
	}

	logger.Init(
		cfg.IsDevelopment,
		&logger.Config{
			EnableConsole:     true,
			ConsoleJSONFormat: false,
			ConsoleLevel:      cfg.LogLevel,
		},
	)

	sqlDB, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		logger.WorkLogger.Error("Failed close DB connection", zap.Error(err))
		os.Exit(1)
	}
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			logger.WorkLogger.Error("Failed close DB connection", zap.Error(err))
			os.Exit(1)
		}
	}()

	err = sqlDB.PingContext(ctx)
	if err != nil {
		logger.WorkLogger.Error("Failed close DB connection", zap.Error(err))
		os.Exit(1)
	}

	logger.WorkLogger.Info("server started")

	srv := controller.NewServer(
		ctx,
		core.New(
			ctx,
			cfg,
			sqlDB,
		),
	)
	err = srv.Run("8080")
	if err != nil {
		log.Fatalln("Failed to run service: ", err)
	}
}
