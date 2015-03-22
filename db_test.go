package rocksdbgo

import (
	// "bytes"
	// "strings"
	"fmt"
	"os"
	"testing"
)

// func TestOpen(t *testing.T) {
// }

// func TestPut(t *testing.T) {
// 	db, err := Open("./a", nil)

// 	if err != nil {
// 		t.Errorf("db open error: %s", err)
// 	}

// 	db.Put(nil, []byte("a"), []byte("b"))
// 	db.Close()
// }

// func TestGet(t *testing.T) {
// 	db, err := Open("./a", nil)

// 	if err != nil {
// 		t.Errorf("db open error: %s", err)
// 	}

// 	v, err := db.Get(nil, []byte("a"))

// 	if !bytes.Equal([]byte("b"), v) {
// 		t.Errorf("expected [%s], actual [%s]; error %s", "a", v, err)
// 	}

// 	db.Close()
// }

func TestIterator(t *testing.T) {
	db, err := Open("./a", nil)
	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	db.Put(nil, []byte("a1"), []byte("v"))
	db.Put(nil, []byte("a2"), []byte("v"))
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
