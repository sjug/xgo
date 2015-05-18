package slice

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	subStr := make([]string, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		subStr[i] = s[i : i+n]
	}
	return subStr
}

func Frist(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
