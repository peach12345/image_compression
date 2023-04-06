package imageModels

import "fmt"

type MyImage struct {
	Width       int
	Height      int
	PictureName string
}

type ImagesToShrink struct {
	AllImages []MyImage
}

func (e *ImagesToShrink) HighResolution() []MyImage {

	var highResolution []MyImage

	for _, x := range e.AllImages {
		if x.Width > 1920 && x.Height > 1080 {
			highResolution = append(highResolution, x)
		}
	}

	return highResolution

}

func (mI MyImage) ShowDetails() {
	fmt.Printf("Picture size is %dx%d", mI.Width, mI.Height)
}
