package hash_table

import (
	"fmt"
	"testing"
)

func TestChainHT(t *testing.T) {
	ht := New(1)
	t.Log(ht.Size())
	t.Log(ht.TblSize())
	t.Log(ht.LoadFactor())

	ht.Put(Integer(1), 1)
	ht.Put(Integer(2), 2)
	ht.Put(Integer(3), 3)
	ht.Put(Float(3.14), 3.14)
	ht.Put(Float(0.00001), 0.00001)
	ht.Put(Float(99999999999999999999999999999999), float64(99999999999999999999999999999999))
	ht.Put(Str("zhang"), "zhang")
	ht.Put(Str("hao"), "hao")
	ht.Put(Str("你好"), "你好")

	t.Log(ht.Size())
	t.Log(ht.TblSize())
	t.Log(ht.LoadFactor())

	fmt.Println(ht.Get(Integer(1)))
	fmt.Println(ht.Get(Integer(2)))
	fmt.Println(ht.Get(Integer(3)))
	fmt.Println(ht.Get(Float(3.14)))
	fmt.Println(ht.Get(Float(0.00001)))
	fmt.Println(ht.Get(Float(99999999999999999999999999999999)))
	fmt.Println(ht.Get(Str("zhang")))
	fmt.Println(ht.Get(Str("hao")))
	fmt.Println(ht.Get(Str("你好")))

	if ht.Get(Str("世界")) != nil {
		t.Fatal()
	}
	ht.Put(Float(3.14), nil)
	if ht.Get(Float(3.14)) != nil {
		t.Fatal()
	}
	ht.Put(Str("你好"), nil)
	ht.Put(Str("世界"), nil)
	ht.Put(Str("zhang"), nil)
	ht.Put(Str("hao"), nil)
	ht.Put(Integer(1), nil)
	ht.Put(Integer(2), nil)
	ht.Put(Integer(3), nil)
	ht.Put(Float(99999999999999999999999999999999), nil)

	t.Log(ht.Size())
	t.Log(ht.TblSize())
	t.Log(ht.LoadFactor())
}
