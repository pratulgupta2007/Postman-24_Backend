package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {

	l := len(os.Args)
	if l == 1 {
		log.Fatal("No file provided!")
	}

	totalwordcount := 0
	var counter = struct {
		sync.Mutex
		a map[string]int
	}{a: make(map[string]int)}

	fmt.Println("Word Frequency Analysis Report")
	fmt.Println("----------------------------------")

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
			totalwordcount += len(s)
		}()
	}
	wg.Wait()

	uqcount := len(counter.a)
	keys := make([]string, 0, uqcount)
	for k := range counter.a {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return counter.a[keys[i]] > counter.a[keys[j]]
	})

	h := 1
	for _, k := range keys[:15] {
		fmt.Printf("%d. %s: %d\n", h, k, counter.a[k])
		h++
	}

	fmt.Println("----------------------------------")
	fmt.Println("Total unique words:", uqcount)
	fmt.Println("Total word count:", totalwordcount)
	fmt.Println("Files processed:", l-1)

}
