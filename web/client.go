package web

import (
	"log"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			//log.Printf(err)
			return
		}
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {

}
