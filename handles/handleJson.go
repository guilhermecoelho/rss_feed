package handles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/guilhermecoelho/rssReader/models"
)

func GetLinks() []string {

	var urls []string

	jsonFile, err := os.Open("Links.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var links models.Links
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &links)

	for item := range links.Links {
		urls = append(urls, links.Links[item].Url)
	}

	return urls
}
