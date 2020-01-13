package main

import (
	"flag"
	"fmt"
	//"log"
	"os"
	"os/signal"
    "syscall"
	"github.com/spf13/viper"
	//"github.com/panjf2000/gnet"	
)

var (
	endchan    chan bool   = make(chan bool)
)

func main() {
	var port int
	var multicore bool
	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9003, "--port 9000")
	flag.BoolVar(&multicore, "multicore", true, "--multicore true")
	flag.Parse()

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	datestart := viper.GetString("Default.Process_DateStart")
	fmt.Println("date start process ", datestart)

	redisaddr := viper.GetString("redis.addr")
	kahost := viper.GetString("kafka.host")

	RedisInit(redisaddr)
	InitKafuka(kahost)
	go BoottimeTimingSettlement()

	go InitServer(multicore, port, &endchan)

	c := make(chan os.Signal)
	signal.Notify(c,syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c
	endchan <- true
	CloseKafKa()
}

