package main

import (
	"fmt"
	"github.com/prometheus/alertmanager/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {

	bs, err := ioutil.ReadFile("./alertmanager.yaml")
	if err != nil {
		fmt.Println("ioutil.ReadFile error:%v", err)
		return
	}
	cfg := &config.Config{}

	err = yaml.UnmarshalStrict(bs, cfg)
	if err != nil {
		fmt.Println("yaml unmarshal error:", err)
		return
	}
	fmt.Println(cfg)
}

