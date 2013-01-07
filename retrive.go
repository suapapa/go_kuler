package kuler

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Theme struct {
	XMLName       xml.Name `xml:"themeItem"`
	ID            string   `xml:"themeID"`
	Title         string   `xml:"themeTitle"`
	Image         string   `xml:"themeImage"`
	AuthorId      string   `xml:"themeAuthor>authorID"`
	AuthorLabel   string   `xml:"themeAuthor>authorLabel"`
	Tags          string   `xml:"themeTags"`
	Rating        int      `xml:"themeRating"`
	DownloadCount int      `xml:"themeDownloadCount"`
	CreatedAt     string   `xml:"themeCreatedAt"`
	EditedAt      string   `xml:"themeEditedAt"`
	Swatches      []Swatch `xml:"themeSwatches>swatch"`
}

func (t Theme) String() string {
	s := t.Title + "("
	for i, c := range t.Swatches {
		if i != 0 {
			s += ","
		}
		s += fmt.Sprintf("%s", c)
	}
	s += ")"
	return s
}

type themeFeed struct {
	XMLName xml.Name `xml:"rss"`
	// Title          string   `xml:"channel>title"`
	// Link           string   `xml:"channel>link"`
	// Description    string   `xml:"channel>descriptions"`
	// Language       string   `xml:"channel>language"`
	// PubDate        string   `xml:"channel>pubDate"`
	// LastBuildDate  string   `xml:"channel>lastBuildDate"`
	// Docs           string   `xml:"channel>docs"`
	// Generator      string   `xml:"channel>generator"`
	// ManagingEditor string   `xml:"channel>managingEditor"`
	// WebMaster      string   `xml:"channel>webMaster"`
	RecordCount int     `xml:"channel>recordCount"`
	StartIndex  int     `xml:"channel>startIndex"`
	ItemPerPage int     `xml:"channel>itemPerPage"`
	Items       []Theme `xml:"channel>item>themeItem"`
}

func retriveThemes(qUrl string) ([]Theme, error) {
	logD("fetching " + qUrl + "...")
	res, err := http.Get(qUrl)
	if err != nil {
		return nil, errors.New("server error")
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, errors.New("error on read body")
	}

	logD("unmarshlling...")
	var rc themeFeed
	if err := xml.Unmarshal(data, &rc); err != nil {
		logE("%s", err)
		if err := unmarshalServiceError(data); err != nil {
			return nil, err
		} else {
			errStr := "failed to unmarshal service error"
			logE(errStr)
			return nil, errors.New(errStr)
		}
	}

	logD("done")
	return rc.Items, nil
}
