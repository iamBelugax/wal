package main

import (
	"fmt"
	"os"

	"github.com/iamBelugax/wal/pkg/wal"
)

func main() {
	name, err := wal.MakeSegmentName(1234567890)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	fmt.Println("Segment Name :", name)

	id, err := wal.ExtractSegmentID(name)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	fmt.Println("Extracted ID :", id)

	if _, err = wal.Open(); err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
}
