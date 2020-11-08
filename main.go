package main

import (
	"Go-Banana/glog"
	"Go-Banana/server"
)

func main() {
	glog.Debug("Go Banana!")
	server.Start()
}
