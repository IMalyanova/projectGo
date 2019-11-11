package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func main() {
	
}

const (
	BotToken = "310805560:AAENzjDSJPKABY9Hw1GZOdKBxxrhOHkfo_k"
	WebhookURL = "http://ea731f5c.ngrok.io"
)

var rss = map[string]string {
	"Habr": "https://habrahabr.ru/rss/best/",
}

type RSS struct {
	Items []Item `xml:"channel>item"`
}

type Item struct {
	URL   string `xml:"guid"`
	Title string `xml:"title"`
}

func getNews(url string) (*RSS, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	rss := new(RSS)
	err = xml.Unmarshal(body, rss)

	if err != nil {
		return nil, err
	}

	return rss, nil
}