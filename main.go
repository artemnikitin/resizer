package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"

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
	defer file.Close()

	typ := getImageType(*originalFile)
	switch typ {
	case "jpg":
		image := processJPEG(file)
		resizeIt(typ, &image)
	case "png":
		image := processPNG(file)
		resizeIt(typ, &image)
	default:
		fmt.Println("Only .jpg and .png are supported at the moment!")
		os.Exit(1)
	}
}

func processJPEG(file io.Reader) image.Image {
	image, err := jpeg.Decode(file)
	processError(err, "Can't convert .jpeg file to Image")
	return image
}

func processPNG(file io.Reader) image.Image {
	image, err := png.Decode(file)
	processError(err, "Can't convert .png file to Image")
	return image
}

func resizeIt(typ string, image *image.Image) {
	*image = resize.Resize(*width, *height, *image, resize.Lanczos3)
	save(typ, *saveTo, *image)
	fmt.Println("Resizing done!")
}

func getImageType(path string) string {
	bytes, err := ioutil.ReadFile(path)
	processError(err, "Can't get content of file")
	result := ""
	if len(bytes) > 0 {
		if bytes[0] == 0xFF && bytes[1] == 0xD8 {
			result = "jpg"
		}
		if bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
			result = "png"
		}
	}
	return result
}

func save(typ, path string, image image.Image) {
	file, err := os.Create(path)
	processError(err, "Can't create file to save resized image")
	if typ == "png" {
		png.Encode(file, image)
	} else {
		jpeg.Encode(file, image, nil)
	}
}

func processError(err error, text string) {
	if err != nil {
		log.Fatal(text+" ", err)
	}
}
