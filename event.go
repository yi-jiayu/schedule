package main

import (
	"encoding/xml"
)

type Feed struct {
	Entry Entry `xml:"entry"`
}

type Entry struct {
	VideoID      string `xml:"videoId"`
	ChannelID    string `xml:"channelId"`
	Title        string `xml:"title"`
	ChannelTitle string `xml:"author>name"`
}

func parseFeed(data string) (feed Feed, err error) {
	err = xml.Unmarshal([]byte(data), &feed)
	return
}
