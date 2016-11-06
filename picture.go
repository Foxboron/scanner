package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/nfnt/resize"
)

type PictureInterface interface {
	getPictureReader()
	createThumbnails()
	createThumbnailSize()
	cachePicture()
	createCacheDir()
	getExif()
}

type Picture struct {
	exif      exifData
	name      string
	cacheName string
	content   []byte
	jpg       image.Image
}

var sizes = []uint{72, 150, 640, 1024, 1600}

const cacheDir = "./cache"

func getCacheName(content []byte) string {
	reader := bytes.NewReader(content)
	buf := make([]byte, 2048)
	_, _ = reader.Read(buf)
	sha := sha256.Sum256(buf[:])
	return hex.EncodeToString(sha[:])
}

func createPicture(name string) Picture {
	content, _ := ioutil.ReadFile(name)
	cacheName := getCacheName(content)
	reader := bytes.NewReader(content)
	img, err := jpeg.Decode(reader)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return Picture{
		name:      name,
		cacheName: cacheName,
		content:   content,
		jpg:       img,
	}
}

func (picture *Picture) createThumbnailSize(size uint) {
	m := resize.Resize(size, 0, picture.jpg, resize.Lanczos3)

	picture.createDirectory()
	picture.cachePicture(m, size)

	log.Printf("Created thumbnail %d for %s", size, picture.name)
}

func (picture *Picture) createThumbnails() {
	if _, err := os.Stat(cacheDir + "/" + picture.cacheName); err == nil {
		log.Print("Thumbnails exists. Skipping creation")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(sizes))

	for _, val := range sizes {
		go func(size uint) {
			defer wg.Done()
			picture.createThumbnailSize(size)
		}(val)
	}

	wg.Wait()
}

func (picture *Picture) getPictureReader() io.Reader {
	reader := bytes.NewReader(picture.content)
	return reader
}

func (picture *Picture) cachePicture(m image.Image, size uint) {
	cachePicture := fmt.Sprintf("%s/%s/%d.jpg", cacheDir, picture.cacheName, size)
	out, err := os.Create(cachePicture)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
}

func (picture *Picture) createDirectory() {
	cacheDirectory := fmt.Sprintf("%s/%s", cacheDir, picture.cacheName)
	err := os.MkdirAll(cacheDirectory, 0744)
	if err != nil {
		log.Fatal(err)
	}
}
