package main

func main() {
	fname := "./bilder/DSC03328.JPG"
	pic := createPicture(fname)
	pic.createThumbnails()
}
