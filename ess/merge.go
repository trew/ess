package ess

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/trew/ess/shamir"
	"hash/adler32"
	"strings"
)

func Merge(parts []string) (string, error) {
	partsBytes := make([][]byte, len(parts))
	for i, arg := range parts {
		decoded, err := decode(arg)
		if err != nil {
			return "", fmt.Errorf("failed to decode secret: %v", err)
		}
		partsBytes[i] = decoded
	}
	merged, err := shamir.Combine(partsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to merge secret: %v", err)
	}

	// ensure checksum is extractable
	if len(merged) < 4 {
		return "", fmt.Errorf("failed to merge secret: parts are too short")
	}

	secret := merged[0 : len(merged)-4]
	checksumBytes := merged[len(merged)-4:]
	expectedChecksum := binary.LittleEndian.Uint32(checksumBytes)
	checksum := adler32.Checksum(secret)

	if expectedChecksum != checksum {
		return "", fmt.Errorf("failed to merge secret with the given parts")
	}

	return string(secret), nil
}

func decode(part string) ([]byte, error) {
	return hex.DecodeString(strings.ToLower(part))
}
