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

/*
extern void rocksdb_iter_destroy(rocksdb_iterator_t*);
extern unsigned char rocksdb_iter_valid(const rocksdb_iterator_t*);
extern void rocksdb_iter_seek_to_first(rocksdb_iterator_t*);
extern void rocksdb_iter_seek_to_last(rocksdb_iterator_t*);
extern void rocksdb_iter_seek(rocksdb_iterator_t*, const char* k, size_t klen);
extern void rocksdb_iter_next(rocksdb_iterator_t*);
extern void rocksdb_iter_prev(rocksdb_iterator_t*);
extern const char* rocksdb_iter_key(const rocksdb_iterator_t*, size_t* klen);
extern const char* rocksdb_iter_value(const rocksdb_iterator_t*, size_t* vlen);
extern void rocksdb_iter_get_error(const rocksdb_iterator_t*, char** errptr);
*/
