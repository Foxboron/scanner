package main

import (
	"bytes"
	"log"

	"github.com/rwcarlsen/goexif/exif"
)

type exifData struct {
	cacheName            string    `json:"cacheName"`
	exposureTime         []float64 `json:"exposureTime"`
	name                 string    `json:"name"`
	focalLength          []float64 `json:"focalLength"`
	make                 string    `json:"make"`
	mediaType            string    `json:"mediaType"`
	dateTime             string    `json:"dateTime"`
	aperture             []float64 `json:"aperture"`
	exposureCompensation []float64 `json:"exposureCompensation"`
	iso                  []int     `json:"iso"`
	dateTimeOriginal     string    `json:"dateTimeOriginal"`
	date                 string    `json:"date"`
	model                string    `json:"model"`
	size                 []int     `json:"size"`
	dateTimeFile         string    `json:"dateTimeFile"`
	orientation          string    `json:"orientation"`
}

func (picture *Picture) returnExif() exifData {
	content := picture.getPictureContent()
	r := bytes.NewReader(content)
	x, err := exif.Decode(r)
	if err != nil {
		log.Fatal(err)
	}

	cacheName := picture.cacheName

	// exposureTime
	exposureTimeExif, _ := x.Get(exif.ExposureTime)
	numer, denom, _ := exposureTimeExif.Rat2(0) // retrieve first (only) rat. value
	val := float64(numer) / float64(denom)
	exposureTime := []float64{val}

	// focalLength
	focalLengthExif, _ := x.Get(exif.FocalLength)
	numer, denom, _ = focalLengthExif.Rat2(0) // retrieve first (only) rat. value
	val = float64(numer) / float64(denom)
	focalLength := []float64{val}

	// make
	makeExif, _ := x.Get(exif.Make)
	make, _ := makeExif.StringVal()

	// mediaType
	mediaType := "photo"

	// dateTime
	dateTimeExif, _ := x.Get(exif.DateTime)
	dateTime, _ := dateTimeExif.StringVal()

	// aperture
	apertureExif, _ := x.Get(exif.FNumber)
	numer, denom, _ = apertureExif.Rat2(0) // retrieve first (only) rat. value
	val = float64(numer) / float64(denom)
	aperture := []float64{val}

	// exposureCompensation
	exposureCompensationExif, _ := x.Get(exif.ExposureBiasValue)
	numer, denom, _ = exposureCompensationExif.Rat2(0) // retrieve first (only) rat. value
	val = float64(numer) / float64(denom)
	exposureCompensation := []float64{val}

	// iso
	isoExif, _ := x.Get(exif.ISOSpeedRatings)
	isoVal, _ := isoExif.Int(0)
	iso := []int{isoVal}

	// dateTimeOriginal
	dateTimeOriginalExif, _ := x.Get(exif.DateTimeOriginal)
	dateTimeOriginal, _ := dateTimeOriginalExif.StringVal()

	// date
	date := dateTimeOriginal

	// model
	camModelExif, _ := x.Get(exif.Model)
	camModel, _ := camModelExif.StringVal()

	// size
	XVal, _ := x.Get(exif.PixelXDimension)
	X, _ := XVal.Int(0)
	YVal, _ := x.Get(exif.PixelYDimension)
	Y, _ := YVal.Int(0)
	size := []int{X, Y}

	// dateTimeFile
	dateTimeFileExif, _ := x.Get(exif.DateTime)
	dateTimeFile, _ := dateTimeFileExif.StringVal()

	// orientation
	orientationExif, _ := x.Get(exif.Orientation)
	orientationVal, _ := orientationExif.Int(0)
	orientation := ""
	switch {
	case orientationVal == 1:
		orientation = "Horizontal (normal)"
	case orientationVal == 0:
		orientation = "Vertical (normal)"
	}

	return exifData{
		cacheName:            cacheName,
		name:                 picture.name,
		model:                camModel,
		make:                 make,
		mediaType:            mediaType,
		dateTimeOriginal:     dateTimeOriginal,
		dateTimeFile:         dateTimeFile,
		dateTime:             dateTime,
		orientation:          orientation,
		exposureTime:         exposureTime,
		focalLength:          focalLength,
		aperture:             aperture,
		exposureCompensation: exposureCompensation,
		iso:                  iso,
		date:                 date,
		size:                 size,
	}

}
