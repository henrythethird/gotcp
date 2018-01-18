package gotcp

import "encoding/binary"

type packer struct{}

func newPacker() *packer {
	return &packer{}
}

func (p *packer) Pack(message string) []byte {
	len := calcPayloadLength(message)
	return append(len, message...)
}

func calcPayloadLength(m string) []byte {
	length := make([]byte, 8)
	binary.LittleEndian.PutUint64(length, uint64(len(m)))

	return length
}
