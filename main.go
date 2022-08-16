package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("cannot open the video capture")
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Println("cannot read device")
		return
	}

	if img.Empty() {
		fmt.Println("no image")
		return
	}

	// flip xy
	gocv.Flip(img, &img, -1)

	gocv.IMWrite("photo.jpg", img)

}
