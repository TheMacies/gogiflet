package main

import (
	"image/gif"
	"image/color"
	"fmt"
	"image"
)

var pal color.Palette
var fullSize image.Rectangle

func initializePalette() {
	//transparency color
	cl := color.RGBA{}
	cl.R = uint8(0)
	cl.G = uint8(0)
	cl.B = uint8(0)
	cl.A = uint8(0)
	pal = append(pal,cl)

	//I am trying to make us of as much colors as i can and I want to them to vary
	for i :=0;i<6;i++ {
		for j:=0;j<6;j++{
			for k:=0;k<6;k++{
				cl := color.RGBA{}
				cl.R = uint8(10 + 40*i)
				cl.G = uint8(10 + 40*j)
				cl.B = uint8(10 + 40*k)
				cl.A = uint8(255)
				pal = append(pal,cl)
			}
		}
	}

	cl = color.RGBA{}
	cl.R = uint8(255)
	cl.G = uint8(255)
	cl.B = uint8(255)
	cl.A = uint8(255)
	pal = append(pal,cl)
}

func letterizeGIF(g *gif.GIF) *gif.GIF {
	for i,im := range g.Image {
		g.Image[i] = letterizeFrame(im)
		fmt.Println("frames letterized :" ,i+1,"/" ,len(g.Image) )
	}
	return g
}


func letterizeFrame(im *image.Paletted) *image.Paletted {
	var avgColors []color.Color
	//Find color of each little square
	for x := 0;; x+=pixelStep {
		xBound := x + pixelStep
		if xBound > fullSize.Size().X {
			xBound = fullSize.Size().X
		}
		for y := 0;; y+=pixelStep {
			yBound := y + pixelStep
			if yBound > fullSize.Size().Y {
				yBound = fullSize.Size().Y
			}
			avgColors = append(avgColors, getAverageColor(im,x,y,xBound,yBound))
			if yBound == fullSize.Size().Y{
				break
			}
		}
		if xBound == fullSize.Size().X{
			break
		}
	}
	clIndex := 0
	//Create new image based on colors and random letters
	newImage := image.NewPaletted(fullSize, pal)
	for x := 0;; x+=pixelStep {
		xBound := x + pixelStep
		if xBound > fullSize.Size().X {
			xBound = fullSize.Size().X
		}
		for y := 0;; y+=pixelStep {
			yBound := y + pixelStep
			if yBound > fullSize.Size().Y {
				yBound = fullSize.Size().Y
			}
			l := getRandomLetterIndex()
			drawLetter(newImage,x,y,xBound,yBound,avgColors[clIndex],l)
			clIndex++
			if yBound == fullSize.Size().Y{
				break
			}
		}
		if xBound == fullSize.Size().X{
			break
		}
	}
	return newImage
}

func getAverageColor(im *image.Paletted, x0,y0,x,y int) color.Color{
	//I just pick a color in the middle of a square

	if !image.Pt((x+x0)/2,(y+y0)/2).In(im.Bounds()) {
		return nil
	}
	cl := im.At((x+x0)/2,(y+y0)/2)
	if _,_,_,a := cl.RGBA(); a ==0 {
		return  nil
	}

	return cl
}

func drawLetter(im *image.Paletted, x0,y0,xEnd,yEnd int,cl color.Color, letterIndex int) {
	if cl == nil {
		return
	}

	index := im.Palette.Index(im.Palette.Convert(cl))
	for x:=x0;x<xEnd;x++ {
		for y:=y0;y<yEnd;y++ {
			if lettersTable[letterIndex][x-x0][y-y0] == 0 {
				im.SetColorIndex(x, y, uint8(index))
			} else {
				im.SetColorIndex(x,y,0)
			}
		}
	}
}
