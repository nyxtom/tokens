package tokens

import "sort"

// IntArray represents a sortable collection of int arrays
type intMultiArray [][]int

func (a intMultiArray) Len() int           { return len(a) }
func (a intMultiArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intMultiArray) Less(i, j int) bool { return a[i][0] < a[j][0] }

// SplitNatural to split and return the list of strings tokenized by all common word patterns
func SplitNatural(text string) []string {
	return Split(text, URLIndex, EmailIndex, EmoticonWordPunctIndex)
}

// Split to return the strings by passing the text through a pre-filter prior to a post-filter where the prefix is executed first before continuing to tokenize on the surrounding text (such as in Email & WordPunct)
func Split(text string, filters ...func(t string) [][]int) []string {
	indices := SplitIndex(text, filters...)
	results := make([]string, len(indices))
	sort.Sort(intMultiArray(indices))
	for i, r := range indices {
		results[i] = text[r[0]:r[1]]
	}
	return results
}

// SplitIndex to return the indices of all the tokens that passed through the pre and post filters. The prefix is executed first and the postfix is executed on the surrounding text to the tokens found by the prefix filter.
func SplitIndex(text string, filters ...func(t string) [][]int) [][]int {
	var results [][]int
	search := [][]int{[]int{0, len(text)}}
	for _, f := range filters {
		var newSearch [][]int
		for _, t := range search {
			indices := f(text[t[0]:t[1]])
			index := t[0]
			for _, match := range indices {
				if match[0] > index && index < t[1] {
					newSearch = append(newSearch, []int{index, match[0]})
				}
				results = append(results, []int{t[0] + match[0], t[0] + match[1]})
				index = match[1]
			}
			if index < t[1] {
				newSearch = append(newSearch, []int{index, t[1]})
			}
		}
		search = newSearch
	}
	return results
}
