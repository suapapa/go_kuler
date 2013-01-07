// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kuler

import (
	"fmt"
)

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

func (s *Service) List(start, max uint) ([]Theme, error) {
	return s.ListRecent(start, max)
}

func (s *Service) ListRecent(start, max uint) ([]Theme, error) {
	opt := listOption{"recent", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

func (s *Service) ListPopular(start, max uint) ([]Theme, error) {
	opt := listOption{"popular", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

func (s *Service) ListRating(start, max uint) ([]Theme, error) {
	opt := listOption{"rating", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

func (s *Service) ListRandom(start, max uint) ([]Theme, error) {
	opt := listOption{"random", &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

func (s *Service) Search(key string, start, max uint) ([]Theme, error) {
	opt := searchOption{key, &pageOption{start, max}}
	return retriveThemes(s.queryUrl(&opt))
}

func (s *Service) SearchThemeID(key string, start, max uint) ([]Theme, error) {
	return s.Search("themeID:"+key, start, max)
}

func (s *Service) SearchUserID(key string, start, max uint) ([]Theme, error) {
	return s.Search("userID:"+key, start, max)
}

func (s *Service) SearchEmail(key string, start, max uint) ([]Theme, error) {
	return s.Search("email:"+key, start, max)
}

func (s *Service) SearchTag(key string, start, max uint) ([]Theme, error) {
	return s.Search("tag:"+key, start, max)
}

func (s *Service) SearchHex(key string, start, max uint) ([]Theme, error) {
	return s.Search("hex:"+key, start, max)
}

func (s *Service) SearchTitle(key string, start, max uint) ([]Theme, error) {
	return s.Search("title:"+key, start, max)
}
