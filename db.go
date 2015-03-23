package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type DB struct {
	rocksdb            *C.rocksdb_t
	defaultWriteOption *WriteOption
	defaultReadOption  *ReadOption
	option             *Option
	flashOption        *FlushOption
}

func Open(dbpath string, option *Option) (*DB, error) {
	db := &DB{}

	dbpathCstring := C.CString(dbpath)
	defer C.free(unsafe.Pointer(dbpathCstring))

	var errInfo *C.char

	if option == nil {
		db.option = NewOption()
		db.option.SetCreateIfMissing(true)
	}

	db.rocksdb = C.rocksdb_open(db.option.option, dbpathCstring, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, fmt.Errorf("Rocksdb [Error] %s", er)
	}

	db.defaultReadOption = NewReadOption()
	db.defaultWriteOption = NewWriteOption()
	db.flashOption = NewFlushOption()

	return db, nil
}

func (this *DB) Put(wo *WriteOption, key []byte, value []byte) error {
	var errInfo *C.char

	k, v := C.CString(string(key)), C.CString(string(value))
	defer func() {
		C.free(unsafe.Pointer(k))
		C.free(unsafe.Pointer(v))
	}()

	w := this.defaultWriteOption.writeOption
	if wo != nil {
		w = wo.writeOption
	}

	C.rocksdb_put(this.rocksdb, w, k, C.size_t(len(key)), v, C.size_t(len(value)), &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return fmt.Errorf("Rocksdb [Error] %s", er)
	}

	return nil
}

func (this *DB) Get(ro *ReadOption, key []byte) ([]byte, error) {
	var l C.size_t
	var errInfo, value *C.char

	k := C.CString(string(key))
	defer C.free(unsafe.Pointer(k))

	r := this.defaultReadOption.readOption
	if ro != nil {
		r = ro.readOption
	}

	value = C.rocksdb_get(this.rocksdb, r, k, C.size_t(len(key)), &l, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, fmt.Errorf("Rocksdb [Error] %s", er)
	}

	v := C.GoBytes(unsafe.Pointer(value), C.int(l))

	return v, nil
}

func (this *DB) Delete(wo *WriteOption, key []byte) error {
	var l C.size_t
	var errInfo *C.char

	k := C.CString(string(key))
	defer C.free(unsafe.Pointer(k))

	w := this.defaultWriteOption.writeOption
	if wo != nil {
		w = wo.writeOption
	}

	C.rocksdb_delete(this.rocksdb, w, k, l, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return fmt.Errorf("Rocksdb [Error] %s", er)
	}

	return nil
}

func (this *DB) Close() {
	this.defaultReadOption.Close()
	this.defaultWriteOption.Close()
	this.option.Close()
	this.flashOption.Close()
	C.rocksdb_close(this.rocksdb)
}

func (d *DB) String() string {
	return ""
}

// extern void rocksdb_flush(rocksdb_t* db,const rocksdb_flushoptions_t* options,char** errptr);
func (d *DB) Flush() error {
	if d.flashOption == nil {
		return fmt.Errorf("Rocksdb [Error] FlushOption Nil")
	}

	var errInfo *C.char

	C.rocksdb_flush(d.rocksdb, d.flashOption.flushOption, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return errors.New(er)
	}

	return nil
}

func (d *DB) NewIterator(ro *ReadOption, forward bool, start string, end string) *Iterator {
	r := d.defaultReadOption
	if ro != nil {
		r = ro
	}

	return newIterator(d, r, forward, start, end)
}

/*
extern void rocksdb_write(
    rocksdb_t* db,
    const rocksdb_writeoptions_t* options,
    rocksdb_writebatch_t* batch,
    char** errptr);
*/
func (d *DB) Write(wo *WriteOption, wb *WriteBatch) error {
	var errInfo *C.char

	w := d.defaultWriteOption.writeOption
	if wo != nil {
		w = wo.writeOption
	}

	C.rocksdb_write(d.rocksdb, w, wb.writeBatch, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return fmt.Errorf("Rocksdb [Error] %s", er)
	}

	return nil
}

/*
extern const rocksdb_snapshot_t* rocksdb_create_snapshot(
    rocksdb_t* db);
*/
func (d *DB) CreateSnapshot() *Snapshot {
	return &Snapshot{
		snapshot: C.rocksdb_create_snapshot(d.rocksdb),
	}
}
