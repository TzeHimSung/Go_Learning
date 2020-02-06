// package main

// import "fmt"

// func main() {
// 	var x uint8 = 1<<1 | 1<<5
// 	var y uint8 = 1<<1 | 1<<2
// 	fmt.Printf("%08b\n", x)
// 	fmt.Printf("%08b\n", y)
// 	fmt.Printf("%08b\n", x&y)
// 	fmt.Printf("%08b\n", x|y)
// 	fmt.Printf("%08b\n", x^y)
// 	fmt.Printf("%08b\n", x&^y)
// 	for i := uint(0); i < 8; i++ {
// 		if x&(1<<i) != 0 {
// 			fmt.Println(i)
// 		}
// 	}
// 	fmt.Printf("%08b\n", x<<1)
// 	fmt.Printf("%08b\n", x>>1)
// }

// package main

// import "fmt"

// func main() {
// 	var z float64
// 	fmt.Println(z, -z, 1/z, -1/z, z/z)
// }

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// package main

// import "fmt"

// func main() {
// 	var x, y complex128 = complex(1, 2), complex(3, 4)
// 	fmt.Println(x * y)
// 	fmt.Println(real(x * y))
// 	fmt.Println(imag(x * y))
// }

// package main

// import (
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"math/cmplx"
// 	"os"
// )

// func main() {
// 	const (
// 		xmin, ymin, xmax, ymax = -2, -2, +2, +2
// 		width, height          = 1024, 1024
// 	)
// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	for py := 0; py < height; py++ {
// 		y := float64(py)/height*(ymax-ymin) + ymin
// 		for px := 0; px < width; px++ {
// 			x := float64(px)/width*(xmax-xmin) + xmin
// 			z := complex(x, y)
// 			img.Set(px, py, mandelbrot(z))
// 		}
// 	}
// 	png.Encode(os.Stdout, img)
// }

// func mandelbrot(z complex128) color.Color {
// 	const iterations = 200
// 	const contrast = 15

// 	var v complex128
// 	for n := uint8(0); n < iterations; n++ {
// 		v = v*v + z
// 		if cmplx.Abs(v) > 2 {
// 			return color.Gray{255 - contrast*n}
// 		}
// 	}
// 	return color.Black
// }
