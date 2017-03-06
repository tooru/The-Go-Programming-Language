package main

import (
    "image"
    "image/color"
    "image/png"
    "log"
    "os"
)

func main() {
    src, err := png.Decode(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }

	dest := image.NewRGBA(src.Bounds())

    rect := src.Bounds()

    for y := 0; y <= rect.Max.Y; y++ {
        for x := 0; x <= rect.Max.X; x++ {
            dest.Set(x, y, supersampling(src, x, y, rect.Max))
        }
    }
    png.Encode(os.Stdout, dest)
}

func supersampling(img image.Image, x int, y int, max image.Point) color.Color {
    var n uint32 = 0

    var dr, dg, db uint32 = 0, 0, 0

    if x > 0 {
        if y > 0 {
            addRGB(&dr, &dg, &db, img.At(x-1, y-1))
            n++
        }
        addRGB(&dr, &dg, &db, img.At(x-1, y))
        n++
        if y < max.Y {
            addRGB(&dr, &dg, &db, img.At(x-1, y+1))
            n++
        }
    }        
    if y > 0 {
        addRGB(&dr, &dg, &db, img.At(x, y-1))
        n++
    }
    c := img.At(x, y)
    _, _, _, a := c.RGBA()
    addRGB(&dr, &dg, &db, c)
    n++
    if y < max.Y {
        addRGB(&dr, &dg, &db, img.At(x, y+1))
        n++
    }
    if x < max.X {
        if y > 0 {
            addRGB(&dr, &dg, &db, img.At(x+1, y-1))
            n++
        }
        addRGB(&dr, &dg, &db, img.At(x+1, y))
        n++
        if y < max.Y {
            addRGB(&dr, &dg, &db, img.At(x+1, y+1))
            n++
        }
    }        

    log.Printf("%d %d,%d,%d", n, dr/n, dg/n, db/n)

    return color.RGBA{uint8(dr / n), uint8(dg / n), uint8(db / n), uint8(a)}
}

func addRGB(dr, dg, db *uint32, c color.Color) {
    r, g, b, _ := c.RGBA()

    *dr += r >> 8
    *dg += g >> 8
    *db += b >> 8
}
