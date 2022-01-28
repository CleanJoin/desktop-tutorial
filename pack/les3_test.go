package pack

import (
	"reflect"
	"testing"
)

// +
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

// +
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

func TestOutputOfTenWords(t *testing.T) {
	inwords := []Wordseq{{Word: "день", Swq: 4}, {Word: "ночь", Swq: 4}, {Word: "дочь", Swq: 4}, {Word: "ров", Swq: 4}, {Word: "лом", Swq: 4}, {Word: "сон", Swq: 3}, {Word: "корч", Swq: 3}, {Word: "слон", Swq: 3}, {Word: "пес", Swq: 3}, {Word: "сон", Swq: 2}, {Word: "клич", Swq: 1}, {Word: "дичь", Swq: 1}}

	outtext := []string{"день", "ночь", "дочь", "ров", "лом", "сон", "корч", "слон", "пес", "сон"}

	c := OutputOfTenWords(inwords)
	if reflect.DeepEqual(c, outtext) == false {
		t.Fatalf("bad count for %q: got %q expected %q", inwords, c, outtext)
	}

}

func TestSortWordCountsFromDictToStruct(t *testing.T) {
	intext := map[string]int{
		"день": 4,
		"дичь": 1,
		"дочь": 7,
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
	// outdict := make([]Wordseq, 0, len(intext))
	// outdict := []Wordseq{{word: "день", swq: 4}, {word: "ночь", swq: 4}, {word: "дочь", swq: 4}, {word: "ров", swq: 4}, {word: "лом", swq: 4}, {word: "сон", swq: 3}, {word: "корч", swq: 3}, {word: "слон", swq: 3}, {word: "пес", swq: 3}, {word: "сон", swq: 2}, {word: "клич", swq: 1}, {word: "дичь", swq: 1}}
	outtext := []string{"дочь", "день", "лом", "ночь", "ров", "корч", "кот", "пес", "слон", "сон"}

	comparestruct := SortWordCountsFromDictToStruct(intext)
	if reflect.DeepEqual(comparestruct, outtext) == false {
		t.Fatalf("bad count for %q: got %q expected %q", intext, comparestruct, outtext)
	}
}
