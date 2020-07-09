package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	covidData CovidData
)

const (
	sourceUrl          = "https://dashboards-dev.sprinklr.com/data/9043/global-covid19-who-gis.json"
	originalDataFolder = "original"
)

type (
	CovidData struct {
		Dimensions     []Dimensions
		Metrics        []Metrics
		Rows           []Rows
		Totals         []string
		License        string
		Copyright      string
		Created        string
		CreatedTime    string
		LastUpdateTime string
	}

	Dimensions struct {
		Name     string
		Type     string
		DataType string
	}

	Metrics struct {
		Name     string
		Type     string
		DataType string
	}

	Rows struct {
		row []string
	}
)

func main() {
	byteData := GetData()
	SaveOriginalData(byteData)
}

func GetData() []byte {
	fmt.Println("get data！")
	res, err := http.Get(sourceUrl)
	if err != nil {
		log.Fatal(err)
	}
	byteArray, _ := ioutil.ReadAll(res.Body)
	// convertErr := json.Unmarshal(byteArray, &covidData)
	// if err != nil {
	// 	panic(convertErr)
	// }
	return byteArray
}

func SaveOriginalData(byteData []byte) {
	if _, err := os.Stat(originalDataFolder); os.IsNotExist(err) {
		os.Mkdir(originalDataFolder, 0777)
	}

	currentDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(currentDir + "/" + originalDataFolder)
	if err != nil {
		log.Fatal(err)
	}

	originalFile, fpErr := os.Create(covidData.LastUpdateTime + ".json")
	if fpErr != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	originalFile.WriteString(string(byteData))

}
