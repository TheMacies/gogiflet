package main

import (
	"os"
	"image/gif"
	"log"
)


func main() {
	initializePalette()
	initializeLettersTable()
	initializeRand()
	for _,path := range os.Args[1:] {
		inputFile, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		g, err := gif.DecodeAll(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		fullSize = g.Image[0].Bounds()
		g = letterizeGIF(g)

		outputFile, err := os.OpenFile(path + "letterized", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		err = gif.EncodeAll(outputFile, g)
		if err != nil {
			log.Fatal(err)
		}
		inputFile.Close()
		outputFile.Close()
	}
}