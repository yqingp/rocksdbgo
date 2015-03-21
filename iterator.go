package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type Iterator struct {
	Iterator *C.rocksdb_iterator_t
}

/*
extern rocksdb_iterator_t* rocksdb_create_iterator(
    rocksdb_t* db,
    const rocksdb_readoptions_t* options);
*/
func newIterator(db *DB, ro *ReadOption) *Iterator {
	return &Iterator{
		Iterator: C.rocksdb_create_iterator(db.Rocksdb, ro.Option),
	}
}

/*
extern void rocksdb_iter_destroy(rocksdb_iterator_t*);
*/
func (i *Iterator) Close() {
	C.rocksdb_iter_destroy(i.Iterator)
}
