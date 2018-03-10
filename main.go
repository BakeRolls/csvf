package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s foo.csv Hello %%1!]", os.Args[0])
	}

	format := strings.Join(os.Args[2:], " ")
	re := regexp.MustCompile("%([0-9]+)")
	indices := []int{}
	for _, v := range re.FindAllStringSubmatch(format, -1) {
		i, err := strconv.Atoi(v[1])
		if err != nil {
			log.Fatalf("could not parse int: %v", err)
		}
		indices = append(indices, i)
	}
	format = re.ReplaceAllString(format, "%s")

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("unexpected error in file: %v", err)
			}
			break
		}
		values := make([]interface{}, len(indices))
		for i, j := range indices {
			if i >= len(record) {
				log.Fatalf("index %d bigger than max value %d", j, len(record)-1)
			}
			values[i] = record[j]
		}
		fmt.Printf(format+"\n", values...)
	}
}
