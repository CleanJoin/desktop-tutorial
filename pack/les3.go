package pack

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func RegString(s string) (d []string) {

	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	str45 := re.ReplaceAllString(s, " ")
	d = strings.Fields(str45)

	return d
}
func StrToDict(d []string) []string {
	sort.Strings(d)
	var s2 string
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
	delete(cache, "")
	strount := SortWordCountsDict(cache)
	return strount

}

func SortWordCountsDict(cache map[string]int) (str []string) {
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
	for i := 0; i < len(kvs) && i < 10; i++ {
		str2 = append(str2, kvs[i].k)
	}
	return str2
}
