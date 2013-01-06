package kuler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTheme(t *testing.T) {
	d, _ := ioutil.ReadFile("_testdata/theme_example.xml")
	var kt theme
	if err := xml.Unmarshal(d, &kt); err != nil {
		t.Fatalf("%s", err)
	}

	fmt.Println(kt)
}

func TestSwatch(t *testing.T) {
	d, _ := ioutil.ReadFile("_testdata/swatch_example.xml")
	var s swatch
	if err := xml.Unmarshal(d, &s); err != nil {
		t.Fatalf("%s", err)
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
