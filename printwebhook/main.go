package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)



func main(){
	port := os.Getenv("webhook_port")
	http.HandleFunc("/",printHandle)
	http.ListenAndServe(port,nil)
}

func printHandle(w http.ResponseWriter,r *http.Request){
	bs,err := ioutil.ReadAll(r.Body)
	if err !=nil{
		fmt.Println(time.Now(),"ioutil ReadAll error:",err)
	}
	fmt.Println()
	fmt.Println(time.Now())
	fmt.Printf("%s\n",string(bs))
}