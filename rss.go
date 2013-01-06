package kuler

import (
	"encoding/xml"
)

type rssChannel struct {
	XMLName        xml.Name `xml:"rss"`
	Title          string   `xml:"channel>title"`
	Link           string   `xml:"channel>link"`
	Description    string   `xml:"channel>descriptions"`
	Language       string   `xml:"channel>language"`
	PubDate        string   `xml:"channel>pubDate"`
	LastBuildDate  string   `xml:"channel>lastBuildDate"`
	Docs           string   `xml:"channel>docs"`
	Generator      string   `xml:"channel>generator"`
	ManagingEditor string   `xml:"channel>managingEditor"`
	WebMaster      string   `xml:"channel>webMaster"`
	RecordCount    int      `xml:"channel>recordCount"`
	StartIndex     int      `xml:"channel>startIndex"`
	ItemPerPage    int      `xml:"channel>itemPerPage"`
	Items          []theme  `xml:"channel>item>themeItem"`
}
