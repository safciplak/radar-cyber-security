package numbersapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func New() *Client {
	return &Client{}
}

type Client struct{}

func (c *Client) makeRequest(req *http.Request) (string, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("numbersapi: unable to read response body: %w", err)
	}

	return string(body), err
}

func (c *Client) GetRandomText(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://numbersapi.com/random/math", nil)

	if err != nil {
		return "", err
	}

	res, err := c.makeRequest(req)
	if err != nil {
		return "", err
	}

	return res, nil
}
