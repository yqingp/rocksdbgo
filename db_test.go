package rocksdbgo

import (
	"bytes"
	// "strings"
	"testing"
)

func TestOpen(t *testing.T) {
	// _, err := Open("./a")

	// if err == nil {
	// 	fmt.Println("db opened")
	// }

	// db.Get([]byte("a"))
	// db.Put([]byte("a"), []byte("b"))
}

func TestPut(t *testing.T) {
	db, err := Open("./a", nil)

	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	db.Put(nil, []byte("a"), []byte("b"))
	db.Close()
}

func TestGet(t *testing.T) {
	db, err := Open("./a", nil)

	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	v, err := db.Get(nil, []byte("a"))

	if !bytes.Equal([]byte("b"), v) {
		t.Errorf("expected [%s], actual [%s]; error %s", "a", v, err)
	}

	db.Close()
}
