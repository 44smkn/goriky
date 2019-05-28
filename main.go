package main

import (
	flags "github.com/jessevdk/go-flags"
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

type Options struct {
	LogfilePath string `short:"f" long:"logfile-path" description:"create requests from specificated logfile"`
	Logformat string `short:"l" long:"logformat" description:"format of logfile"`
	ReqPerMinute int  `short:"r" long:"reqPerMinute" description:"request count per one minute"`
}


func main() {

	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	f, err := os.Open(opts.LogfilePath)
	if err != nil {
		log.Fatalln("reading file is failed because of " + err.Error())
	}
	defer f.Close()

	format := opts.Logformat
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

	for _, req := range reqs {
		ctx := context.Background()
		c.SendRequest(ctx, req)
	}

	t := time.NewTicker(time.Duration(60/opts.ReqPerMinute) * time.Second)
	defer t.Stop()
	i := 0
	for range t.C {
		if i <= len(reqs) {
			i = 0
		}
		ctx := context.Background()
		res, _ := client.SendRequest(ctx, reqs[i])
		fmt.Println(res)
		i++
	}
}
