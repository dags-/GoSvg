package svg

import (
	"image"
	"fmt"
	"image/color"
	"io"
)

const headerFormat = "<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 %v %v' shape-rendering='crispEdges'>"
const noAlpha = "<path d='M%v %vH%v' stroke='rgb(%v,%v,%v)'/>"
const withAlpha = "<path d='M%v %vH%v' stroke='rgba(%v,%v,%v,%v)'/>"

func Convert(img image.Image, wr io.Writer) {
	max := img.Bounds().Max
	fmt.Fprintf(wr, headerFormat, max.X, max.Y)
	for y := 0; y < max.Y; y++ {
		yy := 0.5 + float32(y)
		for x := 0; x < max.X; x++ {
			col := color.RGBAModel.Convert(img.At(x, y))
			c := col.(color.RGBA)
			fmt.Fprintf(wr, noAlpha, x, yy, x + 1, c.R, c.G, c.B)
		}
	}
	fmt.Fprint(wr, "</svg>")
}

func ConvertWithAlpha(img image.Image, wr io.Writer)  {
	max := img.Bounds().Max
	fmt.Fprintf(wr, headerFormat, max.X, max.Y)
	for y := 0; y < max.Y; y++ {
		yy := 0.5 + float32(y)
		for x := 0; x < max.X; x++ {
			col := color.RGBAModel.Convert(img.At(x, y))
			c := col.(color.RGBA)
			fmt.Fprintf(wr, withAlpha, x, yy, x + 1, c.R, c.G, c.B, c.A)
		}
	}
	fmt.Fprint(wr, "</svg>")
}