// http://books.studygolang.com/gopl-zh/ch1/ch1-04.html
package lissajous

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"

	"golang.org/x/exp/slices"
)

var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 2 // next color in palette
)

func Generator(out io.Writer, query map[string][]string) {
	var (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	usefulKeys := []string{"cycles", "size", "nframes", "delay"}

	for k, v := range query {
		if slices.Contains(usefulKeys, k) {

			_v, err := strconv.Atoi(v[0])

			if err != nil {
				fmt.Printf("dup: %v\n", err)
				continue
			}

			size = _v
		}
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
