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
	gen := patternGenerator("chevrons", "Test Phrase", hexGenerator(), hexGenerator())
	writer("result.svg", gen)

	//fmt.Println(x)
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
