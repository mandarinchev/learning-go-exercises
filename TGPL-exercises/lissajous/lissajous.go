// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
)

var palette = []color.Color{color.Black,
    color.RGBA{R: 0x10, G: 0xdd, B: 0x10, A: 0xff},
    color.RGBA{R: 0xdd, G: 0x35, B: 0x10, A: 0xff},
    color.RGBA{R: 0x10, G: 0x10, B: 0xe7, A: 0xff},
    color.RGBA{R: 0xa2, G: 0x79, B: 0x30, A: 0xff},
    color.RGBA{R: 0x11, G: 0xbb, B: 0xcc, A: 0xff},
}

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 300   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )
    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    var colorIndex uint8
    for i := 0; i < nframes; i++ {
        colorIndex++
        if colorIndex == uint8(len(palette)) {
            colorIndex = 1
        }
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                colorIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
