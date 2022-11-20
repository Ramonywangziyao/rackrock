package component

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	loggerV2 "gorm.io/gorm/logger"
	"rackrock/config"
	"rackrock/logger"
)

var DB *gorm.DB

func InitDatabase() {
	var cfg = config.Cfg.Db
	logger.Logger.Infof("connect db....")

	var db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: cfg.Dsn(),
	}), &gorm.Config{
		Logger: loggerV2.Default.LogMode(loggerV2.Info),
	})

	if err != nil {
		logger.Logger.Error("init db session err: %s", err.Error())
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		logger.Logger.Error("connect db err: %s", err.Error())
		return
	} else {
		sqlDb.SetMaxIdleConns(cfg.MaxIdleConn)
		sqlDb.SetMaxOpenConns(cfg.MaxOpenConn)
	}

	DB = db
}
