package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	imageFilePath := "Lena.png"
	img := gocv.IMRead(imageFilePath, gocv.IMReadGrayScale)
	if img.Empty() {
		log.Panic("Can not read Image file : ", imageFilePath)
		return
	}

	log.Println("read", imageFilePath)
	log.Println("size:", img.Size())
	log.Println("type:", img.Type())

	img1 := gocv.NewMat()
	defer img1.Close()

	gx := gocv.NewMat()
	defer gx.Close()

	gy := gocv.NewMat()
	defer gy.Close()

	gocv.Laplacian(img, &img1, 200, 5, 1, 0, gocv.BorderConstant)
	gocv.Canny(img, &img1, 100, 200)
	gocv.Sobel(img, &img1, 3, 0, 1, 7, 1, 0, gocv.BorderConstant)

	window := gocv.NewWindow("Lena")
	window.IMShow(img)
	window.IMShow(img1)
	window.WaitKey(0)
}
