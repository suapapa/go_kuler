package kuler

import (
	"encoding/xml"
)

type swatch struct {
	XMLName   xml.Name `xml:"swatch"`
	HexColor  string   `xml:"swatchHexColor"`
	ColorMode string   `xml:"swatchColorMode"`
	Channel1  float32  `xml:"swatchChannel1"`
	Channel2  float32  `xml:"swatchChannel2"`
	Channel3  float32  `xml:"swatchChannel3"`
	Channel4  float32  `xml:"swatchChannel4"`
	Index     int      `xml:"swatchIndex"`
}

func (s swatch) String() string {
	return "#" + s.HexColor
}

func (s swatch) RGBA() (r, g, b, a uint32) {
	r = s.toU16(s.Channel1)
	g = s.toU16(s.Channel2)
	b = s.toU16(s.Channel3)
	a = 0xffff
	return
}

func (s swatch) toU16(c float32) uint32 {
	r := uint32(0xffff*c + 0.5)
	return r
}
