package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word := strings.Fields(s)
	m := make(map[string]int)
	
	for i:=0; i<len(word); i++ {
		count, found := m[word[i]]
		if(found == true) {
			m[word[i]] = count + 1
		}else {
			m[word[i]] = 1 
		}
			
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
