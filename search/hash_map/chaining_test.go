package hash_map

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"testing"
)

func TestChainHT(t *testing.T) {
	ht := New(1)
	t.Log(ht.Size())
	t.Log(ht.TblSize())
	t.Log(ht.LoadFactor())

	ht.Put(search.Integer(1), 1)
	ht.Put(search.Integer(2), 2)
	ht.Put(search.Integer(3), 3)
	ht.Put(search.Float(3.14), 3.14)
	ht.Put(search.Float(0.00001), 0.00001)
	ht.Put(search.Float(99999999999999999999999999999999), float64(99999999999999999999999999999999))
	ht.Put(search.Str("zhang"), "zhang")
	ht.Put(search.Str("hao"), "hao")
	ht.Put(search.Str("你好"), "你好")

	t.Log(ht.Size())
	t.Log(ht.TblSize())
	t.Log(ht.LoadFactor())

	fmt.Println(ht.Get(search.Integer(1)))
	fmt.Println(ht.Get(search.Integer(2)))
	fmt.Println(ht.Get(search.Integer(3)))
	fmt.Println(ht.Get(search.Float(3.14)))
	fmt.Println(ht.Get(search.Float(0.00001)))
	fmt.Println(ht.Get(search.Float(99999999999999999999999999999999)))
	fmt.Println(ht.Get(search.Str("zhang")))
	fmt.Println(ht.Get(search.Str("hao")))
	fmt.Println(ht.Get(search.Str("你好")))

	if ht.Get(search.Str("世界")) != nil {
		t.Fatal()
	}
	ht.Put(search.Float(3.14), nil)
	if ht.Get(search.Float(3.14)) != nil {
		t.Fatal()
	}
	ht.Put(search.Str("你好"), nil)
	ht.Put(search.Str("世界"), nil)
	ht.Put(search.Str("zhang"), nil)
	ht.Put(search.Str("hao"), nil)
	ht.Put(search.Integer(1), nil)
	ht.Put(search.Integer(2), nil)
	ht.Put(search.Integer(3), nil)
	ht.Put(search.Float(99999999999999999999999999999999), nil)

	t.Log(ht.Size())
	t.Log(ht.TblSize())
	t.Log(ht.LoadFactor())
}
