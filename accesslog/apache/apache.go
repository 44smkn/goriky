package apache

import (
	"goriky/infrastructure/client"
	"path"
	"fmt"
	"net/http"
	"regexp"
)

type LogParser struct{}

var logRe = regexp.MustCompile(
	`^(?:(\S+(?:,\s\S+)*)\s)?` + // %v(The canonical ServerName/virtual host) - 192.168.0.1 or 192.168.0.1,192.168.0.2, 192.168.0.3
		`(\S+)\s` + // %h(Remote Hostname) $remote_addr
		`(\S+)\s` + // %l(Remote Logname)
		`([\S\s]+)\s` + // $remote_user
		`\[(\d{2}/\w{3}/\d{2}(?:\d{2}:){3}\d{2} [-+]\d{4})\]\s` + // $time_local
		`(.*)`)

func (p *LogParser) Parse(line string) (*http.Request, error) {
	matches := logRe.FindStringSubmatch(line)
	if len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse apachelog (not matched): %s", line)
	}

	host := matches[2]
	spath := matches[1]
	method := matches[3]
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
