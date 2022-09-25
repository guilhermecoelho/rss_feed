package handles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/guilhermecoelho/rssReader/models"
)

func SerializeLinkJson(filePath string) models.Links {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var links models.Links
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &links)
	return links
}

func ChangeLinkLastUpdate(url string, lastDate string) (string, error) {

	layout := time.RFC1123
	returnLayout := "2006-01-02 15:04:05"

	//convert lastDate format
	date, err := time.Parse(layout, lastDate)
	if err != nil {
		return date.Format(returnLayout), err
	}

	links := SerializeLinkJson("..\\files_test\\Links_test.json")

	var lastUpdateFromJson string
	for item := range links.Links {
		if url == links.Links[item].Url {
			lastUpdateFromJson = links.Links[item].LastUpdate
			break
		}
	}

	return lastUpdateFromJson, nil
}
