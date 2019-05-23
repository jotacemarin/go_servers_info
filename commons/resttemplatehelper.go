package commons

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPGet : func
func HTTPGet(url string) ([]byte, error) {
	log.Printf("request to %s", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data, nil
}
