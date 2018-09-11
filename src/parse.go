package main

import (
	"strconv"
	"strings"
)

func parseIntString(s string) (r []int) {
	if len(s) == 0 {
		return
	}
	lst := strings.Split(s, ",")
	for _, v := range lst {
		f, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, int(f))
	}
	return
}
