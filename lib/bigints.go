package lib

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

// take 2 slices of bytes (big ints), combine it, and return hex string
func EncodeBigInts(a, b []byte) string {
	bytes := append(a, b...)
	return fmt.Sprintf("%x", bytes)
}

// decode hex string. and divide by 2 big ints
func RestoreBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}
	firstHalfBytes := bytes[:len(bytes)/2]
	sencodHalfBytes := bytes[len(bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalfBytes)
	bigB.SetBytes(sencodHalfBytes)
	return &bigA, &bigB, nil
}
