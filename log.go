package rocksdbgo

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "Rocksdb: ", log.Lshortfile|log.Ldate|log.Ltime)
