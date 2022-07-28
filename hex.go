package main

import "time"

func bytesToInt16(bytes []byte) uint16 {
	data := uint16(bytes[0])<<8 + uint16(bytes[1])
	return data
}

func bytesToInt32(bytes []byte) uint32 {
	return uint32(bytes[0])<<24 +
		uint32(bytes[1])<<16 +
		uint32(bytes[2])<<8 +
		uint32(bytes[3])
}

func bytesToInt64(bytes []byte) uint64 {
	return uint64(bytes[0])<<56 +
		uint64(bytes[1])<<48 +
		uint64(bytes[2])<<40 +
		uint64(bytes[3])<<32 +
		uint64(bytes[4])<<24 +
		uint64(bytes[5])<<16 +
		uint64(bytes[6])<<8 +
		uint64(bytes[7])
}

func hexToDate(bytes []byte) string {
	seconds := bytesToInt32(bytes)
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	now := start.Add(time.Duration(seconds) * time.Second)
	return now.Format("2006-01-02 15:04:05")
}
