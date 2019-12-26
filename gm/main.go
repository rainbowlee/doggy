package main

import (
	"fmt"
	"net/http"
	"flag"
	"time"
	"io"
    "html/template"
	"github.com/golang/glog"
	"os"
	"strconv"
)


var t = template.Must(template.ParseGlob("views/*"))

type MyHandler struct{

}

func (*MyHandler) ServeHTTP( w http.ResponseWriter, r *http.Request){
	fmt.Println("MyHandler ServeHTTP")
	//w.WriteContent("aaaaaaa")

	switch r.URL.Path{
	case "/a":
		//io.WriteString(w, "a!\n")
		glog.Info("a request")
		renderTemplate(w, "a.html", "hello a!")
		break;
	case "/b":
		glog.Info("b request")	
		io.WriteString(w, "b!\n")
		break;
	default:
		io.WriteString(w, "Hello, TLS!\n")
	}

	/*
	if r.URL.Path() == "a"{
		io.WriteString(w, "a!\n")
	} else if r.URL.Path() == "b"{
		io.WriteString(w, "b!\n")
	}
	*/



}

func main() {
	for idx, args := range os.Args {
        fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}
		
	flag.Parse()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        new(MyHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	defer glog.Flush()

	glog.Info("This is info message")
	glog.Infof("This is info message: %v", 12345)
	glog.InfoDepth(1, "This is info message", 12345)

	glog.Warning("This is warning message")
	glog.Warningf("This is warning message: %v", 12345)
	glog.WarningDepth(1, "This is warning message", 12345)

	glog.Error("This is error message")
	glog.Errorf("This is error message: %v", 12345)
	glog.ErrorDepth(1, "This is error message", 12345)

	glog.Fatal("This is fatal message")
	glog.Fatalf("This is fatal message: %v", 12345)
	glog.FatalDepth(1, "This is fatal message", 12345)	
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := t.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, "error 500:"+" "+err.Error(), http.StatusInternalServerError)
    }
}