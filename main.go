package main

import (
	"fmt"
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
	t := time.NewTicker(int(60/reqPerMinute) * time.Second)
	defer t.Stop()
	for {
		for _ = range t.C {
			fmt.Println("req!")
		}
	}
}
