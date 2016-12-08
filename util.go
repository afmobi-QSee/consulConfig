package consulConfig

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet(url string) (string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return ""
	}

	return string(body)
}