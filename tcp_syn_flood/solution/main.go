package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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
	NetworkType  uint32
}

type PacketHeader struct {
	Timestamp         uint32
	TimestampOffset   uint32
	Length            uint32
	LengthUntruncated uint32
}

type PacketData struct {
	LinkHeader []byte
	IP         IPHeader
	TCP        TCPHeader
	Payload    []byte
}

type IPHeader struct {
	InternetHeaderLength uint8 // 4 bits representing number of 32-but words in header
	HeaderLenBytes       uint16
	TotalLenBytes        uint16
	Protocol             uint8
}

type TCPHeader struct {
	SourcePort      uint16
	DestinationPort uint16
	HeaderLenBytes  uint8
	SequenceNumber  uint32
	AckNumber       uint32
	AckSet          bool
	SynSet          bool
}

var (
	pcapMajorVersion                         uint16 = 2
	pcapMinorVersion                         uint16 = 4
	differentByteOrder                       uint32 = 0xd4c3b2a1
	sameByteOrderAndDifferentTimestamps      uint32 = 0xa1b23c4d
	differentByteOrderAndDifferentTimestamps uint32 = 0x4d3cb2a1
)

const (
	SYN = 0x02
	ACK = 0x10
)

func main() {
	file, err := os.Open("./synflood.pcap")
	if err != nil {
		log.Fatal(err)
	}

	result, err := evaluateCompromisedFile(file)
	fmt.Println(*result, err)
	file.Close()
}

func evaluateCompromisedFile(f *os.File) (*float64, error) {
	var header PcapHeader

	/*
		binary.Read(r, order, data) reads structured binary data from r into data.
		r is a Reader interface that behaves like a stream of data.
		You need to pass the pointer of header &header so that Read
		writes to the exact memory address of header.
		Otherwise, passing `header` would mean passing a copy with
		a different memory address and not writing the data to the
		variable we were intending to.
	*/

	binary.Read(f, binary.LittleEndian, &header)

	err := validatePcapHeader(header)

	if err != nil {
		return nil, fmt.Errorf("validatingPcapHeader: %w", err)
	}

	return evaluatePackets(f)
}

func evaluatePackets(f *os.File) (*float64, error) {
	syns := 0.0
	acks := 0.0

	// Define an anonymous function
	results := func(err error) (*float64, error) {
		if errors.Is(err, io.EOF) {
			fmt.Printf("%v/%v = %v ack/syn\n", acks, syns, acks/syns)
			result := acks / syns
			return &result, nil
		}
		return nil, fmt.Errorf("error is not EOF: %w", err)
	}

	for {
		// Parse each packet
		p := PacketHeader{}
		err := binary.Read(f, binary.LittleEndian, &p) // why does this return 0?
		if err != nil {
			// EOF is the error we want and that will stop the loop
			return results(err)
		}

		// Make a []byte of p.Length to read exactly the number of bytes as in the packet's payload
		data := make([]byte, p.Length)

		// Parse the packet payload: note that network protocols are big endian
		binary.Read(f, binary.BigEndian, &data)
		// loop interface, IPv4, TCP (where SYN/ACK is) -> n
		tcpData := data[24:]
		sourcePort := binary.BigEndian.Uint16(tcpData[0:2])

		if sourcePort == 80 && (tcpData[13]&ACK == ACK) {
			acks++
		} else if sourcePort != 80 && (tcpData[13]&SYN == SYN) {
			syns++
		}
	}
}

func validatePcapHeader(h PcapHeader) error {
	// Validate PCAP file format by checking the Magic Number (see man)
	if h.MagicNumber == sameByteOrderAndDifferentTimestamps {
		// Further work: consider this case and treat it as microseconds
		return errors.New("packet timestamps in seconds and nanoseconds rather than seconds and microseconds")
	} else if h.MagicNumber == differentByteOrder {
		// Further work: swap the byte ordering of the file contents instead of returning an error
		return errors.New("this computer has a different byte ordering to the file's")
	} else if h.MagicNumber == differentByteOrderAndDifferentTimestamps {
		// Further work: swap the byte ordering of the file contents instead of returning an error and consider packets and second and nanoseconds
		return errors.New("this computer has a different byte ordering to the file's and packet timestamps are in seconds and nanoseconds")
	}

	// Validate PCAP version numbers
	if h.MajorVersion != pcapMajorVersion {
		// Further work: some way to handle the right combination of possible major and minor version numbers?
		// Caveat: 2.4 has been in use since 1998, so it shouldn't be a problem?
		return fmt.Errorf("this file does not use the current PCAP major version, %d", h.MajorVersion)
	}

	if h.MinorVersion != pcapMinorVersion {
		// Further work: some way to handle the right combination of possible major and minor version numbers?
		// Caveat: 2.4 has been in use since 1998, so it shouldn't be a problem?
		return errors.New("this file does not use the current PCAP minor version")
	}

	if h.NetworkType != 0 {
		return errors.New("should be using the loopback interface")
	}

	return nil
}
