package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

type pcapHeader struct {
	MagicNumber       uint32
	MajorVersion      uint16
	MinorVersion      uint16
	TimezoneOffset    uint32
	TimestampAccuracy uint32
	SnapshotLength    uint32
	LinklayerType     uint32
}

type packetHeader struct {
	Timestamps        uint64
	Length            uint32
	UntruncatedLength uint32
}

type TCPHeader struct {
	SourcePort        uint16
	DestinationPort   uint16
	SequenceNumber    uint32
	AckNumber         uint32
	OffsetAndReserved uint8
	Flags             uint8
	WindowSize        uint16
	Checksum          uint16
	UrgentPointer     uint16
}

// Cannot be constructed like this
type packetPayloadHeaders struct {
	LinkLayerHeader [4]byte
	IP              [20]byte
	TCP             TCPHeader
}

const (
	expectedMagicNumber         = 0xa1b2c3d4
	pcapMajorVersion            = 2
	pcapMinorVersion            = 4
	expectedTimezoneOffset      = 0
	expectedTimestampAccuracy   = 0
	linklayerLoopbackValue      = 0
	packetPayloadHeadersInBytes = 44
	localhostPort               = 80
)

func main() {
	file, err := os.Open("./synflood.pcap")

	if err != nil {
		panic(err)
	}

	num, err := evaluatePCAPFile(file)
	if err != nil {
		fmt.Println(nil, err)
	} else {
		fmt.Println(*num, err)
	}
}

func evaluatePCAPFile(f *os.File) (*float64, error) {
	err := validatePCAPHeader(f)

	if err != nil {
		return nil, fmt.Errorf("validating PCAP Header: %w", err)
	}

	numPackets := 0.0

	for {
		var h packetHeader
		err = binary.Read(f, binary.LittleEndian, &h)

		if err != nil {
			if errors.Is(err, io.EOF) {
				return &numPackets, nil
			} else {
				return nil, fmt.Errorf("reading file into packet header: %w", err)
			}
		}

		// 4 bytes: Linklayer header
		// 20 bytes: IP Header
		// 20 bytes: TCP Header
		var payloadHeaders packetPayloadHeaders
		err = binary.Read(f, binary.BigEndian, &payloadHeaders)
		if err != nil {
			return nil, fmt.Errorf("validating packet payload headers: %w", err)
		}

		// Last 4 bits of first byte is the Internet Header Length.
		// IHL has the size of the IPv4 header: its 4 bits specify the number of
		// 32 bit words (e.g. 5 is 32 bits x 5 words = 160 bits = 20 bytes).
		if (payloadHeaders.IP[0] & 0b1111) != 5 {
			return nil, fmt.Errorf("expected IPv4 of length 20 bytes")
		}

		// Confirm source and destination port are localhost (80)
		if payloadHeaders.TCP.SourcePort != localhostPort {
			return nil, fmt.Errorf("source port\nactual: %d; expected: %d", payloadHeaders.TCP.SourcePort, localhostPort)
		}

		// Confirm source and destination port are localhost (80)
		if payloadHeaders.TCP.DestinationPort != localhostPort {
			return nil, fmt.Errorf("destination port\nactual: %d; expected: %d", payloadHeaders.TCP.DestinationPort, localhostPort)
		}

		payload := make([]byte, h.Length-packetPayloadHeadersInBytes)
		err = binary.Read(f, binary.BigEndian, &payload)
		if err != nil {
			return nil, fmt.Errorf("validating packet payload: %w", err)
		}
		numPackets += 1
	}
}

func validatePCAPHeader(f *os.File) error {
	var header pcapHeader
	err := binary.Read(f, binary.LittleEndian, &header)

	if err != nil {
		return fmt.Errorf("reading file into PCAP Header: %w", err)
	}

	if header.MagicNumber != expectedMagicNumber {
		return fmt.Errorf("actual PCAP Magic Number: %d, expected: %d", header.MagicNumber, expectedMagicNumber)
	}

	if header.MajorVersion != pcapMajorVersion {
		return fmt.Errorf("actual PCAP Major Version: %d, expected: %d", header.MajorVersion, pcapMajorVersion)
	}

	if header.MinorVersion != pcapMinorVersion {
		return fmt.Errorf("actual PCAP Minor Version: %d, expected: %d", header.MinorVersion, pcapMinorVersion)
	}

	if header.TimezoneOffset != expectedTimezoneOffset {
		return fmt.Errorf("actual PCAP Timezone Offset: %d, expected: %d", header.TimezoneOffset, expectedTimezoneOffset)
	}

	if header.TimestampAccuracy != expectedTimestampAccuracy {
		return fmt.Errorf("actual PCAP Timestamp Accuracy: %d, expected: %d", header.TimestampAccuracy, expectedTimestampAccuracy)
	}

	if header.LinklayerType != linklayerLoopbackValue {
		return fmt.Errorf("actual PCAP Linklayer Type value: %d, expected: %d", header.LinklayerType, linklayerLoopbackValue)
	}

	return nil
}
