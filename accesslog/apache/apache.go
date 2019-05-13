package apache

import (
	"regexp"
	"goriky/infrastructure/client"
	"path"
	"net/http"
	"strings"
	li "goriky/accesslog/apache/logitems"
)

type LogParser struct{
	logFormat []string
}

func NewParser(logFormat string) *LogParser {
	return &LogParser {
		logFormat: strings.Split(logFormat, " "),
	}
}

func (p *LogParser) Parse(line string) (*http.Request, error) {

	var items map[string]string
	for i, v := range strings.Split(line, " ") {
		items[p.logFormat[i]] = v
	}

	host := items[li.RemoteHost]
	spath := items[li.Path]
	method := items[li.Method]
	
	c, err := client.New(path.Join(host, spath))
	if err != nil {
		return nil, err
	}
	u := *c.URL
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
        return nil, err
	}

	for k, v := range items {
		re := regexp.MustCompile(li.RequestHeaderRe)
		if m := re.FindStringSubmatch(k); len(m) > 1 {
			req.Header.Add(m[1], v)
		}
	}

	return req, nil
}
