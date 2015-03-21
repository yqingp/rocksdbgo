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
	db, err := Open("./a")

	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	option := NewWriteOption()

	db.Put(option, []byte("a"), []byte("b"))
	option.Close()
	db.Close()
}

func TestGet(t *testing.T) {
	db, err := Open("./a")

	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	option := NewReadOption()
	v, err := db.Get(option, []byte("a"))

	if !bytes.Equal([]byte("b"), v) {
		t.Errorf("expected [%s], actual [%s]; error %s", "a", v, err)
	}
	option.Close()

	db.Close()
}
