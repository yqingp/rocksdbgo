package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

import (
	"bytes"
	"unsafe"
)

type Iterator struct {
	Iterator  *C.rocksdb_iterator_t
	forward   bool
	start     string
	startChar *C.char
	startLen  C.size_t
	end       string
	endChar   *C.char
	endLen    C.size_t
	isFirst   bool
}

type Itval struct {
	K string
	V string
}

/*
extern rocksdb_iterator_t* rocksdb_create_iterator(
    rocksdb_t* db,
    const rocksdb_readoptions_t* options);
*/
func newIterator(db *DB, ro *ReadOption, forward bool, start string, end string) *Iterator {
	i := &Iterator{
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
	C.rocksdb_iter_seek(i.Iterator, i.startChar, i.startLen)
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

func (i *Iterator) Next() (Itval, bool) {
	if !i.valid() {
		return Itval{}, false
	}

	k := i.Key()
	v := i.Value()

	itval := Itval{K: k, V: v}

	if i.isFirst {
		if i.end != "" {
			if i.forward && bytes.Compare([]byte(k), []byte(i.end)) < 0 {
				i.Move()
				return itval, true
			}
			if !i.forward && bytes.Compare([]byte(k), []byte(i.end)) > 0 {
				i.Move()
				return itval, true
			}

			return Itval{}, false
		}

		i.Move()
		return itval, true
	}

	if i.end != "" {
		if i.forward && bytes.Compare([]byte(k), []byte(i.end)) < 0 {
			i.Move()
			return itval, true
		}
		if !i.forward && bytes.Compare([]byte(k), []byte(i.end)) > 0 {
			i.Move()
			return itval, true
		}
		return Itval{}, false
	}

	i.Move()
	return itval, true
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
	if k == nil {
		return ""
	}

	return string(C.GoBytes(unsafe.Pointer(k), C.int(l)))
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

	return string(C.GoBytes(unsafe.Pointer(v), C.int(l)))
}

/*
extern void rocksdb_iter_destroy(rocksdb_iterator_t*);
*/
func (i *Iterator) Close() {
	if i.Iterator != nil {
		C.rocksdb_iter_destroy(i.Iterator)
	}
	if i.startChar != nil {
		C.free(unsafe.Pointer(i.startChar))
	}

	if i.endChar != nil {
		C.free(unsafe.Pointer(i.endChar))
	}
}

/*
extern void rocksdb_iter_next(rocksdb_iterator_t*);
extern void rocksdb_iter_prev(rocksdb_iterator_t*);
*/
func (i *Iterator) Move() {
	if i.isFirst {
		i.isFirst = false
	}

	if i.forward {
		C.rocksdb_iter_next(i.Iterator)
	} else {
		C.rocksdb_iter_prev(i.Iterator)
	}
}

/*
extern void rocksdb_iter_destroy(rocksdb_iterator_t*);
extern void rocksdb_iter_get_error(const rocksdb_iterator_t*, char** errptr);
*/
