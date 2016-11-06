package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/nfnt/resize"
)

type PictureInterface interface {
	getPictureContent()
	createThumbnails()
	createThumbnailSize()
	createCacheDir()
	getCachedPictures()
}

type Picture struct {
	exif      exifData
	name      string
	cacheName string
	content   bytes.Reader
}

var sizes = []uint{72, 150, 640, 1024, 1600}

const cacheDir = "./cache"

func getCacheName(f string) string {
	content, _ := ioutil.ReadFile(f)
	return hex.EncodeToString(content[:2048])
}

func createPicture(name string) Picture {
	cacheName := getCacheName(name)
	return Picture{
		name:      name,
		cacheName: cacheName,
	}
}

func createThumbnailSize(size uint, content []byte) {
	fmt.Println("lol")
	r := bytes.NewReader(content)
	img, err := jpeg.Decode(r)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	m := resize.Resize(size, 0, img, resize.Lanczos3)

	fmt.Println("lol2")
	cacheDirectory := fmt.Sprintf("%s/%s", cacheDir, "test")
	err = os.MkdirAll(cacheDirectory, 0744)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("lol3")
	cachePicture := fmt.Sprintf("%s/%s/%d.jpg", cacheDir, "test", size)
	out, err := os.Create(cachePicture)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Println("lol4")
	jpeg.Encode(out, m, nil)
	log.Printf("Created thumbnail %d for test", size)
}

func (picture *Picture) createThumbnails() {
	for _, v := range sizes {
		go createThumbnailSize(v, picture.getPictureContent())
	}
}

func (picture *Picture) getPictureContent() []byte {
	ret, _ := ioutil.ReadFile(picture.name)
	return ret
}

func cachePicture() {
}

func createDirectory() {
}
