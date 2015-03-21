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

	db.Put([]byte("a"), []byte("b"))
	db.Close()
}

func TestGet(t *testing.T) {
	db, err := Open("./a")

	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	v, err := db.Get([]byte("a"))

	if !bytes.Equal([]byte("b"), v) {
		t.Errorf("expected [%s], actual [%s]; error %s", "a", v, err)
	}

	// v1, err := db.Get([]byte("b"))

	// if !strings.EqualFold("", v1) {
	// 	t.Errorf("expected [%s], actual [%s]", "", v1)
	// }

	db.Close()
}
