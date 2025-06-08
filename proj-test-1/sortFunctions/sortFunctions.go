package sortfunctions

import (
	"fmt"
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

type Words []Word

func (_words Words) Len() int {
	return len(_words)
}

func (_words Words) Less(i int, j int) bool {
	return _words[i].count > _words[j].count
}

func (_words Words) Swap(i int, j int) {
	_words[i], _words[j] = _words[j], _words[i]
}

// example how to solr a basic slice of strings.
func SortStringsExample() {
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Printf("sorting strings resault: %v \n", names)
}

func CharecterCounter() {

	str := `This is a
multiline
string.cccccccccccccccccccccccccccccccccccccccccc
strings are nice and multi multi line is also some nice string with nice stuff 
non logic here so dont keep on reading
"u can but quotes in here too =) " `

	cahrs := map[string]int{}
	for _, cahr := range strings.Split(str, "") {
		cahrs[cahr] += 1
	}

	fmt.Println(cahrs)
}

//          *** example 1 ***

// creates a map of words (string) as keys and accurances of the word as value (int)
func WordsCounter() map[string]int {

	str := `This is a
multiline
string
strings are nice and multi multi line is also some nice string with nice stuff 
non logic here so dont keep on reading nice nice nice
u can but quotes in here tooare nice are they not nice ?`

	words := map[string]int{}
	for _, word := range strings.Fields(str) {
		words[word] += 1
	}

	fmt.Println(words)
	return words
}

// this example show how to sort with "sort.Slice" when u have a struct with a comparable feild
// in this case word "count ""
func SortTop3Words(words map[string]int) {

	wordsCount := []Word{}
	for word, count := range words {
		wordsCount = append(wordsCount, Word{word, count})
	}

	fmt.Printf("before: %v \n", wordsCount)

	sort.Slice(wordsCount,
		func(i, j int) bool {
			return wordsCount[i].count > wordsCount[j].count
		})
	fmt.Printf("after sorting: %v \n", wordsCount)
	fmt.Printf("top 3 words: %v, \n", wordsCount[0:3])

}

//          *** example 2 ***

// in this example i am implementing the sort interface of go.
func SortWords(words map[string]int) {
	var wordsCount Words

	for word, count := range words {
		wordsCount = append(wordsCount, Word{word, count})
	}

	fmt.Printf("before sorting v2: %v \n", wordsCount)
	sort.Sort(wordsCount)
	fmt.Printf("after sorting v2: %v \n", wordsCount)
	sort.Sort(sort.Reverse(wordsCount))
	fmt.Printf("after Reverse v2: %v \n", wordsCount)
}
