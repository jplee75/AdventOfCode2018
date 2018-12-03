package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var inputData []string

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		inputData = append(inputData, scan.Text())
	}

	var differences int

	for _, x := range inputData {
		for _, y := range inputData {
			for i := 0; i < len(x); i++ {
				if x[i] != y[i] {
					differences++
				}
				if differences > 1 {
					break
				}
			}
			if differences == 1 {
				fmt.Println("String 1 " + x + " String 2 " + y)
			} else {
				differences = 0
			}

		}
	}
}
