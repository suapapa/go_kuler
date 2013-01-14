# go_kuler : It makes kuler 20% cooler!

Go package for Adobe's color theme service, [Kuler][1].
 
This is an unofficial wrapper for [Kuler API][2].
Docs are avaiable at [godoc.org][3].

## Install

    $ go get github.com/suapapa/go_kuler

## Example

Search pony themes:

    import github.com/suapapa/go_kuler
    ...
    ks := kuler.NewService("YOUR_APIKEY")
    themes, _ := ks.Search("MLP", 0, 20)
    for _, t := range themes {
        fmt.Println(t)
    }



You need an [Kuler API Key][4] to use the service.
A theme is consist of 5 swatches. Swatch is [color.Color][5]. 


[1]:http://kuler.adobe.com
[2]:https://learn.adobe.com/wiki/display/kulerdev/A.+Kuler+API+Documentation
[3]:http://godoc.org/github.com/suapapa/go_kuler
[4]:https://kuler.adobe.com/api/
[5]:http://godoc.org/image/color#Color
