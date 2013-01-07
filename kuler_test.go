// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kuler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const TEST_API_KEY = "your_apikey_here"

func testLogout() {
	LogTo(os.Stdout)
	/* LogTo(ioutil.Discard) */
}

func testApiKey() string {
	k, _ := ioutil.ReadFile(".apikey")
	return strings.TrimSpace(string(k))
}

func TestSearch(t *testing.T) {
	testLogout()
	s := NewService(testApiKey())
	if ts, err := s.Search("MLP", 0, 0); err == nil {
		for _, t := range ts {
			logD("%s", t)
		}
	} else {
		logE("%s", err)
	}
	logD("---- end of Search MLP ----")
}

func TestListRandom(t *testing.T) {
	testLogout()
	s := NewService(testApiKey())
	if ts, err := s.ListRandom(0, 0); err == nil {
		for _, t := range ts {
			logD("%s", t)
		}
	} else {
		logE("%s", err)
	}
	logD("---- end of ListRandom ----")
}

func TestFixInvaildOptions(t *testing.T) {
	testLogout()
	s := NewService("dummy_key")
	opt := listOption{"cooler", &pageOption{20, 0}}
	expect := "http://kuler-api.adobe.com/rss/get.cfm?" +
		"listType=recent&startIndex=20&itemsPerPage=20&key=dummy_key"
	if got := s.queryUrl(&opt); got != expect {
		t.Fatalf("expect %s, got %s", expect, got)
	}
}

func TestServiceQueryUrl(t *testing.T) {
	testLogout()
	s := NewService("dummy_key")
	opt := listOption{"popular", &pageOption{20, 40}}

	expect := "http://kuler-api.adobe.com/rss/get.cfm?" +
		"listType=popular&startIndex=20&itemsPerPage=40&key=dummy_key"
	if got := s.queryUrl(&opt); got != expect {
		t.Fatalf("expect %s, got %s", expect, got)
	}
}

func TestRssChannel(t *testing.T) {
	testLogout()
	d, _ := ioutil.ReadFile("_testdata/highest_rated.xml")
	var rc themeFeed
	if err := xml.Unmarshal(d, &rc); err != nil {
		t.Fatalf("%s", err)
	}

	for _, c := range rc.Items {
		logD("%s", c)
	}
}

func TestTheme(t *testing.T) {
	testLogout()
	d, _ := ioutil.ReadFile("_testdata/theme_example.xml")
	var kt Theme
	if err := xml.Unmarshal(d, &kt); err != nil {
		t.Fatalf("%s", err)
	}

	logD("%s", kt)
}

func TestSwatch(t *testing.T) {
	testLogout()
	d, _ := ioutil.ReadFile("_testdata/swatch_example.xml")
	var s Swatch
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
	testLogout()
	c, _ := ioutil.ReadFile("_testdata/error_0.xml")
	if err := unmarshalServiceError(c); err == nil {
		t.Fatalf("failed to unmarshal service error!")
	}
}
