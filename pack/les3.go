package pack

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func SubString(s string) (d []string) {
	// s := `жопа жопа дрын ок ок сын сын& дин брат брат брат нет нет нет да да да! да жиза жиза жиза жиза про про про про сиси сиси,линивец,  линивец`
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	str45 := re.ReplaceAllString(s, " ")
	// fmt.Println(str45)
	d = strings.Fields(str45)
	// fmt.Println(d[5])
	return d
}
func StrDict(d []string) []string {
	sort.Strings(d)
	// fmt.Println(d)
	var s2 string
	// intslice := []int{}
	// stringslice := []string{}
	cache := make(map[string]int)

	count := 1
	index := 0

	for _, s := range d {
		if s == s2 {
			count++
		} else if s != s2 {
			cache[s2] = count

			count = 1
			index++
		}
		s2 = s
	}
	strount := StrDeleteDict(cache)

	return strount

}

func StrDeleteDict(cache map[string]int) (str []string) {
	delete(cache, "")
	type kv struct {
		k string
		v int
	}
	var str2 []string
	kvs := make([]kv, 0, len(cache))
	for k, v := range cache {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].v > kvs[j].v
	})
	fmt.Println(kvs)
	for i := 0; i < 10; i++ {
		// fmt.Println(kvs[i].k)
		str2 = append(str2, kvs[i].k)
	}
	return str2
}
