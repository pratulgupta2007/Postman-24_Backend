package main

import (
	"fmt"
	"log"
	"os"
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
				c <- string(f)
			}()
		}
		wg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Printf(i)
	}

}
