package main

import (
	"log"

	"github.com/rwcarlsen/goexif/exif"
)

type exifData struct {
	CacheName            string    `json:"cacheName"`
	ExposureTime         []float64 `json:"exposureTime"`
	Name                 string    `json:"name"`
	FocalLength          []float64 `json:"focalLength"`
	Make                 string    `json:"make"`
	MediaType            string    `json:"mediaType"`
	DateTime             string    `json:"dateTime"`
	Aperture             []float64 `json:"aperture"`
	ExposureCompensation []float64 `json:"exposureCompensation"`
	Iso                  []int     `json:"iso"`
	DateTimeOriginal     string    `json:"dateTimeOriginal"`
	Date                 string    `json:"date"`
	Model                string    `json:"model"`
	Size                 []int     `json:"size"`
	DateTimeFile         string    `json:"dateTimeFile"`
	Orientation          string    `json:"orientation"`
}

func (picture *Picture) getExif() exifData {
	r := picture.getPictureReader()
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
	case orientationVal == 8:
		orientation = "Vertical (normal)"
	}

	return exifData{
		CacheName:            cacheName,
		Name:                 picture.name,
		Model:                camModel,
		Make:                 make,
		MediaType:            mediaType,
		DateTimeOriginal:     dateTimeOriginal,
		DateTimeFile:         dateTimeFile,
		DateTime:             dateTime,
		Orientation:          orientation,
		ExposureTime:         exposureTime,
		FocalLength:          focalLength,
		Aperture:             aperture,
		ExposureCompensation: exposureCompensation,
		Iso:                  iso,
		Date:                 date,
		Size:                 size,
	}

}
