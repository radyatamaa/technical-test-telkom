package main

import (
	"fmt"
	"strings"
)


const base = 16777619

func Search(txt string, patterns []string) ([]string,[]string) {
	in := indices(txt, patterns)
	matches := make([]string, len(in))
	notMatchs := make([]string,0)
	i := 0
	for j, p := range patterns {
		if _, ok := in[j]; ok {
			matches[i] = p
			i++
		}else {
			notMatchs = append(notMatchs,p)
		}

	}

	return matches,notMatchs
}

func indices(txt string, patterns []string) map[int]int {
	n, m := len(txt), minLen(patterns)
	matches := make(map[int]int)

	if n < m || len(patterns) == 0 {
		return matches
	}

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := hashPatterns(patterns, m)
	h := hash(txt[:m])
	for i := 0; i < n-m+1 && len(hp) > 0; i++ {
		if i > 0 {
			h = h - mult*uint32(txt[i-1])
			h = h*base + uint32(txt[i+m-1])
		}

		if mps, ok := hp[h]; ok {
			for _, pi := range mps {
				pat := patterns[pi]
				e := i + len(pat)
				if _, ok := matches[pi]; !ok && e <= n && pat == txt[i:e] {
					matches[pi] = i
				}
			}
		}
	}
	return matches
}

func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func hashPatterns(patterns []string, l int) map[uint32][]int {
	m := make(map[uint32][]int)
	for i, t := range patterns {
		h := hash(t[:l])
		if _, ok := m[h]; ok {
			m[h] = append(m[h], i)
		} else {
			m[h] = []int{i}
		}
	}

	return m
}

func minLen(patterns []string) int {
	if len(patterns) == 0 {
		return 0
	}

	m := len(patterns[0])
	for i := range patterns {
		if m > len(patterns[i]) {
			m = len(patterns[i])
		}
	}

	return m
}

func main() {
	input1 := "telkom"
	input2 := "telecom"
	var result bool
	anyCharmatches,notMatches := Search(input1, strings.Split(input2,""))

	if len(notMatches) == 0 {
		result = true
	}

	fmt.Println("result : ",result)
	fmt.Println("character any match : ",anyCharmatches)
}

