package main

import (
	"myapp/httpserv"
	"myapp/infra"
	"myapp/listen"
)

func init() {
	infra.InitConfig()
}

func main() {
	infra.InitWriter()
	infra.InitReader()
	httpserv.Run()
	listen.Run()
}
