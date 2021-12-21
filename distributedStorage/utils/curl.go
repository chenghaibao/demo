package utils

import (
	"io/ioutil"
	"net/http"
)

func GetCurl(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body)
}
