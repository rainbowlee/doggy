package main

import (
	"github.com/panjf2000/gnet"
	"log"
	"fmt"
)

var(

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

func InitServer(multicore bool, port int, endchan* chan bool ){
	linecodec := new(gnet.LineBasedFrameCodec)
	echo := new(statServer)
	log.Fatal(gnet.Serve(echo, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore), gnet.WithCodec(linecodec)))
}
