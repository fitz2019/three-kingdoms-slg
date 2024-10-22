package main

import (
	"fmt"
	"os"

	"github.com/llr104/slgserver/config"
	"github.com/llr104/slgserver/net"
	"github.com/llr104/slgserver/server/chatserver"
)

func getChatServerAddr() string {
	host := config.File.MustValue("chatserver", "host", "127.0.0.1")
	//host := config.File.MustValue("chatserver", "host", "")
	port := config.File.MustValue("chatserver", "port", "8002")
	return host + ":" + port
}

func main() {
	fmt.Println(os.Getwd())
	chatserver.Init()
	needSecret := config.File.MustBool("chatserver", "need_secret", false)
	s := net.NewServer(getChatServerAddr(), needSecret)
	s.Router(chatserver.MyRouter)
	s.Start()
}
