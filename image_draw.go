package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"

	"os"
)

// https://www.sanarias.com/blog/1214PlayingwithimagesinHTTPresponseingolang
var root = flag.String("root", ".", "file system path")

func main() {
	http.HandleFunc("/draw-image/", drawImageHandler)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func drawImageHandler(w http.ResponseWriter, r *http.Request) {
	image1, err := os.Open("images/master_banner.jpg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	image2, err := os.Open("images/PartsEazyLogo.jpeg")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second, err := jpeg.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	offset := image.Pt(250, 100)
	b := first.Bounds()
	// fmt.Println("Width:", b.Max.X, "Height:", b.Max.Y)
	bounds := second.Bounds()
	fmt.Println("Width:", bounds.Max.X, "Height:", bounds.Max.Y)
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.ZP, draw.Src)
	draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)
	var img image.Image = image3
	writeImage(w, &img)
}

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
