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
	"runtime"
	"unsafe"
)

type DB struct {
	rocksdb             *C.rocksdb_t
	rocksdbBackupEngine *C.rocksdb_backup_engine_t
	writeOption         *WriteOption
	readOption          *ReadOption
	option              *Option
}

func Open(dbpath string) (*DB, error) {
	db := &DB{}

	db.option = NewOption()

	C.rocksdb_options_increase_parallelism(db.option.Option, C.int(runtime.NumCPU()))
	C.rocksdb_options_optimize_level_style_compaction(db.option.Option, 0)
	C.rocksdb_options_set_create_if_missing(db.option.Option, 1)

	dbpathCstring := C.CString(dbpath)

	var errInfo *C.char

	db.rocksdb = C.rocksdb_open(db.option.Option, dbpathCstring, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, errors.New(fmt.Sprintf("Store Rocksdb [Open] Error %s", er))
	}

	db.readOption = NewReadOption()
	db.writeOption = NewWriteOption()

	C.free(unsafe.Pointer(dbpathCstring))

	return db, nil
}

func (this *DB) Put(key []byte, value []byte) error {
	var errInfo *C.char

	k := C.CString(string(key))
	v := C.CString(string(value))

	C.rocksdb_put(this.rocksdb, this.writeOption.Option, k, C.size_t(len(key)), v, C.size_t(len(value)), &errInfo)

	C.free(unsafe.Pointer(k))
	C.free(unsafe.Pointer(v))

	if errInfo != nil {
		er := C.GoString(errInfo)
		return errors.New(fmt.Sprintf("Store Rocksdb [Put] Error %s", er))
	}

	return nil
}

func (this *DB) Get(key []byte) ([]byte, error) {
	k := C.CString(string(key))
	defer C.free(unsafe.Pointer(k))
	var l C.size_t
	var errInfo, value *C.char

	value = C.rocksdb_get(this.rocksdb, this.readOption.Option, k, C.size_t(len(key)), &l, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, errors.New(fmt.Sprintf("Store Rocksdb [Get] Error %s", er))
	}

	v := C.GoBytes(unsafe.Pointer(value), C.int(l))
	return v, nil
}

func (this *DB) Delete(key []byte) error {
	k := C.CString(string(key))
	defer C.free(unsafe.Pointer(k))

	var l C.size_t
	var errInfo *C.char
	C.rocksdb_delete(this.rocksdb, this.writeOption.Option, k, l, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return errors.New(fmt.Sprintf("Store Rocksdb [Delete] Error %s", er))
	}

	return nil
}

func (this *DB) Close() {
	this.writeOption.Close()
	this.readOption.Close()
	this.option.Close()
	C.rocksdb_close(this.rocksdb)
}

func (d *DB) String() string {
	return ""
}
