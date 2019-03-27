package client

import (
	"io/ioutil"
	"net/http"
)

type Client struct{}

func (s *Client) DoRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
