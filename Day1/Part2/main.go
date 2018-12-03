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
	var match = false
	dupFrequencies := map[int]bool{0: true}

	for !match {
		for _, x := range inputData {
			finalCount += x
			if _, ok := dupFrequencies[finalCount]; ok {
				fmt.Printf("The duplier is : %v\n", finalCount)
				match = true
				break
			}
			dupFrequencies[finalCount] = true
		}
		loop += 1
		fmt.Println(loop)
	}
	var repeat = finalCount
	fmt.Println("First Repeated Value = " + strconv.Itoa(repeat))

}
