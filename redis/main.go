package main

import (
	"flag"

	"github.com/golang/glog"

	_ "rw/go-exercise/redis/cmd"
	_ "rw/go-exercise/redis/db"
)

func main() {
	flag.Parse()
	glog.Info("test end")
}
