package pack

import (
	"reflect"
	"testing"
)

func TestFormatTextToLowerAndReplace(t *testing.T) {
	intext := [2]string{`слон
	кот слон
	пес слон кот
	сон слон кот пес
	лом  слон кот пес сон
	ров слон кот пес сон лом
	день слон кот пес сон лом ров
	ночь слон кот пес сон лом ров день
	дочь слон кот пес сон лом ров день ночь
	корч слон кот пес сон лом ров день ночь дочь
	сыч слон кот пес сон лом ров день ночь дочь корч
	клич  слон кот пес сон лом ров день ночь дочь корч сыч
	дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`,
		`жопа жопа дрын ок ок сын сын& дин брат брат брат нет нет нет да да да! да жиза жиза жиза жиза про про про про сиси сиси,линивец,  линивец`}
	outtext := [2]string{`слон  кот слон  пес слон кот  сон слон кот пес  лом  слон кот пес сон  ров слон кот пес сон лом  день слон кот пес сон лом ров  ночь слон кот пес сон лом ров день  дочь слон кот пес сон лом ров день ночь  корч слон кот пес сон лом ров день ночь дочь  сыч слон кот пес сон лом ров день ночь дочь корч  клич  слон кот пес сон лом ров день ночь дочь корч сыч  дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`,
		`жопа жопа дрын ок ок сын сын  дин брат брат брат нет нет нет да да да  да жиза жиза жиза жиза про про про про сиси сиси линивец   линивец`}

	for i, dwarf := range intext {

		if c := FormatTextToLowerAndReplace(dwarf); c != outtext[i] {
			t.Fatalf("bad count for %s: got %s expected %s", intext, c, outtext)
		}
	}
}

func TestCountingWordsinDict(t *testing.T) {
	intext := [1]string{`он лом ров день ночь дочь  сыч слон кот пес сон лом ров день ночь дочь корч   клич  слон кот пес сон лом ров день ночь дочь корч сыч   дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`}
	outdict := map[string]int{
		"день": 4,
		"дичь": 1,
		"дочь": 4,
		"клич": 2,
		"корч": 3,
		"кот":  3,
		"лом":  4,
		"ночь": 4,
		"он":   1,
		"пес":  3,
		"ров":  4,
		"слон": 3,
		"сон":  3,
	}
	for _, text := range intext {

		inputfunc := CountingWordsinDict(text)

		comparedict := reflect.DeepEqual(inputfunc, outdict)
		// fmt.Println(inputfunc, "\n", outdict, res1)
		if comparedict == false {
			t.Fatalf("bad count for %s: got %q expected %q", intext, outdict, inputfunc)
		}

	}
}

func TestSortWordCountsFromDictToStruct(t *testing.T) {
	intext := map[string]int{
		"день": 4,
		"дичь": 1,
		"дочь": 4,
		"клич": 2,
		"корч": 3,
		"кот":  3,
		"лом":  4,
		"ночь": 4,
		"он":   1,
		"пес":  3,
		"ров":  4,
		"слон": 3,
		"сон":  3,
	}
	outdict := make([]Wordseq, 0, len(intext))
	outdict = []Wordseq{{word: "дочь", swq: 4}, {word: "лом", swq: 4}, {word: "ночь", swq: 4}, {word: "ров", swq: 4}, {word: "день", swq: 4}, {word: "сон", swq: 3}, {word: "кот", swq: 3}, {word: "корч", swq: 3}, {word: "слон", swq: 3}, {word: "пес", swq: 3}, {word: "клич", swq: 2}, {word: "дичь", swq: 1}, {word: "он", swq: 1}}

	if c := SortWordCountsFromDictToStruct(intext); c != nil {
		t.Fatalf("bad count for %s: got %s expected %s", s, c, e)

	}
}

// func TestOutputOfTenWords(t *testing.T) {
// 	s := [2]string{`жопа жопа дрын ок ок сын сын дин брат брат брат нет нет нет да да да да жиза жиза жиза жиза про про про про сиси сиси линивец линивец`, `жопа жопа дрын ок ок сын сын& дин брат брат брат нет нет нет да да да! да жиза жиза жиза жиза про про про про сиси сиси,линивец,  линивец`}
// 	e := []string{"да", "жиза", "про", "брат", "нет", "жопа", "ок", "сиси", "линивец", "дрын"}

// 	for i, dwarf := range s {
// 		fmt.Println(string(StringRePack(dwarf)), string(dwarf[i]))
// 		if c := OutputOfTenWords(dwarf); c != nil {
// 			t.Fatalf("bad count for %s: got %s expected %s", s, c, e)
// 		}
// 	}
// }
