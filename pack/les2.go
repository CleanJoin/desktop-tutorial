package pack

import (
	"strconv"
	"strings"
)

func StringRePack(s string) (e string) {
	// s := "a4bc2d5e"
	// fmt.Println(s)
	var b strings.Builder
	a := ""
	d := ""
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[0])); err == nil {
			break
		}
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			// fmt.Println(d, string(s[i]), string(s[i-2]))
			if (d != string(s[i]) && (d != `\`)) || (d == string(s[i-2])) {

				count, _ := strconv.Atoi(string(s[i]))
				a = strings.Repeat(string(s[i-1]), count-1)
				b.WriteString(a)
			} else {
				b.WriteString(string(s[i]))

			}

		} else if (string(s[i]) != `\` && string(s[i]) != d) || (d == `\`) {
			b.WriteString(string(s[i]))

		}
		d = string(s[i])

	}
	if b.String() == "" {
		b.WriteString("некорректная строка")
	}
	// 	* "a4bc2d5e" => "aaaabccddddde"
	// * "abcd" => "abcd"
	// * "45" => "" (некорректная строка)
	return b.String()
}
func StringEscape() {
	// Дополнительное задание: поддержка escape - последовательности
	// * `qwe\4\5` => `qwe45` (*)
	// * `qwe\45` => `qwe44444` (*)
	// * `qwe\\5` => `qwe\\\\\` (*)
}
