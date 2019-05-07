package apache

import (
	"net/http"
)

type LogParser struct{}

func (p *LogParser) Parse(logs string) (*http.Request, error) {
	return nil, nil
}
