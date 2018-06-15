package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func writeOutputAvg(p Parameters, result float64, label string) {
	str := fmt.Sprintf("%s: %f", label, result)

	if p.Verbose {
		fmt.Println(str)
	}

	file, err := os.Create(p.Output)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	defer w.Flush()
	w.WriteString(str)
}

func writeOutputSD(p Parameters, result []float64, label string) {
	if p.Verbose {
		n := 10
		if n > len(result)-1 {
			n = len(result) - 1
		}
		fmt.Println(fmt.Sprintf("%s: %v (only %d of %d values shown)", label, result[0:n], n, len(result)))
	}

	file, err := os.Create(p.Output)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString(fmt.Sprintf("# %s\n", label))
	for _, v := range result {
		w.WriteString(fmt.Sprintf("%f\n", v))
	}
}
