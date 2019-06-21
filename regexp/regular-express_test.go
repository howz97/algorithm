package regexp

import "testing"

func Test_isMatch(t *testing.T) {
	match := isMatch("bbbc", "ab*c")
	if match {
		t.Error("test failed")
	}
	match = isMatch("abbbc", "ab*c")
	if !match {
		t.Error("test failed")
	}
	match = isMatch("ab", ".*")
	if !match {
		t.Error("test failed")
	}
	match = isMatch("aab", "c*a*b")
	if !match {
		t.Error("test failed")
	}
}
