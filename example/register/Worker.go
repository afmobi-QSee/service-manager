package main

import (
	"log"
	"time"

	sm "service-manager"
)

func main() {

	w, err := sm.NewRegister("login", "192.168.18.111", 7777, []string{
		"http://192.168.18.100:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Register()
	for {
		//log.Println("isActive ->")
		time.Sleep(time.Second * 2)
	}
}