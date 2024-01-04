//Ćwiczenie: Images
//Pamiętasz generator obrazów który napisałeś wcześniej? Spróbujmy napisać kolejny, tym razem zamiast zwracać wycinek danych, zwróci on implementację interfejsu image.Image.
//
//Zdefiniuj swój własny typ Image, zaimplementuj wymagane metody, oraz wywołaj pic.ShowImage.
//
//Metoda Bounds powinna zwrócić wartość typu image.Rectangle, przykładowo image.Rect(0, 0, w, h).
//
//Metoda ColorModel powinna zwrócić color.RGBAModel.
//
//Metoda At powinna zwrócić kolor; wartości v w poprzednim generatorze obrazów odpowiada teraz wartość color.RGBA{v, v, 255, 255}.

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x*y + x) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
