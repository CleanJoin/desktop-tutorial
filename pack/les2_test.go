package pack

import (
	"testing"
)

func TestCount(t *testing.T) {
	s := [6]string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	e := [6]string{"aaaabccddddde", "abcd", "некорректная строка", `qwe45`, `qwe44444`, `qwe\\\\\`}

	for i, dwarf := range s {
		// fmt.Println(string(StringRePack(dwarf)), string(dwarf[i]))
		if c := StringRePack(dwarf); c != e[i] {
			t.Fatalf("bad count for %s: got %s expected %s", s, c, e)
		}
	}
}
