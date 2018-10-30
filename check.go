package gomi

import (
	"fmt"
	"os"
)

func checkW(d Data) {
	if len(d.W) == 0 {
		fmt.Println("W is empty")
		os.Exit(0)
	}
}

func checkA(d Data) {
	if len(d.A) == 0 {
		fmt.Println("A is empty")
		os.Exit(0)
	}
}

func checkS(d Data) {
	if len(d.S) == 0 {
		fmt.Println("S is empty")
		os.Exit(0)
	}
}

func checkWd(d Data) {
	if len(d.Discretised.W) == 0 {
		fmt.Println("Discrete W is empty")
		os.Exit(0)
	}
}

func checkAd(d Data) {
	if len(d.Discretised.A) == 0 {
		fmt.Println("Discrete A is empty")
		os.Exit(0)
	}
}

func checkSd(d Data) {
	if len(d.Discretised.S) == 0 {
		fmt.Println("Discrete S is empty")
		os.Exit(0)
	}
}
