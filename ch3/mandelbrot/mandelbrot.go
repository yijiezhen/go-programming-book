package main

import (
	"image"
	"image/png"
	"image/color"
	"math/cmplx"
	"net/http"
	"log"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height = 1024, 1024
)

var wg sync.WaitGroup

func handler(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		wg.Add(1)
		go func(py int) {
			defer wg.Done()
			y := float64(py) / height * (ymax - ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px) / width * (xmax - xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
		}(py)
	}
	wg.Wait()
	png.Encode(w, img)
	w.Header().Set("ContentType", "image/png")
	end := time.Now()
	log.Printf("total: %s", end.Sub(begin))
}

func mandelbrot(z complex128) color.Color {
	//begin := time.Now()
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast * n}
		}
	}
	//log.Printf("mandelbrot: %s", time.Now().Sub(begin))
	return color.Black
}