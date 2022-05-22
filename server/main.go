package main

import (
	"flag"
	"github.com/Lmineor/mto/server"
	"github.com/golang/glog"
)

func main(){
	flag.Parse()
	glog.Info("start to mto backend")
	defer glog.Flush()

	server.Run()
}