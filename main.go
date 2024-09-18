package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {

	l := len(os.Args)

	if l == 1 {
		log.Fatal("No file provided!")
	}
	var counter = struct {
		sync.Mutex
		a map[string]int
	}{a: make(map[string]int)}
	fmt.Println("Word Frequency Analysis Report")

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
				s := strings.Fields(string(f))
				for _, w := range s {
					counter.Lock()
					counter.a[strings.ToLower(strings.Trim(w, ",.;:?\"'()!"))] += 1
					counter.Unlock()
				}
				counts <- len(s)
			}()
		}
		wg.Wait()
		close(counts)
	}()

	totalwordcount := 0
	for j := range counts {
		totalwordcount += j
	}
	fmt.Println(counter.a)
	fmt.Println("Total word count:", totalwordcount)
	fmt.Println("Files processed:", l-1)

}
