package handles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/guilhermecoelho/rssReader/models"
)

func Test_ChangeLinkTotalItems(t *testing.T) {

	//Arrange
	url := "https://anchor.fm/s/4f366e84/podcast/rss"
	totalNewItems := 10
	totalOldItems := 4
	expectedResult := totalNewItems - totalOldItems

	jsonFile := "..\\files_test\\Links_test_ChangeLinkTotalItems.json"
	os.Remove(jsonFile)

	var links models.Links
	var link models.Link
	link.Url = "https://anchor.fm/s/4f366e84/podcast/rss"
	link.TotalItens = strconv.Itoa(totalOldItems)
	links.Links = append(links.Links, link)
	createLocalJsonFile(links, jsonFile)

	//Act
	result, err := ChangeLinkTotalItems(url, totalNewItems, jsonFile)

	//Assert
	if err != nil {
		t.Error("Erro when convert date: ", err)
	}
	if result != expectedResult {
		t.Error("Expected "+strconv.Itoa(expectedResult)+", but is: ", result)
	}
}

func Test_CreateJsonFile(t *testing.T) {

	//Arrange
	jsonFile := "..\\files_test\\Links_test_CreateJsonFile.json"
	os.Remove(jsonFile)

	var links models.Links

	var link models.Link
	link.Url = "https://hipsters.tech/feed/podcast/"
	link.TotalItens = "5"
	links.Links = append(links.Links, link)

	link.Url = "https://anchor.fm/s/4f366e84/podcast/rss"
	link.TotalItens = "10"
	links.Links = append(links.Links, link)

	//Act
	CreateJsonFile(links, jsonFile)
	result := serializeLinkJson(jsonFile)

	//Assert
	if len(result.Links) != 2 {
		t.Error("Expected 2 but is: ", len(result.Links))
	}
}

func serializeLinkJson(filePath string) models.Links {
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

func createLocalJsonFile(link models.Links, fileName string) {
	content, err := json.Marshal(link)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
