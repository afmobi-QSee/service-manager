package main

import (
	"log"
	"time"

	discover "service-manager"
	"fmt"
)

func main() {
	m, err := discover.NewDiscovery("login", []string{
		"http://192.168.18.100:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("all ->", m.GetNodes())
		time.Sleep(time.Second * 2)
	}
}