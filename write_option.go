package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type WriteOption struct {
	Option *C.rocksdb_writeoptions_t
}

// extern rocksdb_writeoptions_t* rocksdb_writeoptions_create();
func NewWriteOption() *WriteOption {
	opt := &WriteOption{}
	opt.Option = C.rocksdb_writeoptions_create()

	return opt
}

// extern void rocksdb_writeoptions_destroy(rocksdb_writeoptions_t*);
func (w *WriteOption) Close() {
	if w.Option != nil {
		C.rocksdb_writeoptions_destroy(w.Option)

	}
}

// extern void rocksdb_writeoptions_set_sync(rocksdb_writeoptions_t*, unsigned char);
func (w *WriteOption) SetSync(b bool) {
	t := 0
	if b {
		t = 1
	}

	C.rocksdb_writeoptions_set_sync(w.Option, C.uchar(t))
}

// extern void rocksdb_writeoptions_disable_WAL(rocksdb_writeoptions_t* opt, int disable);
func (w *WriteOption) DisableWal(disable int) {
	C.rocksdb_writeoptions_disable_WAL(w.Option, C.int(disable))
}
