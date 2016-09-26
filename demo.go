package main

import (
	"os"
	"image/png"
	"github.com/gorilla/GoSvg/svg"
	"net/http"
)

func main() {
	url := "https://40.media.tumblr.com/0c57428045d4cb0d6ada4643c9913bc2/tumblr_inline_ny1oze2U7G1r70e84_540.png"
	resp, err := http.Get(url)
	if err == nil {
		img, _ := png.Decode(resp.Body)
		out, _ := os.Create("meow.svg")
		svg.Convert(img, out)
	}
}
