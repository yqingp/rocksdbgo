package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

import (
	"bytes"
)

type Iterator struct {
	Iterator  *C.rocksdb_iterator_t
	forward   bool
	start     string
	startChar *C.char
	startLen  *C.size_t
	end       string
	endChar   *C.char
	endLen    *C.size_t
	isFirst   bool
}

/*
extern rocksdb_iterator_t* rocksdb_create_iterator(
    rocksdb_t* db,
    const rocksdb_readoptions_t* options);
*/
func newIterator(db *DB, ro *ReadOption, forward bool, start string, end string) *Iterator {
	i := Iterator{
		Iterator: C.rocksdb_create_iterator(db.Rocksdb, ro.Option),
	}

	i.isFirst = true

	if forward {
		i.Forward()
	} else {
		i.Backward()
	}

	return i
}

/*
extern void rocksdb_iter_seek_to_first(rocksdb_iterator_t*);
*/
func (i *Iterator) Forward() {
	i.forward = true
	C.rocksdb_iter_seek_to_first(i.Iterator)
}

/*
extern void rocksdb_iter_seek_to_last(rocksdb_iterator_t*);
*/
func (i *Iterator) Backward() {
	i.forward = false
	C.rocksdb_iter_seek_to_last(i.Iterator)
}

/*
extern void rocksdb_iter_seek(rocksdb_iterator_t*, const char* k, size_t klen);
*/
func (i *Iterator) seek() {
	C.rocksdb_iter_seek(i.Iterator, i.startChar)
}

func (i *Iterator) Start(start string) {
	if start != "" {
		i.start = start
		i.startChar = C.CString(start)
		i.startLen = C.size_t(len([]byte(start)))
		i.seek()
	}
}

func (i *Iterator) End(end string) {
	if end != "" {
		i.end = end
		i.endChar = C.CString(end)
		i.endLen = C.size_t(len([]byte(end)))
	}
}

func (i *Iterator) ResetDirection() {
	if i.forward {
		i.Forward()
	} else {
		i.Backward()
	}
}

func (i *Iterator) Next() (string, bool) {
	if !i.valid() {
		return "", false
	}

	if i.isFirst {
		v := i.Value()

		if i.end != "" {
			if i.forward && bytes.Compare([]byte(v), []byte(i.end)) < 0 {
				return v, true
			}
			if !i.forward && bytes.Compare([]byte(v), []byte(i.end)) > 0 {
				return v, true
			}

			return "", false
		}

		return v, true
	}

}

/*
extern unsigned char rocksdb_iter_valid(const rocksdb_iterator_t*);
*/
func (i *Iterator) valid() bool {
	return C.rocksdb_iter_valid(i.Iterator) == C.uchar(1)
}

/*
extern const char* rocksdb_iter_key(const rocksdb_iterator_t*, size_t* klen);
*/
func (i *Iterator) Key() string {
	var l C.size_t
	k := C.rocksdb_iter_key(i.Iterator, &l)
	if kdata == nil {
		return ""
	}

	return C.GoString(k)
}

/*
extern const char* rocksdb_iter_value(const rocksdb_iterator_t*, size_t* vlen);
*/
func (i *Iterator) Value() string {
	var l C.size_t
	v := C.rocksdb_iter_value(i.Iterator, &l)
	if v == nil {
		return ""
	}

	return C.GoString(v)
}

/*
extern void rocksdb_iter_destroy(rocksdb_iterator_t*);
*/
func (i *Iterator) Close() {
	C.rocksdb_iter_destroy(i.Iterator)
}

/*
 */

/*
extern void rocksdb_iter_destroy(rocksdb_iterator_t*);
extern void rocksdb_iter_next(rocksdb_iterator_t*);
extern void rocksdb_iter_prev(rocksdb_iterator_t*);
extern void rocksdb_iter_get_error(const rocksdb_iterator_t*, char** errptr);
*/
