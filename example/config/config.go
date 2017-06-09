package main

import (
	"fmt"
	"time"
	. "service-manager"
)

type ServerCfg struct {
	Test1 string `json:"test1"`
	Test2 string `json:"test2"`
}

func main() {
	s := new(ServerCfg)
	_, err := InitConfig("test" ,s ,[]string{
		"http://10.240.36.31:2379",
	})
	if err != nil {
		fmt.Println("init config falure.")
	}
	for {
		//fmt.Println(s)
		time.Sleep(time.Second * 2)
	}
}
