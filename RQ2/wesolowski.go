package main

import (
	"fmt"
	"time"

	vdf "github.com/harmony-one/vdf/src/vdf_go"
)

// Lets say that the POEM block time is 10 secs
// 1) 1 confirmations = 10 secs => 10,000 difficulty
// 2) 10 confirmations = 100 secs => 100,000 difficulty
// 3) 40 confirmations = 400 secs => 400,000 difficulty
// 4) 70 confirmations = 700 secs => 700,000 difficulty
// On the macbook, it seems that 10,000 difficulty takes around 10 secs to run,
// which would give the above assigned difficulty values to each of the
// confirmation values

func main() {
	input := []byte("Hello World!....................")
	vdf := vdf.New(10000, [32]byte(input))

	outputCh := vdf.GetOutputChannel()

	go vdf.Execute()

	select {
	case output := <-outputCh:
		start := time.Now()
		if vdf.Verify(output) {
			fmt.Println("vdf verified", "and took", time.Since(start), "to verify")
		}
	}
}
