package main

import (
	"go-workspace-hands-on/hello"
	"go-workspace-hands-on/server"
)

func main() {
	hello.Hello("kohei")
	server.StartServer()
}
