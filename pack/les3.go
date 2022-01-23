package pack

import (
	"regexp"
	"sort"
	"strings"
)

type wordseq struct {
	word string
	swq  int
}

func FormatTextToLowerAndReplace(text string) string {
	lowertext := strings.ToLower(text)
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	formattext := re.ReplaceAllString(lowertext, " ")

	return formattext
}
func StrToDict(formattext string) []string {
	newwods := strings.Fields(formattext)
	sort.Strings(newwods)
	var previousword string
	wodscounter := make(map[string]int)
	count := 1
	for _, word := range newwods {
		if word == previousword {
			count++
		} else if word != previousword {
			wodscounter[previousword] = count
			count = 1
		}
		previousword = word
	}
	delete(wodscounter, "")
	strount := SortWordCountsFromDictToStruct(wodscounter)
	return strount

}

func SortWordCountsFromDictToStruct(wodscounter map[string]int) []string {
	numberofwords := make([]wordseq, 0, len(wodscounter))
	for k, v := range wodscounter {
		numberofwords = append(numberofwords, wordseq{k, v})
	}
	sort.Slice(numberofwords, func(i, j int) bool {
		return numberofwords[i].swq > numberofwords[j].swq
	})
	readywords := OutputOfTenWords(numberofwords)
	return readywords
}
func OutputOfTenWords(numberofwords []wordseq) []string {
	var readywords []string
	for i := 0; i < len(numberofwords) && i < 10; i++ {
		readywords = append(readywords, numberofwords[i].word)
	}
	return readywords
}
