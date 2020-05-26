package leet

func isPalindrome(n int) bool {

	if n < 0 {
		return false
	}

	help := 1

	for n/help > 10 {
		help *= 10
	}

	for n != 0 {
		if n/help != n%10 {
			return false
		}
		n = n % help / 10
		help /= 100
	}
	return true
}
