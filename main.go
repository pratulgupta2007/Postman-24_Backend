package main

import (
	"fmt"
	"log"
	"os"
)

var dict = map[string]int{}

func main() {

	if len(os.Args) == 1 {
		log.Fatal("No file provided!")
	}

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(readfile(os.Args[i]))
	}

	fmt.Println("Word Frequency Analysis Report")
}

func readfile(dir string) string {
	f, err := os.ReadFile(dir)
	if err != nil {
		log.Fatal(err)
	}
	return string(f)
}
