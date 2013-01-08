// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kuler

import (
	"fmt"
	"net/url"
)

type pageOption struct {
	startIdx     uint
	itemsPerPage uint // 1~100 (default 20)
}

func (q *pageOption) String() string {
	if 1 > q.itemsPerPage || q.itemsPerPage > 100 {
		logW("invaild itemsPerPage, %d. Fix it to 20", q.itemsPerPage)
		q.itemsPerPage = 20
	}
	return fmt.Sprintf("startIndex=%d&itemsPerPage=%d",
		q.startIdx, q.itemsPerPage)
}

type listOption struct {
	listType string // recent, popular, rating, random
	page     *pageOption
}

func (q *listOption) String() string {
	switch q.listType {
	case "recent", "popular", "rating", "random":
		// valid options
	default:
		logW("invalid listType, %s. Fix it to recent", q.listType)
		q.listType = "recent"
	}

	return fmt.Sprintf("rss/get.cfm?listType=%s&%s",
		q.listType, q.page)
}

type searchOption struct {
	searchOption string
	page         *pageOption
}

func (q *searchOption) String() string {
	return fmt.Sprintf("rss/search.cfm?searchQuery=%s&%s",
		url.QueryEscape(q.searchOption), q.page)
}
