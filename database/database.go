package database

import (
	"base_service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogovan-korea/ggx-kr-service-utils/database"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"

	"github.com/jmoiron/sqlx"
)

type ReadDb *sqlx.DB
type WriteDb *sqlx.DB

func Open(cfg *config.DatabaseConfig, logger logger.Logger) (ReadDb, WriteDb) {
	readDb := database.MustConnect(cfg.ReadDbCfg.DbType, cfg.ReadDbCfg.ConnectionString)
	writeDb := database.MustConnect(cfg.WriteDbCfg.DbType, cfg.WriteDbCfg.ConnectionString)
	logger.Info("Connected to read & write database!")
	return readDb, writeDb
}
