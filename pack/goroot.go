package pack

import (
	"fmt"
	"time"
)

var Сh = make(chan struct{}) // не буфферизованный
var Сh2 = make(chan int)     // бцфферизованный

// отправление в канал ch<-x
// получение из канала  x=<-ch или x,ok:=<-ch (ok открыт ли канал0(true,false))
// закрытие канала close(ch)

func ReadChannelClose() {
	var ch = make(chan error)
	close(ch)

	x, ok := <-ch
	fmt.Printf(" x: %v, ok: %v\n", x, ok)
}

// будет panic
func WriteChannelClose() {
	var ch = make(chan int)
	close(ch)

	ch <- 3
}

// отправка блокирует пока не будет выполнено получание и наоборот
// go func() {
// 	fmt.Println("hello")
// 	pack.Сh <-
// }()
// <-pack.Сh
//  однонаправленные каналы: только для записи стрелочка в канал f(out chan<- int), только для чтения стрелочка из канала  f(in <-chan int)

func ReadChannel(in <-chan int) {
	fmt.Printf("%d", <-in)

}

func WriteChannel(out chan<- int) {
	out <- 5

}

func WriteReadChannel() {
	var ch = make(chan int)
	go ReadChannel(ch)
	go WriteChannel(ch)
	time.Sleep(5 * time.Second)
}

// буферезованный канал make( chan int, 3) ёмкость канала cap(ch)
// можно читать из канала for ch:= range channel{}

// Мультиплексирование (select)
// select{
// case<-ch1:
// // ...
// case ch2 <- y:
// 	//....
// default:
// 	//....
// }
// таймеры - сработало через указанное время и тикеры- срабатывает раз в заданное кол-во времени
// таймер
// timer := time.NewTimer(10 * time.Second)
// select {
// case data := <-ch:
//  fmt.Printf("received: %v", data)
// case <-timer.C:
//  fmt.Printf("failed to receive in 10s")
// }

// тикер
// timer := time.NewTimer(10 * time.Second)
// select {
// case data := <-ch:
//  fmt.Printf("received: %v", data)
// case <-timer.C:
//  fmt.Printf("failed to receive in 10s")
// }

// замыкание
// func main() {
// 	for i := 0; i < 5; i++ {
// 	go func() {
// 	fmt.Print(i)
// 	}()
// 	}
// 	time.Sleep(10 * time.Second)
//    }
