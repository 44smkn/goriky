package main

import (
	"context"
	"fmt"
	"goriky/infrastructure/client"
	"goriky/accesslog/apache"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	b, err := ioutil.ReadFile("sample.log")
	if err != nil {
		log.Fatalln("reading file is failed because of " + err.Error())
	}
	accessLogs := strings.Split(string(b), "\n")
	reqPerMinute := 10
	client, _ := client.New("http://hogehoge.com")
	apacheLog := &apache.LogParser{}

	t := time.NewTicker(time.Duration(60/reqPerMinute) * time.Second)
	defer t.Stop()
	i := 0
	for range t.C {
		if i <= len(accessLogs) {
			i = 0
		}
		req, _ := apacheLog.Parse(accessLogs[i])
		ctx := context.Background()
		res, _ := client.SendRequest(ctx, req)
		fmt.Println(res)
	}
}
