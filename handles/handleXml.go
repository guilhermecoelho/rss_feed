package handles

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/guilhermecoelho/rssReader/models"
)

func GetXml(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("GET error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("read body: %v", err)
	}

	return data, nil
}

func ReadExistedXml(rss models.Rss, rssOld *models.Rss) {

	var filename = "files/" + rss.Channel.Title + ".xml"
	filename = strings.Replace(filename, " ", "_", -1)
	oldXml, err := os.Open(filename)
	if err != nil {
		CreateLocalXml(rss)
	}
	defer oldXml.Close()
	byteValue, _ := ioutil.ReadAll(oldXml)

	xml.Unmarshal(byteValue, rssOld)
}

func CreateLocalXml(rss models.Rss) {

	filename := "files/" + rss.Channel.Title + ".xml"
	filename = strings.Replace(filename, " ", "_", -1)
	file, _ := os.Create(filename)

	xmlWrite := io.Writer(file)

	enc := xml.NewEncoder(xmlWrite)
	enc.Indent(" ", "   ")
	enc.Encode(rss)
}
