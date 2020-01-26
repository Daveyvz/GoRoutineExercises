package word

import "strings"

// UseCount no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count the words in a given string.
func Count(s string) int {
	// write the code for this func
	bs := strings.Split(s, " ")
	count := len(bs)
	return count
}
