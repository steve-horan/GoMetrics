package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func openFile(p string) *os.File {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func calcCPU(s []string) {
	// Function used to convert the output of parsing proc, and getting our CPU data
	var cpuData []int

	for i, data := range s {
		if i >= 2 {
			f, _ := strconv.Atoi(data)
			cpuData = append(cpuData, f)
		}
	}

	fmt.Println(cpuData)

}

func main() {
	// Gathering some CPU data
	path := "/proc/stat"
	f := openFile(path)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "cpu ") {
			output := strings.Split(scanner.Text(), " ")
			calcCPU(output)
		}
	}
}
