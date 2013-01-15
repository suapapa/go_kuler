// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kuler

import (
	"encoding/xml"
)

// Swatch is a color of Kuler theme
type Swatch struct {
	XMLName   xml.Name `xml:"swatch"`
	HexColor  string   `xml:"swatchHexColor"`
	ColorMode string   `xml:"swatchColorMode"`
	Channel1  float32  `xml:"swatchChannel1"`
	Channel2  float32  `xml:"swatchChannel2"`
	Channel3  float32  `xml:"swatchChannel3"`
	Channel4  float32  `xml:"swatchChannel4"`
	Index     int      `xml:"swatchIndex"`
}

func (s Swatch) String() string {
	return "#" + s.HexColor
}

func (s Swatch) RGBA() (r, g, b, a uint32) {
	switch s.ColorMode {
	case "rgb":
		r = toU16(s.Channel1)
		g = toU16(s.Channel2)
		b = toU16(s.Channel3)
		a = 0xffff
	case "cmyk":
		c, m, y, k := s.Channel1, s.Channel2, s.Channel3, s.Channel4
		r = toU16((1 - c) * (1 - k))
		g = toU16((1 - m) * (1 - k))
		b = toU16((1 - y) * (1 - k))
		a = 0xffff
	default:
		logE("Unknown color format, %s!", s.ColorMode)
	}
	return
}

func toU16(c float32) uint32 {
	r := uint32(0xffff*c + 0.5)
	return r
}
