package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type Option struct {
	Option *C.rocksdb_options_t
}

// extern rocksdb_options_t* rocksdb_options_create();
func NewOption() *Option {
	opt := &Option{}
	opt.Option = C.rocksdb_options_create()

	return opt
}

// extern void rocksdb_options_increase_parallelism(rocksdb_options_t* opt, int total_threads);
func (o *Option) IncreaseParallelism(n int) {
	C.rocksdb_options_increase_parallelism(o.Option, C.int(n))
}

// extern void rocksdb_options_destroy(rocksdb_options_t*);
func (o *Option) Close() {
	C.rocksdb_options_destroy(o.Option)
}

// extern void rocksdb_options_set_create_if_missing(rocksdb_options_t*, unsigned char);
func (o *Option) SetCreateIfMissing(b bool) {
	t := 0
	if b {
		t = 1
	}

	C.rocksdb_options_set_create_if_missing(o.Option, C.uchar(t))
}
