package gomi

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func writeOutputAvg(p Parameters, result float64, label string, output Output) {
	str := fmt.Sprintf("%s\n%f", p.GenerateString("# "), result)

	if p.Verbose {
		fmt.Println(fmt.Sprintf("Result of %s is %f", label, result))
	}

	file, err := os.Create(p.Output)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	defer w.Flush()
	w.WriteString(str)

	if p.LogData {
		name := strings.TrimSuffix(p.Output, filepath.Ext(p.Output))
		name = fmt.Sprintf("%s.json", name)
		output.SetAvgResult(result)
		output.SetParameters(p)
		output.ExportJSON(name)
	}
}

func writeOutputSD(p Parameters, result []float64, label string, output Output) {
	avg := 0.0
	for _, v := range result {
		avg += v
	}

	avg /= float64(len(result))

	if p.Verbose {
		fmt.Println(fmt.Sprintf("Averaged value: %f", avg))
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

	w.WriteString(p.GenerateString("# "))
	w.WriteString("\n")
	w.WriteString(fmt.Sprintf("# Averaged value: %f\n", avg))
	for _, v := range result {
		w.WriteString(fmt.Sprintf("%f\n", v))
	}

	if p.LogData {
		name := strings.TrimSuffix(p.Output, filepath.Ext(p.Output))
		name = fmt.Sprintf("%s.json", name)
		output.SetAvgResult(avg)
		output.SetPointWiseResult(result)
		output.SetParameters(p)
		output.ExportJSON(name)
	}
}
