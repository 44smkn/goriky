package apache

import (
	"goriky/infrastructure/client"
	"path"
	"net/http"
	"strings"
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

	host := items["%h"]
	spath := items["%U"]
	method := items["%m"]
	c, err := client.New(path.Join(host, spath))
	if err != nil {
		return nil, err
	}
	u := *c.URL
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
        return nil, err
	}

	return req, nil
}
