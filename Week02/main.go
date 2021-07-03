package main

import (
	"log"
	"week02/global"
	"week02/internal/model"
	"week02/internal/routers"
	"week02/setup"
)

func init() {
	err := setup.SetupSetting()
	if err != nil {
		log.Fatal("load config error,", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatal("init db error,", err)
	}
}
func main() {
	err := routers.NewHttpServer()
	log.Fatal(err)
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.MysqlSetting)
	if global.ServerSetting.RunMode == "debug" {
		//global.DBEngine.Logger = global.DBEngine.Logger.LogMode(logger.Info)
		global.DBEngine = global.DBEngine.Debug()
	}
	return err
}
