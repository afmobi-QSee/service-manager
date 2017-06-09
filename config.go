package service_manager

import (
	"log"
	"sync"
	"time"

	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"strings"
	"errors"
)

const CONFIG_ROOT = "/config/"

type Config struct {
	sync.RWMutex
	kapi          client.KeysAPI
	serviceName   string
	serviceStruct interface{}
	items         map[string]interface{}
}

func InitConfig(serviceName string, serviceStruct interface{}, endpoints []string) (*Config, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		HeaderTimeoutPerRequest: time.Second * 2,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	Config := &Config{
		kapi:          client.NewKeysAPI(c),
		serviceName:   serviceName,
		serviceStruct: serviceStruct,
		items:         make(map[string]interface{}),
	}

	Config.fetch(CONFIG_ROOT + serviceName, Config.items)
	Config.reload()

	/// `fetch` Timer may work well too?
	go Config.watch(CONFIG_ROOT + serviceName)

	return Config, err
}


func (cfg *Config) getItemKey(path string)string  {
	subs := strings.Split(path,"/")
	return subs[len(subs)-1]
	//idx := strings.Index(key,CONFIG_ROOT+cfg.serviceName)
	//return key[idx:]
}

func (cfg *Config) getItems(path string) (map[string]interface{}, error)  {
	var result  map[string]interface{}

	subs := strings.Split(path,"/")
	//check
	if(len(subs) < 3){
		return nil,errors.New("error path")
	}
	if(subs[1] != "config"){
		return nil,errors.New("error path, root path is not [/conifg].")
	}
	if(subs[2] != cfg.serviceName){
		return nil,errors.New("error path, serviceName is error.")
	}

	for i, item := range subs{
		if(i == 2){
			result = cfg.items
		}
		if (i > 2){
			if(i == (len(subs)-1)){
				break
			}
			if(result[item] == nil){
				tempItem := make(map[string]interface{})
				result[item] = tempItem
				result = tempItem
			}else {
				res,ok := result[item].(map[string]interface{})
				if(ok){
					result = res
				}
			}
		}
	}
	fmt.Println(result)
	return result, nil
}

func (cfg *Config) reload() {
	jsonb, _ := json.Marshal(cfg.items)
	fmt.Println(string(jsonb))
	err := json.Unmarshal(jsonb,cfg.serviceStruct)
	if(err != nil){
		fmt.Println(err.Error())
	}
	fmt.Println(cfg.serviceStruct)
}

func (cfg *Config) fetch(path string,items map[string]interface{}) error {
	resp, err := cfg.kapi.Get(context.Background(), path, nil)
	if err != nil {
		return err
	}
	if resp.Node.Dir {
		for _, v := range resp.Node.Nodes {
			if v.Dir {
				nKey := cfg.getItemKey(v.Key)
				newitem := make(map[string]interface{})
				items[nKey] = newitem
				cfg.fetch(v.Key,newitem)
			}else {
				nKey := cfg.getItemKey(v.Key)
				items[nKey] = v.Value
			}
		}
	}
	return nil
}

func (cfg *Config) updateItem(node client.Node) {
	nKey := cfg.getItemKey(node.Key)
	items,err := cfg.getItems("")
	if(err != nil){
		fmt.Println(err.Error())
	}
	items[nKey] = node.Value
}

func (cfg *Config) watch(path string) {
	watcher := cfg.kapi.Watcher(path, &client.WatcherOptions{
		Recursive: true,
	})
	for {
		resp, err := watcher.Next(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		switch resp.Action {
		case "set", "update":
			nKey := cfg.getItemKey(resp.Node.Key)
			items,err := cfg.getItems(resp.Node.Key)
			if(err != nil){
				fmt.Println(err.Error())
			}
			items[nKey] = resp.Node.Value
			fmt.Print("update-config---: key: "+ nKey + " value: " + resp.Node.Value)
			break
		case "expire", "delete":
			nKey := cfg.getItemKey(resp.Node.Key)
			items,err := cfg.getItems(resp.Node.Key)
			if(err != nil){
				fmt.Println(err.Error())
			}
			delete(items,nKey)
			fmt.Print("delete-config---: key: "+ nKey + " value: " + resp.Node.Value)
			break
		default:
			log.Println("watchme!!!", "resp ->", resp)
		}

		cfg.reload()
	}
}




