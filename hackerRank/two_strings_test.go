package hackerrank

import (
	"strings"
	"testing"
)

func twoStrings(s1 string, s2 string) string {
	// Write your code here
	s1slice := strings.Split(s1, "")
	s2slice := strings.Split(s2, "")
	s1map := make(map[string]string)
	for _, v := range s1slice {
		s1map[v] = v
	}
	for _, s := range s2slice {
		if s == s1map[s] {
			return "YES"
		}
	}
	return "NO"
}

func TestTwoStrings(t *testing.T) {
	twoStrings("hello", "world")

}
