package main

import (
	"os"
	"log"
	"image/png"
	"image"
	"errors"
	"strconv"
	"math/rand"
	"time"
)

var lettersTable [][][]int
var r *rand.Rand

func initializeRand() {
	src := rand.NewSource(time.Now().Unix())
	r = rand.New(src)
}

func initializeLettersTable () {
	letters := []string{"Q","W","E","R","T","Y","U","I","O","P","A","S","D","F","G","H","V","J","K","V","L","Z","X","C","V","B","N","M"}
	for _,letter := range letters {
		f,err := os.Open("letters/" + letter +".png")
		if err != nil {
			log.Fatal(err)
		}
		im,err := png.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		err = checkSize(im.Bounds())
		if err != nil {
			log.Fatal(errors.New("Error with file : " + letter + ".png :" + err.Error()))
		}
		var letterValues [][]int
		for x:=0;x<pixelStep;x++ {
			var row []int
			for y:=0;y<pixelStep;y++ {
				cl := im.At(x,y)
				if _,_,_,a := cl.RGBA(); a != 0 {
                                  row = append(row,0)
				} else {
					row = append(row,1)
				}
			}
			letterValues = append(letterValues,row)
		}
		lettersTable = append(lettersTable, letterValues)
		f.Close()
	}
}

func checkSize(im image.Rectangle) error {
	if im.Size().X != pixelStep || im.Size().Y != pixelStep {
		return errors.New("Bad image size, got :"  + strconv.Itoa(im.Size().X) + " - " + strconv.Itoa(im.Size().Y) +
		"  expected : " + strconv.Itoa(pixelStep) + " - " + strconv.Itoa(pixelStep))
	}
	return nil
}

func getRandomLetterIndex() int {
	return r.Int()%len(lettersTable)
}