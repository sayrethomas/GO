package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/pravj/geopattern"
	"github.com/zach-klippenstein/goregen"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

// Prints pattern's SVG string without any argument
func main() {
	rand.Seed(time.Now().UnixNano())
	gen := patternGenerator(shapes(rand.Intn(15)), "Test Phrase", hexGenerator(), hexGenerator())
	writer("result.svg", gen)

	//fmt.Println(x)
}
func shapes(rand int) string {
	switch rand {
	case 0:
		return "chevrons"
	case 1:
		return "concentric-circles"
	case 2:
		return "diamonds"
	case 3:
		return "hexagons"
	case 4:
		return "mosaic-squares"
	case 5:
		return "nested-squares"
	case 6:
		return "octagons"
	case 7:
		return "overlapping-circles"
	case 8:
		return "overlapping-rings"
	case 9:
		return "plaid"
	case 10:
		return "plus-signs"
	case 11:
		return "sine-waves"
	case 12:
		return "squares"
	case 13:
		return "tessellation"
	case 14:
		return "triangles"
	default:
		return "xes"
	}
}
func hexGenerator() string {
	randomHex, err1 := regen.Generate("[a-f0-9]{6}")
	check(err1)
	randomHex = "#" + randomHex
	return randomHex
}
func writer(fileName string, generatedString string) {
	file, err2 := os.Create(fileName)
	check(err2)
	writer := bufio.NewWriter(file)
	defer file.Close()
	fmt.Fprintln(writer, generatedString)
	writer.Flush()
}
func patternGenerator(shape string, phrase string, color string, baseColor string) string {
	args := map[string]string{"generator": shape,
		"phrase":    phrase,
		"color":     color,
		"baseColor": baseColor}
	gp := geopattern.Generate(args)
	return gp
}
