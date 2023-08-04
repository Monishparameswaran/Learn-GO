package main

import (
	"log"
	"sync"
	"time"
)

func downloadImage() {
	log.Println("Download image request sent")
	time.Sleep(time.Second * 3)
}

func getImage(wg *sync.WaitGroup, onceObj *sync.Once) {
	defer wg.Done()

	onceObj.Do(downloadImage)
}

func main() {
	var wg sync.WaitGroup
	onceObj := sync.Once{}

	numGoroutines := 3
	wg.Add(numGoroutines)
	

	for i := 0; i < numGoroutines; i++ {
		go getImage(&wg, &onceObj)
	}

	wg.Wait()

	log.Println("All goroutines completed")
}
