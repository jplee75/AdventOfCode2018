package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type GuardInfo struct {
	totalSleep   int
	minuteDetail map[int]int
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var guardMinutes map[int]GuardInfo

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		stripped := strings.Join(strings.Fields(scan.Text()), "")
		id := strings.Split(stripped)
	}
}
