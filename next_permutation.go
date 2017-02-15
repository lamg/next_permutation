// Dijkstra's next permutation algorithm
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var str = flag.String("s", "1234", "string to find next permutation")
	flag.Parse()
	s := []rune(*str)
	b := len(s) > 0
	for i := 1; b && i != len(s); i++ {
		b = s[i-1] >= s[i]
	}
	if !b {
		next_permutation(s)
		println(string(s))
	} else {
		fmt.Fprint(os.Stderr, "Last permutation has no next")
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
		i, j = i+1, j-1
	}
	//tail is ascending (by reversing it)
	//p is the next permutation of the old p
}
