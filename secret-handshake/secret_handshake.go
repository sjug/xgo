package secret

func Handshake(i int) []string {
	if i > 31 || i < 1 {
		return nil
	}
	s := make([]string, 0, 4)
	if i&1 != 0 {
		s = append(s, "wink")
	}
	if i&2 != 0 {
		s = append(s, "double blink")
	}
	if i&4 != 0 {
		s = append(s, "close your eyes")
	}
	if i&8 != 0 {
		s = append(s, "jump")
	}
	if i&16 != 0 {
		for j, k := 0, len(s)-1; j < k; j, k = j+1, k-1 {
			s[j], s[k] = s[k], s[j]
		}
	}
	return s
}
