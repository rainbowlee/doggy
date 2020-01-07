package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
	"github.com/spf13/viper"
)

type statServer struct {
	*gnet.EventServer
}

func (es *statServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumLoops)
	return
}

func (es *statServer) React(c gnet.Conn) (out []byte, action gnet.Action) {
	// Echo synchronously.
	out = c.ReadFrame()
	if out != nil {
		//line data
		linedata := string(out)
		//log.Info("read dataline %s", linedata)
		log.Println("read dataline", linedata)
		Dispatch(linedata)
	}
	return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

func main() {
	var port int
	var multicore bool
	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", true, "--multicore true")
	flag.Parse()

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	datestart := viper.GetString("Default.Process_DateStart")
	fmt.Println("date start process ", datestart)

	redisaddr := viper.GetString("redis.addr")
	kahost := viper.GetString("kafka.host")

	RedisInit(redisaddr)
	InitKafuka(kahost)
	go BoottimeTimingSettlement()

	linecodec:=	new(gnet.LineBasedFrameCodec)
	echo := new(statServer)
	log.Fatal(gnet.Serve(echo, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore), gnet.WithCodec(linecodec)))
}