package gotcp

import (
	"encoding/binary"
)

// Width of a 64 bit integer
const prefixLength = 8

type parser struct {
}

func newParser() *parser {
	return &parser{}
}

// Parse parses a byte array of given length to extract length prefixed payloads
func (p *parser) Parse(buffer []byte, messageLength uint64) []string {
	var offset uint64
	var payloads []string

	for {
		payloadLength := extractPayloadLength(buffer)
		start := offset + prefixLength
		end := start + payloadLength

		payload := string(buffer[start:end])
		payloads = append(payloads, payload)

		if end >= messageLength {
			return payloads
		}

		offset = end
	}
}

func extractPayloadLength(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}
