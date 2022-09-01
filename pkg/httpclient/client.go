package httpclient

import (
	"log"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

type Config struct {
	BasePath string
}

const API_KEY = "b15242a4a7f3eb9c34df3c17da4b89f11e4143c5"

const basePath = "https://suggestions.dadata.ru"

func New(cfg *Config) *Client {
	if cfg.BasePath == "" {
		cfg.BasePath = basePath
	}
	return &Client{
		client: resty.New().SetBaseURL(cfg.BasePath),
	}
}

type Request struct {
	Query string `json:"query"`
}
type Response struct {
	Suggestions []Suggestions `json:"suggestions"`
}
type Suggestions struct {
	Value string `json:"value"`
	Data  Data   `json:"data"`
}
type Data struct {
	KPP        string     `json:"kpp"`
	Management Management `json:"management"`
}
type Management struct {
	Name string `json:"name"`
}

func (c *Client) GetInfo(INN string) (*Response, error) {
	path := "/suggestions/api/4_1/rs/findById/party"

	result := &Response{}
	request := Request{}

	request.Query = INN

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Token "+API_KEY).
		SetBody(request).
		SetResult(&result).
		Post(path)
	if err != nil {
		return nil, err
	}
	log.Println(result, resp.Status())
	return result, nil
}
