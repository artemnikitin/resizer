package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

var (
	originalFile = flag.String("file", "", "Path to original image")
	saveTo       = flag.String("save", "", "Save to image by path")
	width        = flag.Uint("w", 0, "Width of resized image")
	height       = flag.Uint("h", 0, "Height of resized image")
)

func main() {
	flag.Parse()
	if *originalFile == "" || (*width == 0 && *height == 0) {
		fmt.Println("Please specify correct parameters!")
		os.Exit(1)
	}
	if *saveTo == "" {
		*saveTo = *originalFile
	}

	file, err := os.Open(*originalFile)
	processError(err, "Can't open original file")

	typ := getImageType(*originalFile)
	switch typ {
	case "jpg":
		image := processJPEG(file)
		resizeIt(image)
	case "jpeg":
		image := processJPEG(file)
		resizeIt(image)
	case "png":
		image := processPNG(file)
		resizeIt(image)
	default:
		fmt.Println("Only .jpg and .png are supported at the moment!")
		os.Exit(1)
	}
}

func processJPEG(file *os.File) image.Image {
	image, err := jpeg.Decode(file)
	processError(err, "Can't convert original file to Image")
	file.Close()
	return image
}

func processPNG(file *os.File) image.Image {
	image, err := png.Decode(file)
	processError(err, "Can't convert original file to Image")
	file.Close()
	return image
}

func resizeIt(image image.Image) {
	image = resize.Resize(*width, *height, image, resize.Lanczos3)
	save(*saveTo, image)
	fmt.Println("Resizing done!")
}

func getImageType(path string) string {
	pos := strings.LastIndex(path, ".")
	if pos == -1 {
		log.Fatal("Can't find type of image")
	}
	return path[pos+1:]
}

func save(path string, image image.Image) {
	file, err := os.Create(path)
	processError(err, "Can't create file to save resized image")
	typ := getImageType(path)
	if typ == "png" {
		png.Encode(file, image)
	} else {
		jpeg.Encode(file, image, nil)
	}
}

func processError(err error, text string) {
	if err != nil {
		log.Fatal(text, err)
	}
}
