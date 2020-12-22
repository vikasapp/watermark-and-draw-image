package main

import (
    "fmt"
    "image"
    "image/draw"
    "image/jpeg"
    // "image/png"
    "os"
    "log"
)

func main() {
    image1,err := os.Open("images/master_banner.jpg")
    if err != nil {
        log.Fatalf("failed to open: %s", err)
    }
    
    first, err := jpeg.Decode(image1)
    if err != nil {
        log.Fatalf("failed to decode: %s", err)
    }
    defer image1.Close()

    image2,err := os.Open("images/PartsEazyLogo.jpeg")
    if err != nil {
        log.Fatalf("failed to open: %s", err)
    }
    second,err := jpeg.Decode(image2)
    if err != nil {
        log.Fatalf("failed to decode: %s", err)
    }
    defer image2.Close()

    offset := image.Pt(300, 200)
    b := first.Bounds()
    image3 := image.NewRGBA(b)
    draw.Draw(image3, b, first, image.ZP, draw.Src)
    draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)

    third,err := os.Create("result/result.jpg")
    if err != nil {
        log.Fatalf("failed to create: %s", err)
    }
    jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
    defer third.Close()
}