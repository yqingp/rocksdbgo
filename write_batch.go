package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

import (
	"errors"
)

type WriteBatch struct {
	writeBatch *C.rocksdb_writebatch_t
}

/*
extern rocksdb_writebatch_t* rocksdb_writebatch_create();
*/
func NewWriteBatch() {
	return &WriteBatch{
		writeBatch: C.rocksdb_writebatch_create(),
	}
}

/*
extern void rocksdb_writebatch_put(rocksdb_writebatch_t*,const char* key, size_t klen,const char* val, size_t vlen);
*/
func (w *WriteBatch) Put(key []byte, value []byte) {
	var errInfo *C.char

	k, v := C.CString(string(key)), C.CString(string(value))
	defer func() {
		C.free(unsafe.Pointer(k))
		C.free(unsafe.Pointer(v))
	}()

	C.rocksdb_writebatch_put(w.writeBatch, k, C.size_t(len(key)), v, C.size_t(len(value)))

	return nil
}
