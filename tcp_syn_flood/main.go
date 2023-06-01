package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

type PcapHeader struct {
	MagicNumber  uint32
	MajorVersion uint16
	MinorVersion uint16
	Timezone     uint32
	Timestamp    uint32
	Snaplen      uint32
	Network      uint32
}

type PacketHeader struct {
	Timestamp         uint32
	TimestampOffset   uint32
	Length            uint32
	LengthUntruncated uint32
}

func main() {
	file, err := os.Open("./synflood.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalPackets, compromisedPackets int

	var header PcapHeader

	/*
		binary.Read(r, order, data) reads structured binary data from r into data.
	*/

	// You need to pass the pointer of header &header so that Read
	// writes to the exact memory address of header
	// Otherwise, passing `header` would mean passing a copy with
	// a different memory address and not writing the data to the
	// variable we were intending to
	binary.Read(file, binary.LittleEndian, &header)

	// Validate PCAP file format by checking the Magic Number (see man)

	// Validate PCAP version numbers (see man)

	// Parse each packet

	// Validate each packet header

	// Parse the packet payload: loop interface, IPv4, TCP (where SYN/ACK is) -> note that network protocols are big endian!

	// If a packet has been ACK'd, compromisedPackets++
	// Count each packet, then at the end get %

	fmt.Println(header)
}
