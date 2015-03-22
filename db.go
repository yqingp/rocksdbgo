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
	Rocksdb            *C.rocksdb_t
	DefaultWriteOption *WriteOption
	DefaultReadOption  *ReadOption
	Option             *Option
	FlashOption        *FlushOption
}

func Open(dbpath string, option *Option) (*DB, error) {
	db := &DB{}

	dbpathCstring := C.CString(dbpath)
	defer C.free(unsafe.Pointer(dbpathCstring))

	var errInfo *C.char

	if option == nil {
		db.Option = NewOption()
		db.Option.SetCreateIfMissing(true)
	}

	db.Rocksdb = C.rocksdb_open(db.Option.Option, dbpathCstring, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, errors.New(fmt.Sprintf("Store Rocksdb [Open] Error %s", er))
	}

	db.DefaultReadOption = NewReadOption()
	db.DefaultWriteOption = NewWriteOption()

	return db, nil
}

func (this *DB) Put(wo *WriteOption, key []byte, value []byte) error {
	var errInfo *C.char

	k, v := C.CString(string(key)), C.CString(string(value))
	defer func() {
		C.free(unsafe.Pointer(k))
		C.free(unsafe.Pointer(v))
	}()

	w := this.DefaultWriteOption.Option
	if wo != nil {
		w = wo.Option
	}

	C.rocksdb_put(this.Rocksdb, w, k, C.size_t(len(key)), v, C.size_t(len(value)), &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return errors.New(fmt.Sprintf("Store Rocksdb [Put] Error %s", er))
	}

	return nil
}

func (this *DB) Get(ro *ReadOption, key []byte) ([]byte, error) {
	var l C.size_t
	var errInfo, value *C.char

	k := C.CString(string(key))
	defer C.free(unsafe.Pointer(k))

	r := this.DefaultReadOption.Option
	if ro != nil {
		r = ro.Option
	}

	value = C.rocksdb_get(this.Rocksdb, r, k, C.size_t(len(key)), &l, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, errors.New(fmt.Sprintf("Store Rocksdb [Get] Error %s", er))
	}

	v := C.GoBytes(unsafe.Pointer(value), C.int(l))

	return v, nil
}

func (this *DB) Delete(wo *WriteOption, key []byte) error {
	var l C.size_t
	var errInfo *C.char

	k := C.CString(string(key))
	defer C.free(unsafe.Pointer(k))

	w := this.DefaultWriteOption.Option
	if wo != nil {
		w = wo.Option
	}

	C.rocksdb_delete(this.Rocksdb, w, k, l, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return errors.New(fmt.Sprintf("Store Rocksdb [Delete] Error %s", er))
	}

	return nil
}

func (this *DB) Close() {
	this.DefaultWriteOption.Close()
	this.DefaultReadOption.Close()
	this.Option.Close()
	// this.FlashOption.Close()
	C.rocksdb_close(this.Rocksdb)
}

func (d *DB) String() string {
	return ""
}

// extern void rocksdb_delete_file(rocksdb_t* db, const char* name);

// extern const rocksdb_livefiles_t* rocksdb_livefiles(rocksdb_t* db);

// extern void rocksdb_flush(rocksdb_t* db,const rocksdb_flushoptions_t* options,char** errptr);
func (d *DB) Flush() error {
	if d.FlashOption == nil {
		return errors.New("undefined FlashOption")
	}

	var errInfo *C.char

	C.rocksdb_flush(d.Rocksdb, d.FlashOption.Option, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return errors.New(er)
	}

	return nil
}

// extern void rocksdb_disable_file_deletions(rocksdb_t* db,char** errptr);

// extern void rocksdb_enable_file_deletions(rocksdb_t* db,unsigned char force, char** errptr);

func (d *DB) NewIterator(ro *ReadOption, forward bool, start string, end string) *Iterator {
	r := d.DefaultReadOption
	if ro != nil {
		r = ro
	}

	return newIterator(d, r, forward, start, end)
}
