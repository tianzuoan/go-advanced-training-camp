package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"week02/global"
	"week02/pkg/setting"
)

type Model struct {
}

func NewDBEngine(mysqlSetting *setting.MysqlSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		mysqlSetting.UserName,
		mysqlSetting.Password,
		mysqlSetting.Host,
		mysqlSetting.DBName,
		mysqlSetting.Charset,
		mysqlSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == setting.SERVER_RUN_MODE_DEBUG {
		//db.Logger.LogMode()
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(30)
	return db, nil
}
