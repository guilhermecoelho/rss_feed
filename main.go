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

		for _, link := range links.Links {
			rss := checkRssList(link.Url)

			totalitems := len(rss.Channel.Items)
			totalItemsJson, _ := handles.GetTotalItemsJson(link.Url, "Links.json")

			totalNewItems := totalitems - totalItemsJson

			totalCountItems := totalNewItems
			if totalCountItems > 5 {
				totalCountItems = 5
			}
			for count := 0; count < totalCountItems; count++ {
				fmt.Println(rss.Channel.Items[count].Title)
			}

			handles.ChangeLinkTotalItems(link.Url, totalitems, "Links.json")
		}
	}()
}

func checkRssList(url string) models.Rss {
	var rss models.Rss
	byteXml, err := handles.GetXml(url)
	if err != nil {
		fmt.Println(err)
		return rss
	}

	xml.Unmarshal(byteXml, &rss)

	return rss
}
