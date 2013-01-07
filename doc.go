// Package kuler is provides access to the Adobe Kuler API (kuler.adobe.com).
// 
// See https://learn.adobe.com/wiki/display/kulerdev/B.%20Feeds
// 
// Usage example:
// 
//     import github.com/suapapa/go_kuler
//     ...
//     kulerService := kuler.NewService("YOUR_APIKEY")
//     themes, _ := kulerService.Search("MLP", 0, 20)
//     for _, t := range themes {
//         fmt.Println(t)
//     }
package kuler
