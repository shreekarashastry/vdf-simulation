package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

const (
	max_time = 5 * time.Second
)

func main() {
	input := "hello world"
	var prevSum, sum [32]byte
	prevSum = sha256.Sum256([]byte(input))

	timer := time.NewTimer(max_time)
	// Run sha vdf for max_time
	for {
		select {
		case <-timer.C:
			fmt.Printf("Last sum %x", sum)
			return
		default:
			sum = sha256.Sum256(prevSum[:])
			prevSum = sum
		}
	}
}
