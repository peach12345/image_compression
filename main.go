package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main() {

	//Get image max size from user input
	w, h := getWidthAndHeightFromUserInput()
	//Call read image method

	imageToShrink, imageConfig := readImageFromFile("/Users/danielweyck/Documents/image_compression/asset/test2.jpg")

	fmt.Println("------")
	fmt.Println("The actual image size: ", imageConfig.Width, "x", imageConfig.Height)
	fmt.Println("Will be converted to:", w, "x", h)

	imageToShrink = shrinkImage(imageToShrink, w, h)

	writeImage(imageToShrink)
	fmt.Println("Successfully resized all images")
}

func getWidthAndHeightFromUserInput() (int, int) {
	fmt.Println("Select max width")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	width, _ := reader.ReadString('\n')

	width = strings.Replace(width, "\n", "", -1)

	fmt.Println("Select max height")
	fmt.Print("-> ")

	height, _ := reader.ReadString('\n')
	height = strings.Replace(height, "\n", "", -1)

	w, err := strconv.Atoi(width)
	if err != nil {
		log.Fatal(err)
	}

	h, err := strconv.Atoi(height)
	if err != nil {
		log.Fatal(err)
	}
	return w, h
}

// read image from file
func readImageFromFile(path string) (image.Image, image.Config) {

	f, err := os.Open(path)

	if err != nil {
		log.Fatal("Error while reading image", err)
	}
	imageConfig, _, err := image.DecodeConfig(f)

	if err != nil {
		log.Fatal("Error while reading config", err)
	}

	d, err := os.Open(path)

	if err != nil {
		log.Fatal("Error while reading image", err)
	}

	image, _, err := image.Decode(d)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer f.Close()
	defer d.Close()

	return image, imageConfig
}

//Shrink image method

func shrinkImage(imageToShrink image.Image, width int, height int) image.Image {

	image := resize.Resize(uint(width), uint(height), imageToShrink, resize.MitchellNetravali)

	return image
}

func writeImage(imageToWrite image.Image) {
	f, err := os.Create("resized.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, imageToWrite, nil); err != nil {
		log.Printf("failed to encode: %v", err)
	}

}