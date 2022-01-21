package main

import (
	"Tutorial/pack"
	"fmt"
)

func main() {
	pack.Lesson1()
	s := [6]string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	e := [6]string{"aaaabccddddde", "abcd", "некорректная строка", `qwe45`, `qwe44444`, `qwe\\\\\`}

	for i, dwarf := range s {
		fmt.Println(string(pack.StringRePack(dwarf)), string(e[i]))
	}
	pack.SubString()

}
