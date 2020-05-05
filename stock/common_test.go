package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/glog"
)

func TestTime(t *testing.T) {

	now := time.Now().Add(-24 * 60 * 60 * 1e9)
	month := now.Month()
	day := now.Day()

	//monthday := fmt.Sprintf("%02d-%02d", month, day)
	yearmonthday := fmt.Sprintf("%d-%02d-%02d", now.Year(), month, day)

	glog.Info("TestTime month ", month, " day ", day, " monthday ", yearmonthday)
}

func TestAdd(t *testing.T) {
	TestTime(t)
}
