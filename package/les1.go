package package

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)
func lesson1(){
ntpTime, err := ntp.Time("time.apple.com")
		if err != nil {
				fmt.Println(err)
		}

		ntpTimeFormatted := ntpTime.Format(time.UnixDate)

		fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	}