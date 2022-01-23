package pack

import (
	"fmt"
	"testing"
)

func TestSubStr(t *testing.T) {
	s := [2]string{`жопа жопа дрын ок ок сын сын дин брат брат брат нет нет нет да да да да жиза жиза жиза жиза про про про про сиси сиси линивец линивец`, `жопа жопа дрын ок ок сын сын& дин брат брат брат нет нет нет да да да! да жиза жиза жиза жиза про про про про сиси сиси,линивец,  линивец`}
	e := []string{"да", "жиза", "про", "брат", "нет", "жопа", "ок", "сиси", "линивец", "дрын"}

	for i, dwarf := range s {
		fmt.Println(string(StringRePack(dwarf)), string(dwarf[i]))
		if c := StrToDict(FormatTextToLowerAndReplace(dwarf)); c != nil {
			t.Fatalf("bad count for %s: got %s expected %s", s, c, e)
		}
	}
}
