package string_search

const primeRK = 16777619

func IndexRabinKarp(s, substr string) int {
	m := len(substr)
	hp, ht := hashStr(substr), hashStr(s[:m])
	if hp == ht && s[:m] == substr {
		return 0
	}
	rpm := pow(primeRK, uint32(m-1))
	for i := m; i < len(s); i++ {
		ht = (ht-uint32(s[i-m])*rpm)*primeRK + uint32(s[i])
		if ht == hp && s[i-m+1:i+1] == substr {
			return i - m + 1
		}
	}
	return -1
}

func hashStr(sep string) uint32 {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	return hash
}

func pow(x, y uint32) uint32 {
	if y < 0 {
		panic("y must >= 0")
	}
	if y == 0 {
		return 1
	}
	v := uint32(1)
	for i := uint32(0); i < y; i++ {
		v *= x
	}
	return v
}
