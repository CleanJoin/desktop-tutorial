package pack

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup - механизм для ожидания завершения работы
// нескольких горутин.
// type WaitGroup struct {
//  // неэкспортируемые поля
// }
// func (wg *WaitGroup) Add(delta int) - инерементирует счетчик
// func (wg *WaitGroup) Done() - декрементит счетчик на 1
// func (wg *WaitGroup) Wait() - блокируется, пока счетчик
type Dog struct {
	name         string
	walkDuration time.Duration
}

func (d Dog) Walk(wg *sync.WaitGroup) {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
	wg.Done()
}
func WaitGroup() {
	dogs := []Dog{{"vasya", time.Second}, {"john", time.Second * 3}}
	wg := new(sync.WaitGroup)
	for _, d := range dogs {
		wg.Add(1)
		go d.Walk(wg)
	}
	wg.Wait()
	fmt.Println("everybody's home")
}

type HttpPkg struct{}

func (HttpPkg) Get(url string) {}

var http HttpPkg

func IterArea() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}
	wg.Wait()
}

// мютексы
var i int // i == 0
func worker(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	i = i + 1
	go mu.Unlock()
	wg.Done()

}
func UseMutex() {
	// runtime.GOMAXPROCS(1) колл процессоров для обработки
	var wg = &sync.WaitGroup{}
	var mu = &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(wg, mu)
	}
	wg.Wait()
	fmt.Println("value of i after 1000 operations is", i) //<1000 без мьютекса
}

// больше 8128 го рутин незя race
// sync почитать про пакет once(выполнить раз)
