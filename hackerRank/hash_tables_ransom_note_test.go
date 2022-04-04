package hackerrank

import (
	"fmt"
	"testing"
)

func checkMagazine(magazine []string, note []string) {
	// Write your code here
	var result string = "Yes"
	if len(magazine) < len(note) {
		result = "No"
	} else {
		magazineMap := make(map[string]int)
		for _, mag := range magazine {
			if magazineMap[mag] >= 1 {
				magazineMap[mag] += 1
			} else {
				magazineMap[mag] = 1
			}
		}
		for _, n := range note {
			if magazineMap[n] == 0 {
				result = "No"
				break
			} else {
				magazineMap[n] -= 1
			}
		}
	}
	fmt.Println(result)
}
func TestCheckMagazine(t *testing.T) {
	magazine := []string{"two", " times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}
	checkMagazine(magazine, note)
}
