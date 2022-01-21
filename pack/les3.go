package pack

import (
	"fmt"
	"strings"
)

func SubString() {
	s := `жопа жопа2, жопа3, aa - жопа! кискис`
	// string[] splitted = text.Split(new[] { '-', '.', '?', '!', ')', '(', ',', ':' }, StringSplitOptions.RemoveEmptyEntries);
	// char symbol = 'a';
	// fmt.Println("test")
	g := strings.Split(s, "-", ".", '?', '!', ')', '(', ',', ':')

	out := strings.Replace(s, "lamb", "", -1)
	fmt.Println(out)
	d := strings.Fields(s)
	fmt.Println(d[1])

}
