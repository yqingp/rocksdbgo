package rocksdbgo

import (
	"bytes"
	. "github.com/yqingp/rocksdbgo"
	"os"
	"testing"
)

func TestPut(t *testing.T) {
	db, err := Open("./a", nil)

	if err != nil {
		t.Errorf("db open error: %s", err)
	}

	err = db.Put(nil, []byte("a"), []byte("b"))
	if err != nil {
		t.Errorf("put error: %s", err)
	}
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
	os.RemoveAll("./a")
}
