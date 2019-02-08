package main

import (
	"github.com/imroc/req"
)

type Searcher interface {
}

type ImageSearcher struct {
	Config SearchConfig
}

func NewImageSearcher() *ImageSearcher {
	return &ImageSearcher{}
}

func (imageSearcher *ImageSearcher) setConfig(config SearchConfig) {
	imageSearcher.Config = config
}

func (imageSearcher *ImageSearcher) GetImage(query string) {
	SearchImage()
}

func (imageSearcher *ImageSearcher) SearchImage(query string) (string, error) {
	header := req.Header{
		"X-RapidAPI-key": imageSearcher.Config.Key(),
	}
	param := req.Param{
		"autoCorrect": "false",
		"pageNumber":  "1",
		"pageSize":    "1",
		"safeSearch":  "false",
		"q":           query,
	}

	res, err := req.New().Get(imageSearcher.Config.Url(), header, param)
	if err != nil {
		return "", err
	}
	var x ImageSearchResponse

	res.ToJSON(&x)

	if len(x.Value) < 1 {
		return "", NoSearchResultError
	}

	return x.Value[0].URL, nil
}

func DownloadImage(url string) ([]byte, error) {
	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}

type SearchConfig interface {
	Url() string
	Key() string
}

type ImageSearchConfig struct {
	url string
	key string
}

func NewImageSearchConfig() *ImageSearchConfig {
	return &ImageSearchConfig{}
}

func (imageSearchConfig ImageSearchConfig) Url() string {
	return imageSearchConfig.url
}

func (imageSearchConfig ImageSearchConfig) Key() string {
	return imageSearchConfig.key
}

func (imageSearchConfig ImageSearchConfig) SetKey(newKey string) ImageSearchConfig {
	imageSearchConfig.key = newKey
	return imageSearchConfig
}

func (imageSearchConfig ImageSearchConfig) SetUrl(newUrl string) ImageSearchConfig {
	imageSearchConfig.url = newUrl
	return imageSearchConfig
}

type ImageSearchResponse struct {
	Type       string `json:"_type"`
	TotalCount int    `json:"totalCount"`
	Value      []struct {
		URL             string      `json:"url"`
		Height          int         `json:"height"`
		Width           int         `json:"width"`
		Thumbnail       string      `json:"thumbnail"`
		ThumbnailHeight int         `json:"thumbnailHeight"`
		ThumbnailWidth  int         `json:"thumbnailWidth"`
		Base64Encoding  interface{} `json:"base64Encoding"`
	} `json:"value"`
}
