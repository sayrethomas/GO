package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	//Import geopattern library
	"github.com/pravj/geopattern"
	//Import regex library
	"github.com/zach-klippenstein/goregen"
)

//Log errors and terminate
func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

//Export geopattern to SVG file
func main() {
	//Seed random number generator
	rand.Seed(time.Now().UnixNano())
	//Generate SVG string
	gen, err := patternGenerator(shapes(rand.Intn(15)), phraseGenerator(rand.Intn(48)), hexGenerator(), hexGenerator())
	check(err)
	//Send SVG string to export file
	writer("pattern.svg", gen)
}

//Switch between geopattern shapes
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
	case 15:
		return "xes"
	}
	return "plaid"
}

//Generate random phrase less than 49 chars with goregen
func phraseGenerator(rand int) string {
	//Convert random int to string
	random := strconv.Itoa(rand)
	//Create phrase
	randomPhrase, err := regen.Generate("[A-Z0-9a-z!?$&]{" + random + "}")
	check(err)
	return randomPhrase
}

//Generate random hex color with goregen
func hexGenerator() string {
	//Create hex code
	randomHex, err := regen.Generate("[a-e0-9]{6}")
	check(err)
	randomHex = "#" + randomHex
	return randomHex
}

//Takes SVG string and writes it to SVG file
func writer(fileName string, generatedString string) {
	file, err := os.Create(fileName)
	check(err)
	writer := bufio.NewWriter(file)
	defer file.Close()
	fmt.Fprintln(writer, generatedString)
	writer.Flush()
}

//Creates Geopattern SVG string
func patternGenerator(shape string, phrase string, color string, baseColor string) (string, error) {
	//Handle error if necessary requirement is not provided
	if len(shape) == 0 || len(phrase) == 0 || len(color) == 0 || len(baseColor) == 0 {
		return "Error", fmt.Errorf("One or more of the necessary requirements not provided.\nshape = %s\nphrase = %s\ncolor = %s\nbaseColor = %s", shape, phrase, color, baseColor)
	}
	args := map[string]string{"generator": shape,
		"phrase":    phrase,
		"color":     color,
		"baseColor": baseColor}
	gp := geopattern.Generate(args)
	return gp, nil
}
