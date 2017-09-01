package main

import (
	"fmt"
	"time"
	. "service-manager"
	"encoding/json"
)

type ServerCfg struct {
	Test1 string `json:"test1"`
	Test2 string `json:"test2"`
}

func main() {
	s := new(ServerCfg)
	_, err := InitConfigCallFunction("appInfos" ,s ,[]string{
		"http://10.240.36.209:2379",
		"http://10.240.36.210:2379",
		"http://10.240.36.211:2379",
	}, callChange)
	if err != nil {
		fmt.Println("init config falure.")
	}
	for {
		//fmt.Println(s)
		time.Sleep(time.Second * 2)
	}
}

type ThirdInfo struct {
	TwitterApiKey		string	`json:"twitterApiKey"`
	TwitterApiSecret	string	`json:"twitterApiSecret"`
	FacebookClientId	string	`json:"facebookClientId"`
	GoogleClientId		string	`json:"googleClientId"`
}

type ThirdInfoSlice struct {
	ThirdInfoSlice  ThirdInfo   `json:"thirdInfo"`
}

func callChange(abc interface{}) {
	fmt.Println("call change function============================================")
	appInfoMap := make(map[string]ThirdInfoSlice)
	jsonStr, err := json.Marshal(abc)
	if err != nil {
		fmt.Println("error =============")
	}
	if err := json.Unmarshal(jsonStr, &appInfoMap); err != nil {
		fmt.Println(err)
	}
	for k, v := range appInfoMap {
		fmt.Println(k)
		fmt.Println(v)
	}

}
