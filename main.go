package main

import (
	"encoding/xml"
	"fmt"
	"sync"

	"github.com/guilhermecoelho/rssReader/handles"
	"github.com/guilhermecoelho/rssReader/models"
)

func main() {
	links := handles.SerializeLinkJson("Links.json")

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)

	go func() {
		defer wg.Done()

		for item := range links.Links {
			checkRssList(links.Links[item].Url)
		}
	}()
}

func checkRssList(url string) {
	byteXml, err := handles.GetXml(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var rss models.Rss
	xml.Unmarshal(byteXml, &rss)

	var rssOld models.Rss
	handles.ReadExistedXml(rss, &rssOld)

	if rssOld.Channel.LastBuildDate != rss.Channel.LastBuildDate {
		handles.CreateLocalXml(rss)
	}
}
