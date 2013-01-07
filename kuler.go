// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kuler

import (
	"fmt"
)

// Returns a instance of kuler.Service.
// The key is Mandatory. Get one from https://kuler.adobe.com/api
// if you don't have already.
//
// All of it's List*() and Search*() functions have two argument,
// start and max. It because the Kuler limits a list up to 100.
// For example, to get list of 120~130th popular themes,
// You should try like this;
//
//    ks.ListPopular(120, 10)
//
// The max argument should be in the range 1..100. Otherwise it forced to
// be default value, 20.
func NewService(key string) *Service {
	s := new(Service)
	s.apiKey = key
	return s
}

type Service struct {
	apiKey string
}

func (s *Service) queryUrl(option interface{}) string {
	return fmt.Sprintf("http://kuler-api.adobe.com/%s&key=%s",
		option, s.apiKey)
}

// Same as ListRecent
func (s *Service) List(start, max uint) ([]Theme, error) {
	return s.ListRecent(start, max)
}

// Get most recent themes
func (s *Service) ListRecent(start, max uint) ([]Theme, error) {
	opt := listOption{"recent", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

// Get most popular themes
func (s *Service) ListPopular(start, max uint) ([]Theme, error) {
	opt := listOption{"popular", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

// Get highest-rated themes
func (s *Service) ListRating(start, max uint) ([]Theme, error) {
	opt := listOption{"rating", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

// Get random themes
func (s *Service) ListRandom(start, max uint) ([]Theme, error) {
	opt := listOption{"random", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

// Simple search. It looks for titles, tags, author names, themeIDs, and hexValues.
func (s *Service) Search(key string, start, max uint) ([]Theme, error) {
	opt := searchOption{key, &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

// Search on a specific themeIDs.
func (s *Service) SearchThemeID(key string, start, max uint) ([]Theme, error) {
	return s.Search("themeID:"+key, start, max)
}

// Search on specific userID
func (s *Service) SearchUserID(key string, start, max uint) ([]Theme, error) {
	return s.Search("userID:"+key, start, max)
}

// Search on specific email
func (s *Service) SearchEmail(key string, start, max uint) ([]Theme, error) {
	return s.Search("email:"+key, start, max)
}

// Search on a tag word
func (s *Service) SearchTag(key string, start, max uint) ([]Theme, error) {
	return s.Search("tag:"+key, start, max)
}

// Search on a hex color value (can be in the format of "ABCDEF" of "0xABCDEF"
func (s *Service) SearchHex(key string, start, max uint) ([]Theme, error) {
	return s.Search("hex:"+key, start, max)
}

// Search on a theme title
func (s *Service) SearchTitle(key string, start, max uint) ([]Theme, error) {
	return s.Search("title:"+key, start, max)
}
