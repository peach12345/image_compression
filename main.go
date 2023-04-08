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

type MyImage struct {
	width  int
	height int
}

func main() {

	fmt.Println("---Programm is up and running!!---")

	//Ask to read a single file or multiple files
	filesPaths := chooseBetweenDirectoryOrSingleFiles()
	fmt.Println(filesPaths)

	//If directory, then all pictures should be resized

	//Else single file will be resized

	//Get image max size from user input
	var myImage = getWidthAndHeightFromUserInput()

	//Call read image method

	for i := 0; i < len(filesPaths); i++ {
		imageToShrink, imageConfig := readImageFromFile(filesPaths[i])

		fmt.Println("------")
		fmt.Println("The actual image size: ", imageConfig.Width, "x", imageConfig.Height)
		fmt.Println("Will be converted to:", myImage.width, "x", myImage.height)

		imageToShrink = shrinkImage(imageToShrink, myImage.width, myImage.height)

		writeImage(imageToShrink, i)
		fmt.Println("Successfully resized all image")
	}

}

func readSingleFile() []string {
	var result []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')

	input = strings.Replace(input, "\n", "", -1)
	result = append(result, input)
	return result
}

func readFilesFromUserInput() []string {

	var pathArray []string

	for {
		fmt.Println("Enter path of the picutre! Finished press q")

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')

		input = strings.Replace(input, "\n", "", -1)

		if input == "q" {
			break
		}

		pathArray = append(pathArray, input)
	}

	return pathArray
}
func chooseBetweenDirectoryOrSingleFiles() []string {
	fmt.Println("Press 1 for single file or 2 for enter multiple files!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')

	input = strings.Replace(input, "\n", "", -1)

	switch input {

	case "1":
		return readSingleFile()
	case "2":
		return readFilesFromUserInput()

	}
	var emptyArray []string
	return emptyArray
}

func getFilenameFromUserInput() string {
	fmt.Println("Enter filename")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	filename, _ := reader.ReadString('\n')

	return filename
}

func getWidthAndHeightFromUserInput() MyImage {
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
	var myImage = MyImage{height: h, width: w}
	return myImage
}

// read image and image config from file in asset folder
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

// Shrink image method
// resize image to a given width and high
func shrinkImage(imageToShrink image.Image, width int, height int) image.Image {

	image := resize.Resize(uint(width), uint(height), imageToShrink, resize.MitchellNetravali)

	return image
}

func writeImage(imageToWrite image.Image, index int) {
	f, err := os.Create("resized" + strconv.Itoa(index) + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, imageToWrite, nil); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}

// Udemy
func DeleteFromSlice(a []string, i int) []string {

	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
