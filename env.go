package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type Env struct {
	env *C.rocksdb_env_t
}

// extern rocksdb_env_t* rocksdb_create_default_env();
func NewEnv() *Env {
	e := &Env{}
	e.env = C.rocksdb_create_default_env()

	return e
}

// extern void rocksdb_env_destroy(rocksdb_env_t*);
func (e *Env) Close() {
	if e.env != nil {
		C.rocksdb_env_destroy(e.env)
	}
}

// extern void rocksdb_env_set_background_threads(rocksdb_env_t* env, int n);
func (e *Env) SetBackGroundThreads(n int) {
	C.rocksdb_env_set_background_threads(e.env, C.int(n))
}

// rocksdb_env_set_high_priority_background_threads(rocksdb_env_t* env, int n);
func (e *Env) SetHighPriorityBackgroundThreads(n int) {
	C.rocksdb_env_set_high_priority_background_threads(e.env, C.int(n))
}
