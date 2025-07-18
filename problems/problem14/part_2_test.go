package problem14

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		palindrome   string
		isPalindrome bool
	}{
		{"radar", true},
		{"hello", false},
		{"bukub", true},
		{"rasdasd", false},
	}
	for _, v := range cases {
		if v.isPalindrome != IsPalindrome(v.palindrome) {
			t.Error("Ошибка")
		}
	}
}
