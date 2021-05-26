package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func randUint8() uint8 {
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(256))
}

var RED uint8 = randUint8()
var BLUE uint8 = randUint8()
var GREEN uint8 = randUint8()

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		xP := r.URL.Query()["x"]
		yP := r.URL.Query()["y"]
		zoomP := r.URL.Query()["zoom"]

		x := 2
		y := 2
		zoom := 1024

		if len(xP) > 0 {
			x, _ = strconv.Atoi(xP[0])
		}

		if len(yP) > 0 {
			y, _ = strconv.Atoi(yP[0])
		}

		if len(zoomP) > 0 {
			zoom, _ = strconv.Atoi(zoomP[0])
		}

		width := zoom
		height := zoom

		xmin := -x
		xmax := x
		ymin := -y
		ymax := y

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/float64(width)*float64(ymax-ymin) + float64(ymin)
			for px := 0; px < width; px++ {
				x := float64(px)/float64(width)*float64(xmax-xmin) + float64(xmin)
				z := complex(x, y)
				//Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		png.Encode(w, img)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{RED, GREEN, BLUE, 255 - contrast*n}
		}
	}
	return color.Black
}
