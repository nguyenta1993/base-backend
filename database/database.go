package database

import (
	"base_service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ReadDb *sqlx.DB
type WriteDb *sqlx.DB

func Open(cfg *config.DatabaseConfig, logger logger.Logger) (ReadDb, WriteDb) {
	readDb, readErr := sqlx.Connect(cfg.ReadDbCfg.DbType, cfg.ReadDbCfg.ConnectionString)
	writeDb, writeErr := sqlx.Connect(cfg.WriteDbCfg.DbType, cfg.WriteDbCfg.ConnectionString)

	if readErr != nil {
		logger.Info("ReadDb Connection String", zap.String("ConnectionString", cfg.ReadDbCfg.ConnectionString))
		logger.Info("WriteDb Connection String", zap.String("ConnectionString", cfg.WriteDbCfg.ConnectionString))
		logger.Error("Error Opening Read DB", zap.Error(readErr), zap.String("ReadDb Connection String", cfg.ReadDbCfg.ConnectionString), zap.String("WriteDb Connection String", cfg.WriteDbCfg.ConnectionString))
		panic(readErr)
	}

	if writeErr != nil {
		logger.Error("Error Opening Write DB", zap.Error(writeErr))
		panic(writeErr)
	}

	logger.Info("Connected to read & write database!")

	return readDb, writeDb
}
