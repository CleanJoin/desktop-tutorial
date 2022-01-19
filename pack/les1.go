package pack

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func Lesson1() {
	ntpTime, err := ntp.Time("time.apple.com")
	if err != nil {
		fmt.Println(err)
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
}
