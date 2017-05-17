package main

import (
	"log"
	sm "service-manager"
	"fmt"
	"flag"
	"os"
	"os/signal"
)

func main() {

	var port = flag.Int("port", 7777, "listen port")
	flag.Parse()

	fmt.Println(sm.GetIP())
	w, err := sm.NewRegister("login", "", *port, []string{
		"http://192.168.18.100:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Register()


	c := make(chan os.Signal)
	signal.Notify(c)
	//监听指定信号
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2)
	//阻塞直至有信号传入
	s := <-c
	fmt.Println("get signal:", s)
}