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

type PSignInfo struct {
	PSign		[]string	`json:"pSign"`
}

type SignInfoSlice struct {
	SignInfoSlice   map[string]PSignInfo     `json:"signInfo"`
}

func callChange(abc interface{}) {
	fmt.Println("call change function============================================")
	//appInfoMap := make(map[string]ThirdInfoSlice)
	jsonStr, err := json.Marshal(abc)
	if err != nil {
		fmt.Println("error =============")
	}
	//if err := json.Unmarshal(jsonStr, &appInfoMap); err != nil {
	//	fmt.Println(err)
	//}
	//for k, v := range appInfoMap {
	//	fmt.Println(k)
	//	fmt.Println(v)
	//}

	signInfoMap := make(map[string]SignInfoSlice)
	if err := json.Unmarshal(jsonStr, &signInfoMap); err != nil {
		fmt.Println(err)
	}
	//appId := "123456"
	//pName := "com.afmobi.tudc"
	//pSign := "3B2B19B88552719822907F0AD1C964EA1835AC45"
	//for k, v := range signInfoMap {
	//	fmt.Println(k)
	//	fmt.Println(v.SignInfoSlice)
	//	if k == appId {
	//		fmt.Println("appid ======= true")
	//		for f, g := range v.SignInfoSlice {
	//			if f == pName {
	//				fmt.Println("==================")
	//				for _, r := range g.PSign {
	//					if pSign == r {
	//						fmt.Println("dddddddddddddd")
	//					}
	//				}
	//			}
	//		}
	//	}
	//
	//}

	appInfoMap := make(map[string][]Pinfo)
	for k, v := range signInfoMap {
		var pInfos []Pinfo
		for i, j := range v.SignInfoSlice {
			pInfos = append(pInfos, Pinfo{PName:i, PSign:j.PSign})
		}
		appInfoMap[k] = pInfos
	}
	fmt.Println(appInfoMap)
}

type Pinfo struct {
	PName 			string
	PSign			[]string
}
