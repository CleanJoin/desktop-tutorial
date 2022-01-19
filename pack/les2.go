package pack

import (
	"fmt"
	"strconv"
	"strings"
)

func StringRePack() {
	// s := "a4bc2d5e"
	s := `wwrr\44\5\9`
	b := ""
	a := ""
	d := ""
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[0])); err == nil {
			break
		}
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			fmt.Println(d, string(s[i]))
			if (d != string(s[i]) && (d != `\`)) || (d == string(s[i-2])) {

				count, _ := strconv.Atoi(string(s[i]))
				a = strings.Repeat(string(s[i-1]), count-1)
				b = b + a
			} else {
				b = b + string(s[i])

			}

		} else if (string(s[i]) != `\` && string(s[i]) != d) || (d == `\`) {
			b = b + string(s[i])

		}
		d = string(s[i])

	}
	if b == "" {
		fmt.Printf("некорректная строка")
	} else {
		fmt.Println(b)
	}
	// 	* "a4bc2d5e" => "aaaabccddddde"
	// * "abcd" => "abcd"
	// * "45" => "" (некорректная строка)

}
func StringEscape() {
	// Дополнительное задание: поддержка escape - последовательности
	// * `qwe\4\5` => `qwe45` (*)
	// * `qwe\45` => `qwe44444` (*)
	// * `qwe\\5` => `qwe\\\\\` (*)
}
