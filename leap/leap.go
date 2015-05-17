package leap

const TestVersion int = 1

func IsLeapYear(year int) bool {
	switch {
	case year%4 == 0:
		// Edge case detection
		if year%100 == 0 && year%400 != 0 {
			break
		}
		return true
	}
	return false
}
