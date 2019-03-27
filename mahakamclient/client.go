package mahakamclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mahakamcloud/netd/config"
)

const (
	MahakamHostBaseUrl         = "http://%s:%d/v1"
	MahakamHostRegistrationAPI = "/bare-metal-hosts"

	contentTypeJSON = "application/json"
)

type Client struct{}

func (s *Client) RegisterHost(reader *bytes.Buffer) ([]byte, error) {
	client := &http.Client{}

	url := fmt.Sprintf(MahakamHostBaseUrl+MahakamHostRegistrationAPI, config.MahakamIP(), config.MahakamPort())

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentTypeJSON)

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
