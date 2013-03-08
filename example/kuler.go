// +build !appengine

// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/suapapa/go_kuler"
	"os"
)

func main() {
	ks := kuler.NewService(os.Getenv("KULER_APIKEY"))
	themes, _ := ks.Search(os.Args[1], 0, 20)
	for _, t := range themes {
		fmt.Println(t)
	}
}
