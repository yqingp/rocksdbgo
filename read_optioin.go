package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type ReadOption struct {
	Option *C.rocksdb_readoptions_t
}

// extern rocksdb_readoptions_t* rocksdb_readoptions_create();
func NewReadOption() *ReadOption {
	opt := &ReadOption{}
	opt.Option = C.rocksdb_readoptions_create()

	return opt
}

// extern void rocksdb_readoptions_destroy(rocksdb_readoptions_t*);
func (r *ReadOption) Close() {
	C.rocksdb_readoptions_destroy(r.Option)
}

// extern void rocksdb_readoptions_set_verify_checksums(rocksdb_readoptions_t*, unsigned char);
func (r *ReadOption) SetVerifyChecksums(b bool) {
	t := 0
	if b {
		t = 1
	}
	C.rocksdb_readoptions_set_verify_checksums(r.Option, C.uchar(t))
}

// TODO:
// extern void rocksdb_readoptions_set_fill_cache(rocksdb_readoptions_t*, unsigned char);
// extern void rocksdb_readoptions_set_snapshot(rocksdb_readoptions_t*, const rocksdb_snapshot_t*);
// extern void rocksdb_readoptions_set_iterate_upper_bound(rocksdb_readoptions_t*, const char* key, size_t keylen);
// extern void rocksdb_readoptions_set_read_tier(rocksdb_readoptions_t*, int);
// extern void rocksdb_readoptions_set_tailing(rocksdb_readoptions_t*, unsigned char);
