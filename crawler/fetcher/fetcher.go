package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateTimer = time.Tick(100 * time.Microsecond)

func Fetcher(url string) ([]byte, error) {
	<-rateTimer
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
