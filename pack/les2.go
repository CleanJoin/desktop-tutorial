package pack

import (
	"strconv"
	"strings"
)

func StringRePack(s string) (e string) {
	var b strings.Builder
	a := ""
	d := ""
	for i := 0; i < len(s); i++ {
		if _, err := strconv.Atoi(string(s[0])); err == nil {
			break
		}
		if _, err := strconv.Atoi(string(s[i])); err == nil {
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
	return b.String()
}
