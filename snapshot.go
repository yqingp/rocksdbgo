package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type Snapshot struct {
	snapshot *C.rocksdb_snapshot_t
	db       *DB
}

/*
extern void rocksdb_release_snapshot(
    rocksdb_t* db,
    const rocksdb_snapshot_t* snapshot);
*/
func (s *Snapshot) Close() {
	if s.snapshot != nil {
		C.rocksdb_release_snapshot(s.db.rocksdb, s.snapshot)
	}
}
