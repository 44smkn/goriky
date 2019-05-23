package main

import (
	"bufio"
	"context"
	"fmt"
	"goriky/accesslog/apache"
	"goriky/infrastructure/client"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.Open("sample.log")
	if err != nil {
		log.Fatalln("reading file is failed because of " + err.Error())
	}
	defer f.Close()

	format := ""
	p := apache.NewParser(format)
	c := client.New()
	buf := bufio.NewReaderSize(f, 4096) // may change buffer size...?
	reqs := make([]*http.Request, 10)
	for line := ""; err == nil; line, err = buf.ReadString('\n') {
		req, err := p.Parse(line)
		if err != nil {
			log.Fatalln("To gain the line is failed because of " + err.Error())
		}
		reqs = append(reqs, req)
	}
	if err != io.EOF {
		log.Fatalln("reading line is failed." + err.Error())
	}

	client := client.New()
	reqPerMinute := 10

	for _, req := range reqs {
		ctx := context.Background()
		c.SendRequest(ctx, req)
	}

	t := time.NewTicker(time.Duration(60/reqPerMinute) * time.Second)
	defer t.Stop()
	i := 0
	for range t.C {
		if i <= len(reqs) {
			i = 0
		}
		ctx := context.Background()
		res, _ := client.SendRequest(ctx, reqs[i])
		fmt.Println(res)
	}
}
