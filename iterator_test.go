package rocksdbgo

import (
	"fmt"
	. "github.com/yqingp/rocksdbgo"
	"os"
	"testing"
)

func TestIterator(t *testing.T) {
	db, err := Open("./a", nil)
	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	db.Put(nil, []byte("a1"), []byte("v1"))
	db.Put(nil, []byte("a2"), []byte("v2"))
	db.Put(nil, []byte("a3"), []byte("v"))

	db.Put(nil, []byte("b1"), []byte("v"))
	db.Put(nil, []byte("c1"), []byte("v"))
	db.Put(nil, []byte("d1"), []byte("v"))
	db.Put(nil, []byte("d2"), []byte("v"))

	it := db.NewIterator(nil, true, "", "")
	// it := db.NewIterator(nil, false, "", "")

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
