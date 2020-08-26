package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/color"
	"image/jpeg"
	"net/http"
	"os"
	"math"
	"io/ioutil"
	"strconv"
	"time"
)

// find the average color of the picture.
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

// resize an image by its ratio e.g. ratio 2 means reduce the size by 1/2, 10 means reduce the size by 1/10
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	width := bounds.Max.X - bounds.Min.Y
	ratio := width / newWidth
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}
	return *out
}

// populate a tiles database in memory
func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, "when populating tiles db:", err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}

// find the nearest matching image.
func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1000000.0
	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(*db, filename)
	return filename
}

// find the Eucleadian distance between 2 points.
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

// find the square
func sq(n float64) float64 {
	return n * n
}

var TILESDB map[string][3]float64

func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	// building up the source file database
	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(writer, nil)
}

func mosaic(writer http.ResponseWriter, request *http.Request) {
	t0 := time.Now()
	// get the content from the POSTed form
	request.ParseMultipartForm(10485760) // max body in memory is 10MB
	file, _, _ := request.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(request.FormValue("tile_size"))
	// decode and get original image
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	// create a new image for the mosaic
	newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))
	// build up the tiles database
	db := cloneTilesDB()
	// source point for each tile, which starts with 0, 0 of each file
	sp := image.Point{0, 0}
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {
			// use the top left most pixel as the average color
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}
			// get the closest tile from the tiles DB
			nearest := nearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					// resize the tile to the correct size
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					// draw the tile into the mosaic
					draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("error:", err, nearest)
				}
			} else {
				fmt.Println("error:", nearest)
			}
			file.Close()
		}
	}

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newimage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())

	t1 := time.Now()
	images := map[string]string {
		"original": originalStr,
		"mosaic": mosaic,
		"duration": fmt.Sprintf("%v ", t1.Sub(t0)),
	}
	t, _ := template.ParseFiles("results.html")
	t.Execute(writer, images)
}