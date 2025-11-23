package stream

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	fox "github.com/cuhsat/fox/v4/internal"
)

type Schema interface {
	Headers() map[string]string
	String() string
	Write(string)
}

func Post(url string, sc Schema) {
	body := strings.NewReader(sc.String())

	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("user-agent", fmt.Sprintf("%s %s", fox.Product, fox.Version))

	for k, v := range sc.Headers() {
		req.Header.Set(k, v)
	}

	res, err := new(http.Client).Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	if res.StatusCode != 200 {
		log.Println(http.StatusText(res.StatusCode))
	}

	_ = res.Body.Close()
}
