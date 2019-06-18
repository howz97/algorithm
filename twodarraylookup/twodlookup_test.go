package twodarraylookup

import "testing"

func Test_lookup(t *testing.T) {
	tbl := newTable()
	println(tbl.contains(102))
}
