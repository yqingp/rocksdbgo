package rocksdbgo

import (
	. "github.com/yqingp/rocksdbgo"
	"os"
	"testing"
)

func TestWriteBatch(t *testing.T) {
	db, err := Open("./a", nil)
	if err != nil {
		t.Errorf("db open error: %s", err)
	}
	wb := NewWriteBatch()
	wb.Put([]byte("a"), []byte("v"))
	wb.Delete([]byte("a"))

	err = db.Write(nil, wb)

	if err != nil {
		t.Errorf("error: %s", err)
	}

	wb.Close()
	db.Close()
	os.RemoveAll("./a")
}
