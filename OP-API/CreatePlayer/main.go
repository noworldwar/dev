package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/xid"
)

type Info struct {
	Filename string
}

func main() {
	info := &Info{Filename: "log/" + time.Now().Format("2006-01-02-150405-CreatePlayer.txt")}
	count := 500
	ch1 := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	defer close(ch1)
	go func() {
		defer wg.Done()
		// ch1 <- sendPost()
		for i := 0; i < count; i++ {
			ch1 <- sendPost()
		}

	}()

	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			sendLog(ch1, &info.Filename)
		}

	}()

	wg.Wait()
}

func sendPost() string {
	operatorID := "ta2admin1rtsg"
	appSecret := "zHlj6OVedD2jXLh1DafUKGxAWR5GyVkKaquRGoOSEy4="
	requestTime := strconv.FormatInt(time.Now().Unix(), 10)
	playerID := generatePlayer()
	// set operator
	formData := map[string]string{
		"operatorID":  operatorID,
		"appSecret":   appSecret,
		"playerID":    playerID,
		"nickname":    playerID,
		"requestTime": requestTime,
	}

	resp := postWithSignature("https://release-op-api.dvweg.com/player/create", formData, appSecret+playerID+operatorID+playerID+requestTime)
	return string(resp)
}

func generatePlayer() string {
	// Get a unique xid
	id := xid.New()
	return strings.ToUpper(id.String())
}

func postWithSignature(path string, formData map[string]string, signStr string) []byte {
	httpClient := resty.New()
	has := md5.Sum([]byte(signStr))
	signature := fmt.Sprintf("%x", has)
	rsp, err := httpClient.R().EnableTrace().
		SetHeader("signature", signature).
		SetFormData(formData).
		Post(path)
	if err != nil {
		return nil
	}
	return rsp.Body()
}

func printTextToFile(text, filename string) {
	// file
	var f *os.File

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(text + "\r\n")
}

func sendLog(ch1 chan string, filename *string) {
	printTextToFile(<-ch1, *filename)
}
