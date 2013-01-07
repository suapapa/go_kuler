// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kuler

import (
	"fmt"
	// "net/http"
	// "net/url"
)

/* https://learn.adobe.com/wiki/display/kulerdev/B.%20Feeds */

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

func (s *Service) List(startIdx, itemsPerPage uint) ([]Theme, error) {
	return s.ListRecent(startIdx, itemsPerPage)
}

func (s *Service) ListRecent(startIdx, itemsPerPage uint) ([]Theme, error) {
	opt := listOption{"recent", &pageOption{startIdx, itemsPerPage}}
	return retrive(s.queryUrl(&opt))
}

func (s *Service) ListPopular(startIdx, itemsPerPage uint) ([]Theme, error) {
	opt := listOption{"popular", &pageOption{startIdx, itemsPerPage}}
	return retrive(s.queryUrl(&opt))
}

func (s *Service) ListRating(startIdx, itemsPerPage uint) ([]Theme, error) {
	opt := listOption{"rating", &pageOption{startIdx, itemsPerPage}}
	return retrive(s.queryUrl(&opt))
}

func (s *Service) ListRandom(startIdx, itemsPerPage uint) ([]Theme, error) {
	opt := listOption{"random", &pageOption{startIdx, itemsPerPage}}
	return retrive(s.queryUrl(&opt))
}

func (s *Service) Search(key string, startIdx, itemsPerPage uint) ([]Theme, error) {
	opt := searchOption{key, &pageOption{startIdx, itemsPerPage}}
	return retrive(s.queryUrl(&opt))
}

func (s *Service) SearchThemeID(key string, startIdx, itemsPerPage uint) ([]Theme, error) {
	return s.Search("themeID:"+key, startIdx, itemsPerPage)
}

func (s *Service) SearchUserID(key string, startIdx, itemsPerPage uint) ([]Theme, error) {
	return s.Search("userID:"+key, startIdx, itemsPerPage)
}

func (s *Service) SearchEmail(key string, startIdx, itemsPerPage uint) ([]Theme, error) {
	return s.Search("email:"+key, startIdx, itemsPerPage)
}

func (s *Service) SearchTag(key string, startIdx, itemsPerPage uint) ([]Theme, error) {
	return s.Search("tag:"+key, startIdx, itemsPerPage)
}

func (s *Service) SearchHex(key string, startIdx, itemsPerPage uint) ([]Theme, error) {
	return s.Search("hex:"+key, startIdx, itemsPerPage)
}

func (s *Service) SearchTitle(key string, startIdx, itemsPerPage uint) ([]Theme, error) {
	return s.Search("title:"+key, startIdx, itemsPerPage)
}
