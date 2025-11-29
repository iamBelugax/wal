package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/iamBelugax/wal/pkg/wal"
)

func main() {
	dir, _ := os.Getwd()
	json5MBBytes, _ := os.ReadFile("/Users/iamnilotpal/Documents/Coding Workspace/projects/wal/cmd/wal/5MB.json")
	json1MBBytes, _ := os.ReadFile("/Users/iamnilotpal/Documents/Coding Workspace/projects/wal/cmd/wal/1MB.json")

	wal, err := wal.Open(wal.WithDataDir(dir))
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	defer func() {
		wal.Close(context.Background())
	}()

	start := time.Now()

	for i := range 201 {
		lsn, err := wal.Append(context.Background(), json5MBBytes)
		if err != nil {
			fmt.Println("write :", err)
			os.Exit(1)
		}

		entry, err := wal.ReadAt(context.Background(), lsn)
		if err != nil {
			fmt.Println("read :", err)
			os.Exit(1)
		}
		fmt.Printf("LSN -> %d:%d \n", i+1, entry.Header.LSN)
	}

	for i := range 100 {
		lsn, err := wal.Append(context.Background(), json1MBBytes)
		if err != nil {
			fmt.Println("write :", err)
			os.Exit(1)
		}

		entry, err := wal.ReadAt(context.Background(), lsn)
		if err != nil {
			fmt.Println("read :", err)
			os.Exit(1)
		}
		fmt.Printf("LSN -> %d:%d \n", i+1, entry.Header.LSN)
	}

	fmt.Println("Time taken", time.Since(start).Seconds(), time.Since(start).Milliseconds())

	// lsn1, err := wal.Append(context.Background(), json5MBBytes)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	os.Exit(1)
	// }

	// lsn2, err := wal.Append(context.Background(), json5MBBytes)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	os.Exit(1)
	// }

	// lsn3, err := wal.Append(context.Background(), json5MBBytes)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	os.Exit(1)
	// }

	// entry1, _ := wal.ReadAt(context.Background(), lsn1)
	// fmt.Printf("Entry 1 Record: %+v \n", entry1.Record)
	// fmt.Printf("Entry 1 Header : %+v \n", entry1.Header)

	// entry2, _ := wal.ReadAt(context.Background(), lsn2)
	// fmt.Printf("Entry 2 Record: %+v \n", entry2.Record)
	// fmt.Printf("Entry 2 Header : %+v \n", entry2.Header)

	// entry3, _ := wal.ReadAt(context.Background(), lsn3)
	// fmt.Printf("Entry 3 Record: %+v \n", entry3.Record)
	// fmt.Printf("Entry 3 Header : %+v \n", entry3.Header)

	// name, err := wal.MakeSegmentName(1234567890)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("Segment Name :", name)

	// id, err := wal.ExtractSegmentID(name)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("Extracted ID :", id)

	// val := walpb.RecordHeader{}
	// val.SetMagic(999999999)
	// val.SetVersion(999999999)
	// val.SetChecksum(999999999)
	// val.SetLsn(9999999999999999999)
	// val.SetTimestamp(9999999999999999999)
	// val.SetRecordSize(9999999999999999999)
	// val.SetPreviousLsn(9999999999999999999)

	// data, _ := proto.Marshal(&val)
	// _ = data
	// fmt.Println("Set Max Value Disk Size :", len(data))

	// val = walpb.RecordHeader{}
	// val.SetLsn(0)
	// val.SetMagic(0)
	// val.SetVersion(0)
	// val.SetChecksum(0)
	// val.SetTimestamp(0)
	// val.SetRecordSize(0)
	// val.SetPreviousLsn(0)

	// data, _ = proto.Marshal(&val)
	// _ = data
	// fmt.Println("Set Min Value Disk Size :", len(data))

	// val = walpb.RecordHeader{}
	// data, _ = proto.Marshal(&val)
	// _ = data
	// fmt.Println("Unset Min Value Disk Size :", len(data))
}
