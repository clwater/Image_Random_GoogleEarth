package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

type GoogleEarth struct {
	Index int
	Url string
}

var infos []GoogleEarth

func main() {
	index := 1000
	failIndex := 0

	for ;failIndex < 100 ;  {
		url := fmt.Sprintf("%s%d%s", "https://www.gstatic.com/prettyearth/assets/full/", index, ".jpg")
		log.Println(url)

		client := resty.New()
		resp, err := client.R().Get(url)
		if err != nil || resp.StatusCode() == 200 {
			failIndex++
		}else {
			failIndex = 0
			infos = append(infos, GoogleEarth{index, url})

		}
		index++
	}
	
	//log.Println(infos)

	j, _ := json.Marshal(infos)
	log.Println(string(j))

}
