package main

import (
	"fmt"
	"github.com/pravj/geopattern"
	"os"
	"bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Prints pattern's SVG string without any argument
func main() {
	args := map[string]string{"generator": "chevrons"}
	gp := geopattern.Generate(args)
	
	file, err := os.Create("result.svg")
	check(err)
	
	writer := bufio.NewWriter(file)
	defer file.Close()

	fmt.Fprintln(writer, gp)
	writer.Flush()
	//fmt.Println(gp)
}