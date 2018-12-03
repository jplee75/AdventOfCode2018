package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	var twoCount int
	var threeCount int
	var twoCountFound bool
	var threeCountFound bool

	for _, x := range inputData {
		stringSlice := strings.Split(x, "")
		twoCountFound = false
		threeCountFound = false
		for _, c := range stringSlice {

			if strings.Count(x, c) == 2 && !twoCountFound {
				twoCount++
				twoCountFound = true
			}
			if strings.Count(x, c) == 3 && !threeCountFound {
				threeCount++
				threeCountFound = true
			}
			if threeCountFound && twoCountFound {
				break
			}
		}

	}
	fmt.Printf("The checksum is %v", (twoCount * threeCount))
}
