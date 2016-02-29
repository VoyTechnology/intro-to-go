package functions

import (
	"fmt"
)

func convertIntToWeekDayString(dayNumber int) (day string) {
	switch dayNumber % 7 {
	case 0:
		return "Monday"
	case 1:
		return "Tuesday"
	case 2:
		return "Wednesday"
	case 3:
		return "Thursday"
	case 4:
		return "Friday"
	case 5:
		return "Saturday"
	case 6:
		return "Sunday"
	}

	return ""
}

// GetWeekDay is publically exposed function
func GetWeekDay(day int) (string, error) {
	if day < 0 || day > 7 {
		return "", fmt.Errorf("Weekday invalid")
	}
	return convertIntToWeekDayString(day), nil
}

// SumVersion1 sums in 1st way you can do it
func SumVersion1(a, b int) int {
	return a + b
}

// SumVersion2 sums in 2nd way you can do it
func SumVersion2(a int, b int) int {
	return a + b
}

// SumVersion3 sums in 2nd way you can do it
func SumVersion3(a int, b int) (result int) {
	result = a + b
	return result
}
