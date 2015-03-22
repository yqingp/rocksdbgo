package rocksdbgo

/*
#cgo LDFLAGS: -lrocksdb
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"

type Option struct {
	option *C.rocksdb_options_t
}

// extern rocksdb_options_t* rocksdb_options_create();
func NewOption() *Option {
	opt := &Option{}
	opt.option = C.rocksdb_options_create()

	return opt
}

// extern void rocksdb_options_increase_parallelism(rocksdb_options_t* opt, int total_threads);
func (o *Option) IncreaseParallelism(n int) {
	C.rocksdb_options_increase_parallelism(o.option, C.int(n))
}

// extern void rocksdb_options_destroy(rocksdb_options_t*);
func (o *Option) Close() {
	if o.option != nil {
		C.rocksdb_options_destroy(o.option)
	}
}

/*
extern void rocksdb_options_set_compression(rocksdb_options_t*, int);
enum {
  rocksdb_no_compression = 0,
  rocksdb_snappy_compression = 1,
  rocksdb_zlib_compression = 2,
  rocksdb_bz2_compression = 3,
  rocksdb_lz4_compression = 4,
  rocksdb_lz4hc_compression = 5
};
*/
func (o *Option) SetCompression(t int) {
	C.rocksdb_options_set_compression(o.option, C.int(t))
}

/*
enum {
  rocksdb_level_compaction = 0,
  rocksdb_universal_compaction = 1,
  rocksdb_fifo_compaction = 2
};
extern void rocksdb_options_set_compaction_style(rocksdb_options_t*, int);
*/
func (o *Option) SetCompactionStyle(t int) {
	C.rocksdb_options_set_compaction_style(o.option, C.int(t))
}

/*
extern void rocksdb_options_optimize_for_point_lookup(
    rocksdb_options_t* opt, uint64_t block_cache_size_mb);
*/
func (o *Option) OptimizeForPointLookup(size uint64) {
	C.rocksdb_options_optimize_for_point_lookup(o.option, C.uint64_t(size))
}

/*
extern void rocksdb_options_optimize_level_style_compaction(
    rocksdb_options_t* opt, uint64_t memtable_memory_budget);
*/
func (o *Option) OptimizeLevelStyleCompaction(size uint64) {
	C.rocksdb_options_optimize_level_style_compaction(o.option, C.uint64_t(size))
}

/*
extern void rocksdb_options_optimize_universal_style_compaction(
    rocksdb_options_t* opt, uint64_t memtable_memory_budget);
*/
func (o *Option) OptimizeUniversalStyleCompaction(size uint64) {
	C.rocksdb_options_optimize_universal_style_compaction(o.option, C.uint64_t(size))
}

/*
extern void rocksdb_options_set_compression_per_level(
  rocksdb_options_t* opt,
  int* level_values,
  size_t num_levels);
*/
func (o *Option) SetCompressionPerLevel(levelValues int, numLevels int) {
	l := C.int(levelValues)
	C.rocksdb_options_set_compression_per_level(o.option, &l, C.size_t(numLevels))
}

/*
extern void rocksdb_options_set_create_if_missing(rocksdb_options_t*, unsigned char);
*/
func (o *Option) SetCreateIfMissing(b bool) {
	t := 0
	if b {
		t = 1
	}

	C.rocksdb_options_set_create_if_missing(o.option, C.uchar(t))
}

/*
extern void rocksdb_options_set_create_missing_column_families(
    rocksdb_options_t*, unsigned char);
*/
func (o *Option) SetCreateMissingColumnFamilies(b bool) {
	t := C.uchar(0)
	if b {
		t = C.uchar(1)
	}

	C.rocksdb_options_set_create_missing_column_families(o.option, t)
}

/*
extern void rocksdb_options_set_error_if_exists(
    rocksdb_options_t*, unsigned char);
*/
func (o *Option) SetErrorIfExists(b bool) {
	t := C.uchar(0)
	if b {
		t = C.uchar(1)
	}

	C.rocksdb_options_set_error_if_exists(o.option, t)
}

/*
extern void rocksdb_options_set_paranoid_checks(
    rocksdb_options_t*, unsigned char);
*/
func (o *Option) SetParanoidChecks(b bool) {
	t := C.uchar(0)
	if b {
		t = C.uchar(1)
	}

	C.rocksdb_options_set_paranoid_checks(o.option, t)
}

/*
extern void rocksdb_options_set_env(rocksdb_options_t*, rocksdb_env_t*);
*/
func (o *Option) SetEnv(env *Env) {
	C.rocksdb_options_set_env(o.option, env.env)
}

/*
extern void rocksdb_options_set_write_buffer_size(rocksdb_options_t*, size_t);
*/
func (o *Option) SetWriteBufferSize(size uint32) {
	C.rocksdb_options_set_write_buffer_size(o.option, C.size_t(size))
}

/*
extern void rocksdb_options_set_max_open_files(rocksdb_options_t*, int);
*/
func (o *Option) SetMaxOpenFiles(size int) {
	C.rocksdb_options_set_max_open_files(o.option, C.int(size))
}

/*
extern void rocksdb_options_set_max_total_wal_size(rocksdb_options_t* opt, uint64_t n);
*/
func (o *Option) SetMaxTotalWalSize(size uint64) {
	C.rocksdb_options_set_max_total_wal_size(o.option, C.uint64_t(size))
}

/*
extern void rocksdb_options_set_num_levels(rocksdb_options_t*, int);
*/
func (o *Option) SetNumLevels(level int) {
	C.rocksdb_options_set_num_levels(o.option, C.int(level))
}
