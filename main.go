package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

	//Encode qrcode
	// open and decode image file
	fileEncode, _ := os.Open("qrcode.jpg")
	img, _, _ := image.Decode(fileEncode)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result)
}
