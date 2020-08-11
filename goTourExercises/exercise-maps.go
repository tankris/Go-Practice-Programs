package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	stringSplit := strings.Fields(s)
	mapSplit := make(map[string]int, len(stringSplit))

	for _, i := range stringSplit {
		mapSplit[i]++
	}

	return mapSplit
}

func main2() {
	wc.Test(WordCount)
}
