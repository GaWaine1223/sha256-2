// Copyright 2018 Lothar . All rights reserved.
// https://github.com/Blockchain-CN

package sha256

import (
	"crypto/sha256"
	"strconv"
	"sync/atomic"
	"fmt"
)

// Size The size of a SHA256 checksum in bytes.
const Size = 32

var stop int32

// HashwithDifficulty ...
func HashwithDifficulty(data []byte, d int) (result [Size]byte, nonce int64) {
	for nonce = 1; ; nonce++ {
		if atomic.LoadInt32(&stop) == 1 {
			return result, 0
		}
		str := strconv.FormatInt(nonce, 10)
		b := append(data, []byte(str)...)
		result = sha256.Sum256(b)
		if difficulty(result, d) {
			return result, nonce
		}
	}
	return
}

func difficulty(hash [Size]byte, d int) bool {
	dn := d / 2
	sn := d % 2
	for i := 0; i < dn; i++ {
		if hash[i] != 0x00 {
			return false
		}
	}
	if sn != 0 {
		if hash[dn*2+1] > 0x0f {
			return false
		}
	}
	return true
}

// StopHash ...
func StopHash() bool {
	return atomic.CompareAndSwapInt32(&stop, 0, 1)
}

// Verification to test if the data's hash is equal to a string
func Verification(data []byte, hash string) bool {
	return hash == fmt.Sprint(sha256.Sum256(data))
}
