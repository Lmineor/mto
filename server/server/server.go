package server

import (
	"fmt"
	"github.com/Lmineor/mto/config"
	"github.com/Lmineor/mto/global"
	"github.com/Lmineor/mto/models"
	"github.com/Lmineor/mto/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"time"
)


type server interface {
	ListenAndServe() error
}

func Run(){
	cfgPath := "./config.yaml"
	initConfig(cfgPath)
	initDB()

	if global.DB != nil {
		db,_ :=  global.DB.DB()
		defer db.Close()
		models.MigrateTables(global.DB)
	}

	runServer()
}

func initConfig(cfgPath string){
	cfg , err := config.ReadConfig(cfgPath)
	if err != nil{
		glog.Fatal("read config [%s] failed %s", cfgPath,err)
	}
	global.Config = cfg
}

func initDB(){
	db := models.GormMysql()
	global.DB = db
}

func runServer(){
	Router := router.Routers()
	address := fmt.Sprintf("127.0.0.1:%d", global.Config.System.Addr)
	s := initServer(address, Router)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	glog.V(1).Infof("server run success on %s", global.Config.System.Addr)

	fmt.Printf(`欢迎%s\n`, address)

	glog.Error(s.ListenAndServe().Error())
}


func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
