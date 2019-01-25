package main

import (
	"github.com/CameronJHall/image-processing/idx"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	"image/color"
	"log"
	"time"
)

var data idx.IDXData

func main(){
	labelsFileName := "./train-labels-idx1-ubyte"
	dataFileName := "./train-images-idx3-ubyte"

	err := data.ParseLabels(labelsFileName)
	err = data.ParseData(dataFileName)
	if err != nil {
		log.Fatal(err)
	}

	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Numbers :)",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(16, 16))


	nImage := image.NewRGBA(image.Rect(0, 0, data.DataDimensions[1], data.DataDimensions[2]))
	for num:=0; num<data.DataDimensions[0]; num++{
		nImage = image.NewRGBA(image.Rect(0, 0, data.DataDimensions[1], data.DataDimensions[2]))
		for i:=0; i<data.DataDimensions[1]; i++ {
			for j := 0; j < data.DataDimensions[1]; j++ {
				tPixel := uint8(data.Data[(num*data.DataDimensions[1]*data.DataDimensions[2]) + (j*data.DataDimensions[1]) + i])
				nImage.Set(i, j, color.RGBA{255-tPixel, 255-tPixel, 255-tPixel, 0})
			}
		}

		win.Clear(color.Black)
		sprite := pixel.NewSprite(pixel.PictureDataFromImage(nImage), pixel.PictureDataFromImage(nImage).Bounds())
		sprite.Draw(win, mat)
		time.Sleep(time.Duration(1)*time.Second)
		win.Update()
	}



}