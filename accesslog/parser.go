package accesslog

import "net/http"

type Parser interface {
	Parse(log string) (http.Request, error)
}
