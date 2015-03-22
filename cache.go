package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type Cache struct {
	cache *C.rocksdb_cache_t
}

// extern rocksdb_cache_t* rocksdb_cache_create_lru(size_t capacity);
func NewCache(capacity int) *Cache {
	return &Cache{
		cache: C.rocksdb_cache_create_lru(C.size_t(capacity)),
	}
}

// extern void rocksdb_cache_destroy(rocksdb_cache_t* cache);
func (c *Cache) Close() {
	if c.cache != nil {
		C.rocksdb_cache_destroy(c.cache)
	}
}
