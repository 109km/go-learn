package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	date := 821
	chapterId := 647
	sectionId := 1

	i := 0
	j := 0
	k := 0
	for i < 11 {
		for j < 99 {
			for k < 30 {
				getUrl(date+i, chapterId+j, sectionId+k)
				k = k + 1
				time.Sleep(time.Microsecond)
			}
			k = 0
			j = j + 1
		}
		k = 0
		j = 0
		i = i + 1
	}
}

func getUrl(date int, chapterId int, sectionId int) string {

	d := strconv.Itoa(date)
	c := strconv.Itoa(chapterId)
	s := strconv.Itoa(sectionId)

	url := "http://video.feierlaiedu.com/kcschool/20190" + d + "0" + c + "_" + s + ".mp4"

	// fmt.Println(d + " " + c + " " + s)
	res, err := http.Get(url)
	if err != nil {
		errMsg := url + "出现错误！"
		fmt.Println(errMsg)
		return errMsg
	}

	if res.StatusCode == 200 {
		fmt.Println(url)
		return url
	}
	return ""
}
