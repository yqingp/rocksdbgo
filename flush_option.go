package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type FlushOption struct {
	Option *C.rocksdb_flushoptions_t
}

// extern rocksdb_flushoptions_t* rocksdb_flushoptions_create();
func NewFlushOption() *FlushOption {
	return &FlushOption{
		Option: C.rocksdb_flushoptions_create(),
	}
}

// extern void rocksdb_flushoptions_destroy(rocksdb_flushoptions_t*);
func (f *FlushOption) Close() {
	if f.Option != nil {
		C.rocksdb_flushoptions_destroy(f.Option)
	}
}

// extern void rocksdb_flushoptions_set_wait(rocksdb_flushoptions_t*, unsigned char);
func (f *FlushOption) SetWait(b bool) {
	t := 0
	if b {
		t = 1
	}

	C.rocksdb_flushoptions_set_wait(f.Option, C.uchar(t))
}
