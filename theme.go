package kuler

import (
	"encoding/xml"
	"fmt"
)

type theme struct {
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
	Swatches      []swatch `xml:"themeSwatches>swatch"`
}

func (t theme) String() string {
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
