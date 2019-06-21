package regexp

func isMatch(s, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) != 0 && len(p) == 0 {
		return false
	}
	if len(s) == 0 && len(p) != 0 {
		if p[len(p)-1] == '*' {
			return isMatch(s, p[:len(p)-2])
		}
		return false
	}

	switch p[len(p)-1] {
	case '*':
		if isMatch(s, p[:len(p)-2]) {
			return true
		}
		for i := 1; len(s)-i >= 0 && (s[len(s)-i] == p[len(p)-2] || p[len(p)-2] == '.'); i++ {
			if isMatch(s[:len(s)-i], p[:len(p)-2]) {
				return true
			}
		}
		return false
	case '.':
		return isMatch(s[:len(s)-1], p[:len(p)-1])
	default:
		if s[len(s)-1] == p[len(p)-1] {
			return isMatch(s[:len(s)-1], p[:len(p)-1])
		}
		return false
	}
}
