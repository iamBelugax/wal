package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/iamBelugax/wal/pkg/wal"
)

// 50,000 records with 730 Bytes each (36.5 MB) with 1KB Page Size -  51.2 MB (Overhead 47.1%)
// 50,000 records with 730 Bytes each (36.5 MB) with 2KB Page Size - 102.4 MB (Overhead of fucking 194.83%)
// 50,000 records with 730 Bytes each (36.5 MB) with 4KB Page Size - 204.8 MB (Overhead of fucking 488.37%)
//
// The average payload size should be approximately equal to the Page Size.

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

	for i := range 50000 {
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
}
