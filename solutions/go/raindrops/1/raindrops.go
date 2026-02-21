package raindrops

import "strconv"

func Convert(number int) string {
	rules := []struct {
		divisor int
		sound   string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	result := ""

	for _, rule := range rules {
		if number%rule.divisor == 0 {
			result += rule.sound
		}
	}

	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}