package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var inputData []int

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		i, _ := strconv.Atoi(scan.Text())
		inputData = append(inputData, i)
	}

	var finalCount int
	var loop int
	var frequencies []int
	var match = false
	dupFrequencies := map[int]bool{0: true}

	for !match {
		for _, x := range inputData {
			finalCount += x
			frequencies = append(frequencies, finalCount)
			if _, ok := dupFrequencies[finalCount]; ok {
				fmt.Printf("The duplier is : %v\n", finalCount)
				match = true
				break
			}
			dupFrequencies[finalCount] = true
			//fmt.Println(finalCount)
			//sort.Ints(frequencies)
			//y := sort.SearchInts(frequencies, finalCount)
			//if loop >= 1 && y != len(frequencies) {
			//	if finalCount == frequencies[y+1] {
			//		match = true
			//		break
			//	} else {
			//		frequencies = append(frequencies, finalCount)
			//	}
			//} else {
			//
			//}
		}
		loop += 1
		fmt.Println(loop)
	}
	var repeat = finalCount
	fmt.Println("First Repeated Value = " + strconv.Itoa(repeat))

}
