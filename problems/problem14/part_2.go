package problem14

/*
 * 14. Тестирование - Задание 2
 * Напишите функцию isPalindrome(s string) bool.
 * Напишите тест, проверяющий палиндромы ("radar", "hello").
 */
func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	if len(s)%2 == 0 {
		leftSide := string(s[0 : len(s)/2])
		rightSide := string(s[len(s)/2:])
		if leftSide == rightSide {
			return true
		} else {
			return false
		}
	} else {
		leftSide := string(s[0 : (len(s)-1)/2])
		rightSide := string(s[(len(s)+1)/2:])
		if leftSide == reverseString(rightSide) {
			return true
		} else {
			return false
		}
	}
}
