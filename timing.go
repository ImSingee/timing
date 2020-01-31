package timing

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const baseURL = "https://web.timingapp.com/"

var authorization string

func Init(apiKey string) {
	if apiKey != "" {
		authorization = "Bearer " + apiKey
	}
}

func IsInitialed() bool {
	return authorization != ""
}

func request(method string, url string, body string) (*response, error) {
	if !IsInitialed() {
		return nil, &InvalidRequest{"未初始化"}
	}

	log.Printf("%s %s", method, url)

	var req *http.Request
	var err error

	client := http.Client{}

	if len(body) != 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, strings.NewReader(body))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", authorization)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &InvalidResponse{"resp.StatusCode != 200"}
	}

	res := response{
		StatusCode: resp.StatusCode,
	}

	res.Body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
