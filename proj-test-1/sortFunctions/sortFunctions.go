package sortfunctions

import (
	"fmt"
	"sort"
	"strings"
)


func CharecterCounter() {

	str := `This is a
multiline
string.cccccccccccccccccccccccccccccccccccccccccc
strings are nice and multi multi line is also some nice string with nice stuff 
non logic here so dont keep on reading
"u can but quotes in here too =) " `

	cahrs := map[string]int{}
	for _, cahr := range strings.Split(str, "") {
		cahrs[cahr] +=1 
	}

	fmt.Println(cahrs)
}


func WordsCounter() map[string]int {

	str := `This is a
multiline
string
strings are nice and multi multi line is also some nice string with nice stuff 
non logic here so dont keep on reading nice nice nice
u can but quotes in here too`

	words := map[string]int{}
	for _, word := range strings.Fields(str) {
		words[word] +=1 
	}

	fmt.Println(words)
	return words
}

func SortTop3Words(words map[string]int) {

	type Word struct {
		word string
		count int 
	}

	wordsCount := []Word{}
	for k, v := range words {
		wordsCount = append(wordsCount, Word{k, v})
	}

	fmt.Printf("before: %v \n", wordsCount)

	sort.Slice(wordsCount, func(i, j int) bool {
		return wordsCount[i].count > wordsCount[j].count
	})
	fmt.Printf("after sorting: %v \n", wordsCount)
	fmt.Printf("top 3 words: %v, \n", wordsCount[0:3])

}