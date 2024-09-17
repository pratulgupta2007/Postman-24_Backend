package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Word Frequency Analysis Report")
	s1 := readfile("./sample1.txt")
	s2 := readfile("./sample2.txt")
	s3 := readfile("./sample3.txt")
	fmt.Println(s1, s2, s3)

}

func readfile(dir string) string {
	f, err := os.ReadFile(dir)
	if err != nil {
		panic((err))
	}
	s := string(f)
	return s
}
