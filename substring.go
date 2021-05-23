package main

import (
	"fmt"
	"strings"
)

// ExtendedString adds the super useful ContainedIn() method
type ExtendedString string

func compare(longer, shorter string, lindex, sindex int) {
	fmt.Printf("\nComparing at index %d\n", lindex)
	fmt.Println(longer)
	fmt.Println(strings.Repeat(" ", lindex) + "|")
	fmt.Println(shorter)
	fmt.Println(strings.Repeat(" ", sindex) + "|")
}

// ContainedIn returns true if the receiver is a substring of the longer string in params
func (s ExtendedString) ContainedIn(longer string) bool {
	if len(s) == 0 { // empty string is included
		return true
	}
	lrunes := []rune(longer)
	srunes := []rune(s)

	backtrackMap := make(map[rune]int)
	for i, r := range srunes {
		backtrackMap[r] = i // backtrack by i if rune r is found (highest backtrack is kept)
	}

	lindex := 0 // start at the last index of s, no need to waste
	minIndex := 1
	for lindex < len(lrunes) {
		compare(longer, string(s), lindex, 0)
		if lrunes[lindex] == srunes[0] { // we may have a match
			found := true
			for i, r := range srunes {
				if lindex+i >= len(lrunes) || r != lrunes[lindex+i] {
					found = false
					fmt.Printf("Strings didn't match after %d comparisons\n", i+1)
					break
				}
			}
			if found {
				fmt.Printf("Strings matched after %d comparisons\n", len(srunes))
				return true
			}
		}
		if backtrack, ok := backtrackMap[lrunes[lindex]]; ok {
			if lindex-backtrack >= minIndex {
				// backtrack
				fmt.Printf("backtracking by %d to %d because it's an '%s'\n", backtrack, lindex-backtrack, string(lrunes[lindex]))
				lindex -= backtrack
				minIndex = lindex
				fmt.Println("Will then refuse to backtrack below", lindex)
				continue
			}
			fmt.Printf("I would backtrack by %d because there is a '%s', but no backtrack below index %d\n", backtrack, string(lrunes[lindex]), minIndex)
		}
		fmt.Printf("advancing from %d by %d\n", lindex, len(s))
		lindex += len(s) // advance by length of s
	}
	return false
}
