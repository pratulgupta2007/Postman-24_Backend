package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var dict = map[string]int{}

func main() {

	l := len(os.Args)

	if l == 1 {
		log.Fatal("No file provided!")
	}

	fmt.Println("Word Frequency Analysis Report")

	c := make(chan string)
	counts := make(chan int)

	go func() {

		wg := sync.WaitGroup{}
		for i := 1; i < l; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				f, err := os.ReadFile(os.Args[i])
				if err != nil {
					log.Fatal(err)
				}
				s := string(f)
				c <- s
				counts <- len(strings.Fields(s))
			}()
		}
		wg.Wait()
		close(c)
		close(counts)
	}()

	totalwordcount := 0
	for i := range c {
		fmt.Println(i)
		j := <-counts
		totalwordcount += j
	}
	fmt.Println("Total word count:", totalwordcount)
	fmt.Println("Files processed:", l-1)

}
