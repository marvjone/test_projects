package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, ferr := os.Open("myfile.csv")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
