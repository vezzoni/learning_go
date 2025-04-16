package dt

import (
	"fmt"
	"time"
)

func rot13(in rune) rune {
	if in >= 'A' && in <= 'Z' {
		return (((in - 'A') + 13) % 26) + 'A'
	}
	if in >= 'a' && in <= 'z' {
		return (((in - 'a') + 13) % 26) + 'a'
	}

	return in
}

func Example4() {
	fmt.Println("Examples 4	")

	now := time.Now()
	fmt.Println(now)

	nanos := now.UnixNano()
	fmt.Println(nanos)

	day := now.Day()
	fmt.Println(day)

	month := now.Month()
	fmt.Println(month)

	year := now.Year()
	fmt.Println(year)

	yearDay := now.YearDay()
	fmt.Println(yearDay)

	hour := now.Hour()
	fmt.Println(hour)

	minute := now.Minute()
	fmt.Println(minute)

	second := now.Second()
	fmt.Println(second)
}
