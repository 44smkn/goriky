package apache_test

import (
	"github.com/google/go-cmp/cmp"
	"goriky/accesslog/apache"
	"net/http"
	"testing"
)

func TestParse(t *testing.T) {
	a := apache.NewParser(`%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\"`)

	testCases := []struct {
		line string
		want http.Request
	}{
		{
			line: `192.168.0.101 - - [12/May/2014:20:41:48 +0900] "GET /index.html HTTP/1.1" 200 114 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:29.0) Gecko/20100101 Firefox/29.0"`,
			want: http.Request{
				Method: "GET",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.line, func(t *testing.T) {
			got, err := a.Parse(tC.line)
			if err != nil {
				t.Errorf(err.Error())
			}
			if diff := cmp.Diff(tC.want, got); diff != "" {
				t.Errorf("want = %v, got = %v", tC.want, got)
			}
		})
	}
}