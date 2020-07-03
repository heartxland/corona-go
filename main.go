package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	GetData()
}

func GetData() {
	fmt.Println("get data！")
	sourceUrl := "https://dashboards-dev.sprinklr.com/data/9043/global-covid19-who-gis.json"
	res, err := http.Get(sourceUrl)
	if err != nil {
		panic(err)
	}
	byteArray, _ := ioutil.ReadAll(res.Body)
}
