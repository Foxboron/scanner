package main

import "sync"

func main() {
	album := getAlbums()

	var wg sync.WaitGroup
	wg.Add(len(album))

	for _, v := range album {
		go func(v Album) {
			defer wg.Done()
			v.writeAlbumCache()
		}(v)
	}
	wg.Wait()
}
