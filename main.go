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

	fmt.Printf("Case Sensitive (Y/N): ")
	var casesens string
	fmt.Scanln(&casesens)

	fmt.Printf("Exclude Articles (Y/N): ")
	var excludearticles string
	fmt.Scanln(&excludearticles)

	totalwordcount := 0
	var counter = struct {
		sync.Mutex
		a map[string]int
	}{a: make(map[string]int)}

	report := "Word Frequency Analysis Report\n----------------------------------\n"

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

			switch casesens {
			case "Y":
				switch excludearticles {
				case "Y":
					for _, w := range s {
						w0 := strings.Trim(w, ",.;:?\"'()!")
						w1 := strings.ToLower(w0)
						if w1 != "a" && w1 != "an" && w1 != "the" {
							counter.Lock()
							counter.a[w0] += 1
							counter.Unlock()
							continue
						}
						totalwordcount -= 1
					}
				case "N":
					for _, w := range s {
						counter.Lock()
						counter.a[strings.Trim(w, ",.;:?\"'()!")] += 1
						counter.Unlock()
					}
				default:
					log.Fatal("Incorrect option provided for article exclusion. Please write only Y or N")
				}

			case "N":
				switch excludearticles {
				case "Y":
					for _, w := range s {
						w1 := strings.ToLower(strings.Trim(w, ",.;:?\"'()!"))
						if w1 != "a" && w1 != "an" && w1 != "the" {
							counter.Lock()
							counter.a[w1] += 1
							counter.Unlock()
							continue
						}
						totalwordcount -= 1
					}
				case "N":
					for _, w := range s {
						counter.Lock()
						counter.a[strings.Trim(w, ",.;:?\"'()!")] += 1
						counter.Unlock()
					}
				default:
					log.Fatal("Incorrect option provided for article exclusion. Please write only Y or N")
				}

			default:
				log.Fatal("Incorrect option provided for case sensitivity. Please write only Y or N")
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
		report += fmt.Sprintf("%d. %s: %d\n", h, k, counter.a[k])
		h++
	}

	report += "----------------------------------\n"
	report += fmt.Sprintf("Total unique words: %d\n", uqcount)
	report += fmt.Sprintf("Total word count: %d\n", totalwordcount)
	report += fmt.Sprintf("Files processed: %d\n", l-1)

	fmt.Printf(report)

	file, err := os.OpenFile("analysis_report.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.Write([]byte(report))
	if err != nil {
		log.Fatal(err)
	}

}
