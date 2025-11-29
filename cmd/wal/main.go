package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/iamBelugax/wal/pkg/wal"
)

// 50,000 records at 730 bytes each (36.5 MB) with a 1KB page size = 51.2 MB (Overhead: 47.1%)
// 50,000 records at 730 bytes each (36.5 MB) with a 2KB page size = 102.4 MB (Overhead: fucking 194.83%)
// 50,000 records at 730 bytes each (36.5 MB) with a 4KB page size = 204.8 MB (Overhead: fucking 488.37%)
//
// I was just messing around with testing how long it takes without calling Sync
// for each write but damn, it’s so good. For 100K records, each being 4KB, it took 1.04760275 seconds.

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed to get PWD :", err)
	}

	jsonBytes, err := os.ReadFile(filepath.Join(pwd, "cmd", "wal", "data.json"))
	if err != nil {
		log.Fatalln("Failed to read data file :", err)
	}

	wal, err := wal.Open(wal.WithDataDir(filepath.Join(pwd, "cmd", "wal")))
	if err != nil {
		log.Fatalln("Failed to open wal :", err)
	}

	defer func() {
		wal.Close(context.Background())
	}()

	start := time.Now()

	for i := range 100000 {
		lsn, err := wal.Append(context.Background(), jsonBytes)
		if err != nil {
			log.Fatalln("Failed to append :", err)
		}

		entry, err := wal.ReadAt(context.Background(), lsn)
		if err != nil {
			log.Fatalln("Failed to read :", err)
		}

		fmt.Printf("ID - %d, LSN - %d \n", i+1, entry.Header.LSN)
	}

	fmt.Println("\n Time Taken", time.Since(start).Nanoseconds())
}
