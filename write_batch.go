package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type WriteBatch struct {
	WriteBatch *C.rocksdb_writebatch_t
}
