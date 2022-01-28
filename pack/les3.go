package pack

import (
	"regexp"
	"sort"
	"strings"
)

type Wordseq struct {
	Word string
	Swq  int
}

func FormatTextToLowerAndReplace(text string) string {
	lowertext := strings.ToLower(text)
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	formattext := re.ReplaceAllString(lowertext, " ")
	return formattext
}
func CountingWordsinDict(formattext string) map[string]int {
	newwods := strings.Fields(formattext)
	sort.Strings(newwods)
	var previousword string
	wodscounter := make(map[string]int)
	count := 1
	for _, word := range newwods {
		if word == previousword {
			count++
		} else if word != previousword && previousword != "" {
			wodscounter[previousword] = count
			count = 1
		}
		previousword = word
	}
	return wodscounter

}

func SortWordCountsFromDictToStruct(wodscounter map[string]int) []string {
	numberofwords := make([]Wordseq, 0, len(wodscounter))
	for k, v := range wodscounter {
		numberofwords = append(numberofwords, Wordseq{k, v})
	}
	sort.Slice(numberofwords, func(i, j int) bool {

		return numberofwords[i].Swq > numberofwords[j].Swq
	})
	sort.SliceStable(numberofwords, func(i, j int) bool {

		return numberofwords[i].Word < numberofwords[j].Word && numberofwords[i].Swq == numberofwords[j].Swq
	})
	Outwords := OutputOfTenWords(numberofwords)
	return Outwords
}
func OutputOfTenWords(numberofwords []Wordseq) []string {
	var readywords []string
	for i := 0; i < len(numberofwords) && i < 10; i++ {
		readywords = append(readywords, numberofwords[i].Word)
	}
	return readywords
}
