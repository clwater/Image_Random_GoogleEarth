package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"strings"
)

type GoogleEarth struct {
	Index int
	Url string
}

var infos []GoogleEarth

func main() {
	index := 2000
	failIndex := 0

	for ;failIndex < 100 ;  {
		url := fmt.Sprintf("%s%d%s", "https://www.gstatic.com/prettyearth/assets/full/", index, ".jpg")
		log.Println(url)

		client := resty.New()
		resp, err := client.R().Get(url)
		//多维度判断页面内容是否有效
		if err != nil || resp.StatusCode() != 200 || resp.Size() < 2000{
			failIndex++
		}else {
			failIndex = 0
			infos = append(infos, GoogleEarth{index, url})
		}
		index++

		//每检索10此进行一次保存
		if index % 10 == 0 {
			//log.Println(infos)
			j, _ := json.Marshal(infos)
			log.Println(string(j))
			Save(string(j))
			infos = nil
		}

	}
	


}


//
//  Save
//  @Description: 追加相应内容到文件中
//  @param text
//
func Save(text string){
	var filename = "earthview.json"
	var f *os.File
	var err1 error
	f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	defer f.Close()
	if err1 != nil {
		panic(err1)
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	text = strings.ReplaceAll(text, "[", "")
	text = strings.ReplaceAll(text, "]", ",")
	text = text + "\n"
	_, _ = w.WriteString(text)
	//fmt.Printf("写入 %d 个字节n", n)
	w.Flush()
}

