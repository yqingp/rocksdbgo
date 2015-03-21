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

type ReadOption struct {
	Option *C.rocksdb_readoptions_t
}

type Option struct {
	Option *C.rocksdb_options_t
}

func NewWriteOption() *WriteOption {
	opt := &WriteOption{}
	opt.Option = C.rocksdb_writeoptions_create()

	return opt
}

func (w *WriteOption) Close() {
	C.rocksdb_writeoptions_destroy(w.Option)
}

func NewReadOption() *ReadOption {
	opt := &ReadOption{}
	opt.Option = C.rocksdb_readoptions_create()

	return opt
}

func (r *ReadOption) Close() {
	C.rocksdb_readoptions_destroy(r.Option)
}

func NewOption() *Option {
	opt := &Option{}
	opt.Option = C.rocksdb_options_create()

	return opt
}

func (o *Option) Close() {
	C.rocksdb_options_destroy(o.Option)
}
