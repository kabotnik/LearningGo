package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{0xff, 0xff, 0xff, 50},
	color.RGBA{0x00, 0x00, 0xff, 50},
	color.RGBA{0xff, 0x00, 0x00, 50},
	color.RGBA{0xff, 0x00, 0xff, 50},
	color.RGBA{0xff, 0xff, 0x00, 50},
	color.RGBA{0x00, 0xff, 0xff, 50},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	imgSrc := "https://media.giphy.com/media/DsOPKOKBszTaM/giphy.gif"
	// imgSrc := "https://media.giphy.com/media/ZHrXOvxBnhCmc/giphy.gif"
	resp, err := http.Get(imgSrc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	img, err := gif.DecodeAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		os.Exit(1)
	}
	// gif.EncodeAll(os.Stdout, img)
	partify(os.Stdout, img)
}

func partify(out io.Writer, giffy *gif.GIF) {
	loops := len(giffy.Image)
	// fmt.Printf("Loops: %d\n", loops)

	anim := gif.GIF{LoopCount: loops}
	for _, frame := range giffy.Image {

		// rect := frame.Bounds()
		// pIndex := i % len(palette)
		// pColor := palette[pIndex]
		// mask := image.NewRGBA(rect)
		// cMask := image.NewUniform(pColor)
		// draw.Draw(frame, frame.Bounds(), cMask, image.ZP, draw.Over)
		draw.Draw(frame, frame.Bounds(), image.Transparent, image.ZP, draw.)

		// g := frame.Bounds
		// gImage := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		// draw.Draw(gImage, gImage.Bounds(), frame, g.Min, draw.Src)

		anim.Image = append(anim.Image, frame)
		anim.Delay = append(anim.Delay, 0)
	}
	// numCycles := float64(cycles) // number of complete x oscillator revolutions
	// const (
	// 	res     = 0.001 // angular resolution
	// 	size    = 100   // image canvas covers [-size..+size]
	// 	nframes = 64    // number of animation frames
	// 	delay   = 8     // delay between frames in 10ms units
	// )
	// freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	// anim := gif.GIF{LoopCount: nframes}
	// phase := 0.0 // phase difference
	// for i := 0; i < nframes; i++ {
	// 	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	// 	img := image.NewPaletted(rect, palette)
	// 	for t := 0.0; t < numCycles*2*math.Pi; t += res {
	// 		x := math.Sin(t)
	// 		y := math.Sin(t*freq + phase)
	// 		paletteColor := (i % (len(palette) - 1)) + 1
	// 		color := uint8(paletteColor)
	// 		img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
	// 			color)
	// 	}
	// 	phase += 0.1
	// 	anim.Delay = append(anim.Delay, delay)
	// 	anim.Image = append(anim.Image, img)
	// }
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
