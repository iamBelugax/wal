package main

import (
	"fmt"
	"os"

	walpb "github.com/iamBelugax/wal/internal/encoding/proto/__gen__"
	"github.com/iamBelugax/wal/pkg/wal"
	"google.golang.org/protobuf/proto"
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

	val := walpb.RecordHeader{}
	val.SetMagic(999999999)
	val.SetVersion(999999999)
	val.SetChecksum(999999999)
	val.SetLsn(9999999999999999999)
	val.SetTimestamp(9999999999999999999)
	val.SetPreviousLsn(9999999999999999999)

	data, _ := proto.Marshal(&val)
	_ = data
	fmt.Println("Set Max Value Disk Size :", len(data))

	val = walpb.RecordHeader{}
	val.SetLsn(0)
	val.SetMagic(0)
	val.SetVersion(0)
	val.SetChecksum(0)
	val.SetTimestamp(0)
	val.SetPreviousLsn(0)

	data, _ = proto.Marshal(&val)
	_ = data
	fmt.Println("Set Min Value Disk Size :", len(data))

	val = walpb.RecordHeader{}
	data, _ = proto.Marshal(&val)
	_ = data
	fmt.Println("Unset Min Value Disk Size :", len(data))
}
