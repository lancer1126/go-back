package core

import (
	"fmt"
	"go-back/global"
	"go-back/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	switch global.GB_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.GB_CONFIG.Mysql
	if m.Dbname == "" {
		fmt.Printf("Dbname is null")
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// RegisterTables 将model的变动自动映射到数据库中
func RegisterTables() {
	db := global.GB_DB
	err := db.AutoMigrate(
		system.SysApi{},
	)
	if err != nil {
		global.GB_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GB_LOG.Info("register table success")
}
