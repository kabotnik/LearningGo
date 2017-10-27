package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.White,
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
}

// const (
// 	//whiteIndex = 0 // first color in palette
// 	//blackIndex = 1 // next color in palette

// 	blackIndex = 0
// 	greenIndex = 1
// 	whiteIndex = 2
// )

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		// hander := func(w http.ResponseWriter, r *http.Request) {
		// 	lissajous(w)
		// }
		// http.HandleFunc("/", handler)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			qv := r.URL.Query()
			c := qv.Get("cycles")
			cycles, err := strconv.Atoi(c)
			if err != nil {
				cycles = 5
			}
			lissajous(w, cycles)
		})
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, 5)
}

func lissajous(out io.Writer, cycles int) {
	numCycles := float64(cycles) // number of complete x oscillator revolutions
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < numCycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			paletteColor := (i % (len(palette) - 1)) + 1
			color := uint8(paletteColor)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				color)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
