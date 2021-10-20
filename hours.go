package main

import "fmt"

const chineseHours = "子丑寅卯辰巳午未申酉戌亥"

func getChineseHour(hour int) string {

	var cHour rune
	if hour >= 23 || hour < 1 {
		cHour = []rune(chineseHours)[0]
	} else {
		i := (hour + 1) / 2
		cHour = []rune(chineseHours)[i]
	}

	return fmt.Sprintf("%c时", cHour)
}
