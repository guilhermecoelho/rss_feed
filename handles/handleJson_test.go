package handles

import (
	"testing"
)

func Test_ChangeLinkLastUpdate(t *testing.T) {

	//Arrange
	url := "https://anchor.fm/s/4f366e84/podcast/rss"
	//lastDate := "Sat, 24 Sep 2022 09:03:47 GMT"
	lastDate := "Wed, 21 Sep 2022 17:58:53 +0000"
	formatedLastDate := "2022-09-21 17:58:53"

	//Act
	result, err := ChangeLinkLastUpdate(url, lastDate)

	//Assert
	if err != nil {
		t.Error("Erro when convert date: ", err)
	}
	if result != formatedLastDate {
		t.Error("Expected "+formatedLastDate+" but is: ", result)
	}

}
