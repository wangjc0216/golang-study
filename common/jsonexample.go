package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/alertmanager/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Chart struct {
	ID          int             `json:"id,omitempty" db:"id"`
	Name        string          `json:"name,omitempty" db:"name"`
	Type        string          `json:"type,omitempty" db:"type"`
	DashboardID int             `json:"dashboard_id,omitempty"`
	SourceType  string          `json:"source_type,omitempty" db:"source_type"`
	Data        json.RawMessage `json:"graph_data,ommitempty"`
}

func main2() {
	chart := Chart{}
	chart.ID = 1
	chart.Name = "Jishnu"
	str, err := json.Marshal(chart)
	fmt.Println(err)
	fmt.Println(string(str))
}

func main1(){

	bs, err := ioutil.ReadFile("/tmp/alertmanager.yaml")
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


	justPrint(cfg.Receivers)

}
type resp struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}
func justPrint(i interface{}){

	bs,err := json.Marshal(&resp{
		Status: "success",
		Data:   i,
	})
	if err!=nil{
		fmt.Println("json Marshal error:",err)
	}
	fmt.Println(string(bs))
}