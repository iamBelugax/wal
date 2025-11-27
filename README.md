# Write-Ahead Log (WAL) in Go

This project implements a Write-Ahead Log (WAL) in Go, using Protocol Buffers
(default) as the encoding format for all log records. The goal of this WAL is to
provide a reliable foundation for systems that require strong durability
guarantees and predictable recovery behavior.

The WAL records every change in a strictly sequential order and ensures that
each entry is written to stable storage before it is considered committed. By
using Protocol Buffers, each record is encoded in a compact and schema-safe
format, which makes the log easy to extend and efficient to store.
