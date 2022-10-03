package handles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

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

func ChangeLinkTotalItems(url string, totalItems int, jsonFile string) (int, error) {

	totalNewItems := 0

	links := SerializeLinkJson(jsonFile)
	for item := range links.Links {
		if url == links.Links[item].Url {
			totalItemJson, err := strconv.Atoi(links.Links[item].TotalItens)
			if err != nil {
				fmt.Println(err)
			}
			if totalItems > totalItemJson {
				links.Links[item].TotalItens = strconv.Itoa(totalItems)
				UpdateJson(links, jsonFile)
				totalNewItems = totalItems - totalItemJson
				break
			}
		}
	}
	return totalNewItems, nil
}

func UpdateJson(links models.Links, jsonFile string) {
	content, err := json.Marshal(links)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(jsonFile, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateJsonFile(link models.Links, fileName string) {
	content, err := json.Marshal(link)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
