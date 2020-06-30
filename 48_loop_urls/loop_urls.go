package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var gRWLock *sync.RWMutex

func writeToFile(file *os.File, msg string) {
	_, err := file.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
}

func getUrl(date int, chapterId int, sectionId int, logFile *os.File, errFile *os.File) {

	gRWLock = new(sync.RWMutex)

	d := strconv.Itoa(date)
	c := strconv.Itoa(chapterId)
	s := strconv.Itoa(sectionId)

	url := "http://video.feierlaiedu.com/kcschool/20190" + d + "0" + c + "_" + s + ".mp4"
	res, err := http.Get(url)

	if err != nil {
		writeToFile(errFile, url+"\n")
		return
	}

	if res.StatusCode == 200 {
		writeToFile(logFile, url+"\n")
	} else {
		writeToFile(errFile, url+"\n")
	}
}

func main() {

	logFile, err1 := os.OpenFile("/Users/xinhengs/Develop/learn-go-by-examples/48_loop_urls/urls.log", os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_SYNC, os.ModePerm)
	errFile, err2 := os.OpenFile("/Users/xinhengs/Develop/learn-go-by-examples/48_loop_urls/error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)

	if err1 != nil {
		fmt.Println(err1.Error())
	}
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	date := 800
	chapterId := 600
	sectionId := 1

	i := 0
	j := 0
	k := 0
	for i < 21 {
		for j < 99 {
			for k < 20 {
				getUrl(date+i, chapterId+j, sectionId+k, logFile, errFile)
				k = k + 1
				// time.Sleep(time.Microsecond)
			}
			k = 0
			j = j + 1
		}
		k = 0
		j = 0
		i = i + 1
	}

	defer logFile.Close()
	defer errFile.Close()
	fmt.Println("End")
}
