package main

import (
	"log"
	"time"

	discover "service-manager"
)

func main() {
	m, err := discover.NewMaster("service_d", []string{
		"http://192.168.18.100:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("all ->", m.GetNodes())
		log.Println("all(strictly) ->", m.GetNodesStrictly())
		time.Sleep(time.Second * 2)
	}
}