package main

import (
	"Tutorial/pack"
	"fmt"
)

func main() {
	// pack.Lesson1()
	// s := [6]string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	// e := [6]string{"aaaabccddddde", "abcd", "некорректная строка", `qwe45`, `qwe44444`, `qwe\\\\\`}

	// for i, dwarf := range s {
	// 	fmt.Println(string(pack.StringRePack(dwarf)), string(e[i]))
	// }
	s := `слон 
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
	дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`
	d := pack.StrDict(pack.SubString(s))
	fmt.Println(d)

}
