// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package kuler provides access to the Adobe Kuler service (kuler.adobe.com).
// 
// See https://learn.adobe.com/wiki/display/kulerdev/B.%20Feeds
// 
// Usage example:
// 
//     import github.com/suapapa/go_kuler
//     ...
//     ks := kuler.NewService("YOUR_APIKEY")
//     themes, _ := ks.Search("MLP", 0, 20)
//     for _, t := range themes {
//         fmt.Println(t)
//     }
package kuler
