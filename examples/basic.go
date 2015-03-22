package main

import (
	"fmt"
	"github.com/yqingp/rocksdbgo"
	"os"
)

func main() {
	db, _ := rocksdbgo.Open("./a", nil)
	db.Put(nil, []byte("a1"), []byte("v"))
	db.Put(nil, []byte("a2"), []byte("v"))
	db.Put(nil, []byte("a3"), []byte("v"))

	db.Put(nil, []byte("b1"), []byte("v"))
	db.Put(nil, []byte("c1"), []byte("v"))
	db.Put(nil, []byte("d1"), []byte("v"))
	db.Put(nil, []byte("d2"), []byte("v"))

	db.Get(nil, []byte("a1"))

	db.Delete(nil, []byte("a1"))

	it := db.NewIterator(nil, true, "", "") //forward
	// it := db.NewIterator(nil, false, "", "") //backward

	for {
		if v, valid := it.Next(); valid {
			fmt.Println(v)
		} else {
			break
		}
	}
	it.Close()
	db.Close()
	os.RemoveAll("./a")
}
