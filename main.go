package main

import (
	"log"
	"math"
	"math/rand"

	"github.com/CameronJHall/image-processing/idx"
	"gonum.org/v1/gonum/mat"
)

//val toPrint Chan

func main() {
	labelsFileName := "./train-labels-idx1-ubyte"
	dataFileName := "./train-images-idx3-ubyte"

	data := idx.IDXData{}

	err := data.ParseLabels(labelsFileName)
	err = data.ParseData(dataFileName)
	if err != nil {
		log.Fatal(err)
	}

	m := matrixBuilder(data)

	// ldss := makeLDSets(data.Labels, matrices)
	net := NNet{}
	net.Init([]int{784, 16, 10})
	net.FeedForward(m[0].RowView(0))

	//go pixelgl.Run(run())
}

func matrixBuilder(data idx.IDXData) (matrices []mat.Matrix) {
	matSize := data.DataDimensions[1] * data.DataDimensions[2]
	for i := 0; i < data.DataDimensions[0]; i++ {
		v := data.Data[i*matSize : (i+1)*matSize]
		a := mat.NewDense(matSize, 1, v)
		matrices = append(matrices, a)
	}
	return
}

func makeLDSets(labels []float64, matrices []mat.Matrix) (ldss []ldSet) {
	for i, m := range matrices {
		s := mat.SVD{}
		s.Factorize(m, mat.SVDNone)
		val := make([]float64, 28)
		ldss = append(ldss, ldSet{labels[i], s.Values(val)})
	}
	return
}

//func run() {
//	cfg := pixelgl.WindowConfig{
//		Title:  "Numbers :)",
//		Bounds: pixel.R(0, 0, 1024, 768),
//		VSync: true,
//	}
//	win, err := pixelgl.NewWindow(cfg)
//	if err != nil {
//		panic(err)
//	}
//
//	mat := pixel.IM
//	mat = mat.Moved(win.Bounds().Center())
//	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(16, 16))
//
//
//	nImage := image.NewRGBA(image.Rect(0, 0, data.DataDimensions[1], data.DataDimensions[2]))
//	for num:=0; num<data.DataDimensions[0]; num++{
//		nImage = image.NewRGBA(image.Rect(0, 0, data.DataDimensions[1], data.DataDimensions[2]))
//		for i:=0; i<data.DataDimensions[1]; i++ {
//			for j := 0; j < data.DataDimensions[1]; j++ {
//				tPixel := uint8(data.Data[(num*data.DataDimensions[1]*data.DataDimensions[2]) + (j*data.DataDimensions[1]) + i])
//				nImage.Set(i, j, color.RGBA{255-tPixel, 255-tPixel, 255-tPixel, 0})
//			}
//		}
//
//		win.Clear(color.Black)
//		sprite := pixel.NewSprite(pixel.PictureDataFromImage(nImage), pixel.PictureDataFromImage(nImage).Bounds())
//		sprite.Draw(win, mat)
//		//time.Sleep(time.Duration(1)*time.Second)
//		win.SetTitle(fmt.Sprintf("%s | Index: %d", cfg.Title, num))
//		win.Update()
//	}
//}

type ldSet struct {
	label float64
	eVal  []float64
}

type NNet struct {
	LayerSizes []int
	Biases     []float64
	Weights    [][]float64
}

func (n *NNet) Init(LayerSizes []int) {
	n.LayerSizes = LayerSizes
	n.Biases = make([]float64, len(LayerSizes))
	n.Weights = make([][]float64, len(LayerSizes))
	for i, val := range LayerSizes {
		n.Biases[i] = rand.Float64()
		n.Weights[i] = make([]float64, val)
		for j := 0; j < val; j++ {
			n.Weights[i][j] = rand.Float64()
		}
	}
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

func (n *NNet) FeedForward(v mat.Vector) mat.Matrix {

}
