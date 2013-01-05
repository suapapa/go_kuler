package kuler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSwatch(t *testing.T) {
	d, _ := ioutil.ReadFile("_testdata/swatch_example.xml")
	var s swatch
	if err := xml.Unmarshal(d, &s); err != nil {
		fmt.Println(err)
	}

	r, g, b, _ := s.RGBA()
	r = (r * 0xFF) / 0xFFFF
	g = (g * 0xFF) / 0xFFFF
	b = (b * 0xFF) / 0xFFFF
	rgbStr := fmt.Sprintf("#%02X%02X%02X", r, g, b)

	if s.String() != rgbStr {
		t.Fatalf("expect %s, got %s", rgbStr, s)
	}
}

func TestServiceError(t *testing.T) {
	c, _ := ioutil.ReadFile("_testdata/error_0.xml")
	if err := unmarshalServiceError(c); err == nil {
		t.Fatalf("failed to unmarshal service error!")
	}
}
