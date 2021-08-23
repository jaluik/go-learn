package real

import (
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}

func (r *Retriever) NewGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	return string(response)
}
