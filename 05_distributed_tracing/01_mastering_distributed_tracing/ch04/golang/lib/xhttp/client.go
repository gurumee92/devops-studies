package xhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get is...
func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return Do(req)
}

// Do is...
func Do(req *http.Request) ([]byte, error) {
	return DoWithClient(req, http.DefaultClient)
}

// DoWithClient is..
func DoWithClient(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}
	return body, nil
}
