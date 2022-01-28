package main

import (
	"Tutorial/pack"
)

type Wordseq2 struct {
	Word string
	Swq  int
}

func main() {
	// pack.Lesson1()
	// s := [6]string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	// e := [6]string{"aaaabccddddde", "abcd", "некорректная строка", `qwe45`, `qwe44444`, `qwe\\\\\`}

	// for i, dwarf := range s {
	// 	fmt.Println(string(pack.StringRePack(dwarf)), string(e[i]))
	// }
	// s := `слон
	// 	кот слон
	// 	пес слон кот
	// 	сон слон кот пес
	// 	лом  слон кот пес сон
	// 	ров слон кот пес сон лом
	// 	день слон кот пес сон лом ров
	// 	ночь слон кот пес сон лом ров день
	// 	дочь слон кот пес сон лом ров день ночь
	// 	корч слон кот пес сон лом ров день ночь дочь
	// 	сыч слон кот пес сон лом ров день ночь дочь корч
	// 	клич  слон кот пес сон лом ров день ночь дочь корч сыч
	// 	дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`
	// d := pack.OutputOfTenWords(pack.SortWordCountsFromDictToStruct(pack.CountingWordsinDict(s))
	// c := pack.CountingWordsinDict(`он лом ров день ночь дочь  сыч слон кот пес сон лом ров день ночь дочь корч   клич  слон кот пес сон лом ров день ночь дочь корч сыч   дичь слон кот пес сон лом ров день ночь дочь корч сыч клич`)
	// // intext := map[string]int{
	// 	"день": 4,
	// 	"дичь": 1,
	// 	"дочь": 4,
	// 	"клич": 2,
	// 	"корч": 3,
	// 	"кот":  3,
	// 	"лом":  4,
	// 	"ночь": 4,
	// 	"он":   1,
	// 	"пес":  3,
	// 	"ров":  4,
	// 	"слон": 3,
	// 	"сон":  3,
	// }
	//e := pack.SortWordCountsFromDictToStruct(intext)

	// outdict2 := []pack.Wordseq{{Word: "день", Swq: 4}, {Word: "ночь", Swq: 4}, {Word: "дочь", Swq: 4}, {Word: "ров", Swq: 4}, {Word: "лом", Swq: 4}, {Word: "сон", Swq: 3}, {Word: "корч", Swq: 3}, {Word: "слон", Swq: 3}, {Word: "пес", Swq: 3}, {Word: "сон", Swq: 2}, {Word: "клич", Swq: 1}, {Word: "дичь", Swq: 1}}
	// intext := map[string]int{
	// 	"день": 4,
	// 	"дичь": 1,
	// 	"дочь": 4,
	// 	"клич": 2,
	// 	"корч": 3,
	// 	"кот":  3,
	// 	"лом":  4,
	// 	"ночь": 4,
	// 	"он":   1,
	// 	"пес":  3,
	// 	"ров":  4,
	// 	"слон": 3,
	// 	"сон":  3,
	// }
	// ww := pack.SortWordCountsFromDictToStruct(intext)
	// // e := pack.OutputOfTenWords(outdict2)
	// fmt.Println(ww)

	// закрытый канал
	// pack.ReadChannelClose()
	// pack.WriteChannelClose()
	// Запись и чтение из канала
	// pack.WriteReadChannel()
	// var ch = make(chan int, 3)
	// for i := 0; i < 3; i++ {
	// 	ch <- i
	// 	fmt.Println("Записанно!")
	// }

	// for j := 0; j < 4; j++ {
	// 	fmt.Printf("Считаем: %v ", <-ch)
	// }

	// Создаем новый список
	// и помещаем в него несколько чисел.
	// l := list.New()
	// e4 := l.PushBack(4)
	// e1 := l.PushFront(1)
	// l.InsertBefore(3, e4)
	// l.InsertAfter(2, e1)
	// fmt.Printf("Длина списка %v\n", l.Len())

	// fmt.Println("Итерируем по списку с начала.")
	// for e := l.Front(); e != nil; e = e.Next() {
	//     fmt.Println(e.Value)
	// }

	// fmt.Println("Итерируем по списку с конца.")
	// for e := l.Back(); e != nil; e = e.Prev() {
	//     fmt.Println(e.Value)
	// }
	// pack.WaitGroup()
	// pack.IterArea()
	pack.UseMutex()
}
