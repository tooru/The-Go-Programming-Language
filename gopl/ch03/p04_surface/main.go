// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	defaultWidth  = 600 // canvas size in pixels
	defaultHeight = 320
	cells         = 100         // number of grid cells
	xyrange       = 30.0        // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type cornerParams struct {
	width   float64
	height  float64
	xyscale float64
	zscale  float64
}

//!+main

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		width, ok := getFloatParam(r, "width", defaultWidth)

		if !ok {
			return
		}

		height, ok := getFloatParam(r, "height", defaultHeight)
		if !ok {
			return
		}

		w.Header().Set("Content-Type", "image/svg+xml")

		surface(w, width, height)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getFloatParam(r *http.Request, key string, defaultValue float64) (float64, bool) {
	v, ok := r.Form[key]

	if !ok {
		log.Printf("no key: %s", key)
		return defaultValue, true
	}

	value, err := strconv.ParseFloat(v[0], 64)
	if err != nil {
		log.Fatalf("strconv.Atoi failed: %v", err)
		return 0, false
	}
	return value, true
}

func surface(w io.Writer, width float64, height float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	params := new(cornerParams)

	params.width = width
	params.height = height
	params.xyscale = width / 2 / xyrange  // pixels per x or y unit
	params.zscale = float64(height) * 0.4 // pixels per z unit

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, params)
			bx, by := corner(i, j, params)
			cx, cy := corner(i, j+1, params)
			dx, dy := corner(i+1, j+1, params)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i int, j int, params *cornerParams) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := params.width/2 + (x-y)*cos30*params.xyscale
	sy := params.height/2 + (x+y)*sin30*params.xyscale - z*params.zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-main
