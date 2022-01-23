package pack

import (
	"fmt"
	"testing"
)

func TestFormatTextToLowerAndReplace(t *testing.T) {
	s := [2]string{`слон 
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
	e := [2]string{`он лом ров день ночь дочь  сыч слон кот пес сон лом ров день ночь дочь корч   клич  слон кот пес сон лом ров день ночь дочь корч сыч   дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`,
		"жиза"}

	for i, dwarf := range s {
		fmt.Println(string(StringRePack(dwarf)), string(dwarf[i]))
		if c := FormatTextToLowerAndReplace(dwarf); c != e[i] {
			t.Fatalf("bad count for %s: got %s expected %s", s, c, e)
		}
	}
}
func TestCountingWordsinDict(t *testing.T) {
	s := [2]string{`он лом ров день ночь дочь  сыч слон кот пес сон лом ров день ночь дочь корч   клич  слон кот пес сон лом ров день ночь дочь корч сыч   дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`, `жопа жопа дрын ок ок сын сын& дин брат брат брат нет нет нет да да да! да жиза жиза жиза жиза про про про про сиси сиси,линивец,  линивец`}
	e := [2]string{"да", "жиза"}

	for i, dwarf := range s {
		fmt.Println(string(StringRePack(dwarf)), string(dwarf[i]))
		if c := CountingWordsinDict(dwarf); c != nil {
			t.Fatalf("bad count for %s: got %s expected %s", s, c, e)
		}
	}
}
func TestSortWordCountsFromDictToStruct(t *testing.T) {
	s := [2]string{`жопа жопа дрын ок ок сын сын дин брат брат брат нет нет нет да да да да жиза жиза жиза жиза про про про про сиси сиси линивец линивец`, `жопа жопа дрын ок ок сын сын& дин брат брат брат нет нет нет да да да! да жиза жиза жиза жиза про про про про сиси сиси,линивец,  линивец`}
	e := []string{"да", "жиза", "про", "брат", "нет", "жопа", "ок", "сиси", "линивец", "дрын"}

	for i, dwarf := range s {
		fmt.Println(string(StringRePack(dwarf)), string(dwarf[i]))
		if c := SortWordCountsFromDictToStruct(dwarf); c != nil {
			t.Fatalf("bad count for %s: got %s expected %s", s, c, e)
		}
	}
}
func TestOutputOfTenWords(t *testing.T) {
	s := [2]string{`жопа жопа дрын ок ок сын сын дин брат брат брат нет нет нет да да да да жиза жиза жиза жиза про про про про сиси сиси линивец линивец`, `жопа жопа дрын ок ок сын сын& дин брат брат брат нет нет нет да да да! да жиза жиза жиза жиза про про про про сиси сиси,линивец,  линивец`}
	e := []string{"да", "жиза", "про", "брат", "нет", "жопа", "ок", "сиси", "линивец", "дрын"}

	for i, dwarf := range s {
		fmt.Println(string(StringRePack(dwarf)), string(dwarf[i]))
		if c := OutputOfTenWords(dwarf); c != nil {
			t.Fatalf("bad count for %s: got %s expected %s", s, c, e)
		}
	}
}
