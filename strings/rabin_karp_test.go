package strings

import "testing"

func TestIndexRabinKarp(t *testing.T) {
	txt := "kadskdkndkqbabc"
	pattern := "abc"
	i := IndexRabinKarp(txt, pattern)
	if txt[i:i+len(pattern)] != pattern {
		t.Fatalf("not match, i=%d", i)
	}
}
