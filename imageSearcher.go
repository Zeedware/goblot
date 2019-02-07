package main

import (
	"net/http"
	"runtime/debug"
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

func (imageSearcher *ImageSearcher) SearchImage() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", imageSearcher.Config.GetUrl(), nil)
	if err != nil {
		debug.PrintStack()
	}
	req.Header.Set("X-RapidAPI-Key", imageSearcher.Config.GetKey())
	client.Do(req)

	return ""
}

type SearchConfig interface {
	GetUrl() string
	GetKey() string
}

type ImageSearchConfig struct {
	Url string
	Key string
}

func (imageSearchConfig ImageSearchConfig) GetUrl() string {
	return imageSearchConfig.Url
}

func (imageSearchConfig ImageSearchConfig) GetKey() string {
	return imageSearchConfig.Key
}

func (imageSearchConfig ImageSearchConfig) SetKey(newKey string) ImageSearchConfig {
	imageSearchConfig.Key = newKey
	return imageSearchConfig
}

func (imageSearchConfig ImageSearchConfig) SetUrl(newUrl string) ImageSearchConfig {
	imageSearchConfig.Url = newUrl
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
