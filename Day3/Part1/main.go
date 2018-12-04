package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SantaCloth struct {
	id     int
	startX int
	startY int
	sizeX  int
	sizeY  int
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var elfClaims []SantaCloth
	var bigGrid [1000][1000]uint
	var area uint

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		stripped := strings.Join(strings.Fields(scan.Text()), "")
		id, _ := strconv.Atoi(strings.Split(strings.Split(stripped, "@")[0], "#")[1])
		measurements := strings.Split(stripped, "@")[1]
		inits := strings.Split(strings.Split(measurements, ":")[0], ",")
		sizes := strings.Split(strings.Split(measurements, ":")[1], "x")
		initX, _ := strconv.Atoi(inits[0])
		initY, _ := strconv.Atoi(inits[1])
		sizeX, _ := strconv.Atoi(sizes[0])
		sizeY, _ := strconv.Atoi(sizes[1])
		elfClaims = append(elfClaims, SantaCloth{id, initX, initY, sizeX, sizeY})
	}
	for _, s := range elfClaims {
		for x := s.startX; x < (s.startX + s.sizeX); x++ {

			for y := s.startY; y < (s.startY + s.sizeY); y++ {

				bigGrid[x][y]++

			}
		}

	}
	//now what
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if bigGrid[x][y] > 1 {
				area++
			}
		}
	}
	fmt.Println(area)
}
