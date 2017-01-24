package main

import (
	"os"
)

var usage = "Usage: %s string"

func main() {
	if len(os.Args) == 2 {
		s := []rune(os.Args[1])
		next_permutation(s)
		println(string(s))
	}	else {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
}

func next_permutation(p []rune) {
	//p is not the last permutation in lexical ordering
	i := len(p) - 2
	for p[i] >= p[i+1] {
		i--
	}
	//max i such p[i] < p[i+1]
	j := len(p) - 1
	for p[j] <= p[i] {
		j--
	}
	//p[j] = min x: x is in p[i+1..]: x
	p[i], p[j] = p[j], p[i]
	//p has the head up to i of the next permutation of the old p
	i, j = i+1, len(p)-1
	for i < j {
		p[i], p[j] = p[j], p[i]
		i, j = i + 1, j - 1
	}
	//tail is ascending (by reversing it)
	//p is the next permutation of the old p
}

func previous_permutation(p []rune){
	//p is not the first one
	i := len(p) - 2
	for p[i] <= p[i+1] {
		i--
	}
	//max i such p[i] > p[i+1]
	j := len(p) - 1
	for p[j] >= p[i] {
		j--
	}
	//p[j] = max x: x is in p[i+1]: x
	p[i], p[j] = p[j], p[i]
	//p has the head up to i of the previous permutation
	i, j = i+1, len(p) - 1
	for i < j {
		p[i], p[j] = p[j], p[i]
		i, j = i + 1, j - 1
	}
	//tail is descending (by reversing it)
	//p is the previous permutation of the old p
}
