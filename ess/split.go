package ess

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/trew/ess/shamir"
	"hash/adler32"
	"strings"
)

func Split(secret string, partCount int, threshold int) ([]string, error) {
	secretBytes := []byte(secret)
	secretBytes = append(secretBytes, getChecksumBytes(secretBytes)...)

	parts, err := shamir.Split(secretBytes, partCount, threshold)
	if err != nil {
		return nil, fmt.Errorf("failed to split secret: %v", err)
	}

	partStrings := make([]string, len(parts))
	for i, part := range parts {
		partStrings[i] = encode(part)
	}

	return partStrings, nil
}

func getChecksumBytes(bytes []byte) []byte {
	checksum := adler32.Checksum(bytes)
	checksumBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(checksumBytes, checksum)
	return checksumBytes
}

func encode(part []byte) string {
	return strings.ToUpper(hex.EncodeToString(part))
}
