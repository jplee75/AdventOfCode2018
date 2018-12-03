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

	for _, x := range inputData {
		fmt.Println(x)
		finalCount += x
	}
	fmt.Println(finalCount)

}
