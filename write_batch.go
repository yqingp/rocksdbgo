package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

import (
	// "errors"
	"unsafe"
)

type WriteBatch struct {
	writeBatch *C.rocksdb_writebatch_t
}

/*
extern rocksdb_writebatch_t* rocksdb_writebatch_create();
*/
func NewWriteBatch() *WriteBatch {
	return &WriteBatch{
		writeBatch: C.rocksdb_writebatch_create(),
	}
}

/*
extern void rocksdb_writebatch_put(rocksdb_writebatch_t*,const char* key, size_t klen,const char* val, size_t vlen);
*/
func (w *WriteBatch) Put(key []byte, value []byte) {

	k, v := C.CString(string(key)), C.CString(string(value))
	defer func() {
		C.free(unsafe.Pointer(k))
		C.free(unsafe.Pointer(v))
	}()

	C.rocksdb_writebatch_put(w.writeBatch, k, C.size_t(len(key)), v, C.size_t(len(value)))
}

/*
extern void rocksdb_writebatch_delete(
    rocksdb_writebatch_t*,
    const char* key, size_t klen);
*/
func (w *WriteBatch) Delete(key []byte) {
	k := C.CString(string(key))
	defer func() {
		C.free(unsafe.Pointer(k))
	}()

	C.rocksdb_writebatch_delete(w.writeBatch, k, C.size_t(len(key)))
}

/*
extern void rocksdb_writebatch_destroy(rocksdb_writebatch_t*);
*/
func (w *WriteBatch) Close() {
	C.rocksdb_writebatch_destroy(w.writeBatch)
}

/*
extern void rocksdb_writebatch_clear(rocksdb_writebatch_t*);
*/
func (w *WriteBatch) Clear() {
	C.rocksdb_writebatch_clear(w.writeBatch)
}

/*
extern int rocksdb_writebatch_count(rocksdb_writebatch_t*);
*/
func (w *WriteBatch) Count() int {
	return int(C.rocksdb_writebatch_count(w.writeBatch))
}
