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
	rocksdbOptions      *C.rocksdb_options_t
	wirteoptions        *C.rocksdb_writeoptions_t
	readoptions         *C.rocksdb_readoptions_t
}

func Open(dbpath string) (*DB, error) {
	db := &DB{}

	db.rocksdbOptions = C.rocksdb_options_create()
	db.wirteoptions = C.rocksdb_writeoptions_create()
	db.readoptions = C.rocksdb_readoptions_create()

	C.rocksdb_options_increase_parallelism(db.rocksdbOptions, C.int(runtime.NumCPU()))
	C.rocksdb_options_optimize_level_style_compaction(db.rocksdbOptions, 0)
	C.rocksdb_options_set_create_if_missing(db.rocksdbOptions, 1)

	dbpathCstring := C.CString(dbpath)

	var errInfo *C.char

	db.rocksdb = C.rocksdb_open(db.rocksdbOptions, dbpathCstring, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, errors.New(fmt.Sprintf("Store Rocksdb [Open] Error %s", er))
	}

	C.free(unsafe.Pointer(dbpathCstring))

	return db, nil
}

func (this *DB) Put(key []byte, value []byte) error {
	var errInfo *C.char

	k := C.CString(string(key))
	v := C.CString(string(value))

	C.rocksdb_put(this.rocksdb, this.wirteoptions, k, C.size_t(len(key)), v, C.size_t(len(value)), &errInfo)

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

	var l C.size_t
	var errInfo, value *C.char

	value = C.rocksdb_get(this.rocksdb, this.readoptions, k, C.size_t(len(key)), &l, &errInfo)

	if errInfo != nil {
		er := C.GoString(errInfo)
		return nil, errors.New(fmt.Sprintf("Store Rocksdb [Get] Error %s", er))
	}

	v := C.GoBytes(unsafe.Pointer(value), C.int(l))
	return v, nil
}

func (this *DB) Close() {
	C.rocksdb_writeoptions_destroy(this.wirteoptions)
	C.rocksdb_readoptions_destroy(this.readoptions)
	C.rocksdb_options_destroy(this.rocksdbOptions)
	C.rocksdb_close(this.rocksdb)
}

func (d *DB) String() string {
	return ""
}
