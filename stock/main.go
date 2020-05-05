package main

import (
	"flag"

	"github.com/golang/glog"
	mysqldata "github.com/rainbowlee/doggy/stock/sql"
)

func main() {

	flag.Parse()
	defer glog.Flush()

	glog.Info("This is info message")
	walkdata := false
	flag.BoolVar(&walkdata, "gen", false, "this help")

	mysqldata.OrmConnectDB()

	if walkdata == true {
		FundIter()
	}

	//mysqldata.ConnectDB()
	//debt.Test()
	defer mysqldata.OrmCloseDB()
}
