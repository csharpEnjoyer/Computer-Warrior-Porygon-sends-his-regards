package main

import (
	"image"
	"image/color"
	"os"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	height int = 0
	width  int = 0
)

func main() {
	a := app.New()
	w := a.NewWindow("Filter")

	fimg, _ := os.Open("a.png")
	defer fimg.Close()
	img, _, _ := image.Decode(fimg)

	height = img.Bounds().Dy()
	width = img.Bounds().Dx()

	imga := canvas.NewImageFromImage(img)
	imga.FillMode = canvas.ImageFillOriginal

	r := 20.0
	g := 20.0
	b := 20.0

	red := binding.BindFloat(&r)
	green := binding.BindFloat(&g)
	blue := binding.BindFloat(&b)

	slider1 := widget.NewSliderWithData(1, 255, red)
	slider2 := widget.NewSliderWithData(1, 255, green)
	slider3 := widget.NewSliderWithData(1, 255, blue)

	label1 := widget.NewLabelWithData(
		binding.FloatToString(red),
	)
	label2 := widget.NewLabelWithData(
		binding.FloatToString(green),
	)
	label3 := widget.NewLabelWithData(
		binding.FloatToString(blue),
	)

	w.SetContent(
		container.NewVBox(
			imga,
			label1,
			slider1,
			label2,
			slider2,
			label3,
			slider3,
		),
	)
	go func() {
		for {
			updateImage(imga, label1, label2, label3)
		}

	}()

	w.ShowAndRun()
}

func updateImage(img *canvas.Image, r *widget.Label, g *widget.Label, b *widget.Label) {

	var re string = r.Text
	var gr string = g.Text
	var bl string = b.Text

	red, _ := strconv.ParseFloat(re, 0)
	blue, _ := strconv.ParseFloat(gr, 0)
	green, _ := strconv.ParseFloat(bl, 0)

	img.Image = generateImage(img.Image, uint8(red), uint8(blue), uint8(green))
	img.Refresh()
}

func generateImage(imga image.Image, r uint8, g uint8, b uint8) image.Image {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			red, green, blue, a := imga.At(x, y).RGBA()
			col := color.RGBA{uint8(red) + r, uint8(green) + g, uint8(blue) + b, uint8(a)}
			img.Set(x, y, col)
		}
	}
	return img
}
