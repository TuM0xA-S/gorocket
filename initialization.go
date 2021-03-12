package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"image"
	_ "image/png"
	"log"

	"github.com/TuM0xA-S/gorocket/resources"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/opentype"
)

//fonts
var (
	FontBig     font.Face
	FontSmaller font.Face
)

func init() {
	tt, err := opentype.Parse(gomedium.TTF)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	FontSmaller, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    30,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	FontBig, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	GameOverTextX = getCentredPosForText(GameOverText)
	NewRecordTextX = getCentredPosForText(NewRecordText)
}

//some centered consts
var (
	GameOverTextX  float64
	NewRecordTextX float64
)

//RocketImg is a pic of rocket
var (
	RocketImgBase image.Image
	RocketImgH    int
	RocketImgW    int
)

func init() {
	var err error
	RocketImgBase, _, err = image.Decode(bytes.NewReader(resources.RocketPNG))
	if err != nil {
		log.Fatal(err)
	}

	lowerRightCorner := RocketImgBase.Bounds().Max
	RocketImgW = lowerRightCorner.X
	RocketImgH = lowerRightCorner.Y
}

//StartSplashImg used at startup state
var (
	StartSplashImg *ebiten.Image
)

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(resources.StartupPNG))
	if err != nil {
		log.Fatal(err)
	}
	StartSplashImg = ebiten.NewImageFromImage(img)
}

//Cipher using for record file
var Cipher cipher.Block

func init() {
	var err error
	Cipher, err = aes.NewCipher(SecretKey)
	if err != nil {
		log.Fatal(err)
	}
}
