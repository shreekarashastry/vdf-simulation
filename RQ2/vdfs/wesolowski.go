package vdfs

import (
	"fmt"
	"time"

	vdf "github.com/harmony-one/vdf/src/vdf_go"
	"github.com/shreekarashastry/vdf-simulation/RQ1/stats"
)

const difficulty = 10000
const num_iterations = 50

// Lets say that the POEM block time is 10 secs
// 1) 1 confirmations = 10 secs => 10,000 difficulty
// 2) 10 confirmations = 100 secs => 100,000 difficulty
// 3) 40 confirmations = 400 secs => 400,000 difficulty
// 4) 70 confirmations = 700 secs => 700,000 difficulty
// On the macbook, it seems that 10,000 difficulty takes around 10 secs to run,
// which would give the above assigned difficulty values to each of the
// confirmation values

func RunWesolowski() {
	input := []byte("Hello World!....................")

	computeTimes := []uint64{}
	verificationTimes := []uint64{}

	for i := 0; i < num_iterations; i++ {
		vdf := vdf.New(difficulty, [32]byte(input))
		outputCh := vdf.GetOutputChannel()

		startTime := time.Now()

		go vdf.Execute()

		select {
		case output := <-outputCh:
			computeTime := time.Since(startTime).Milliseconds()
			computeTimes = append(computeTimes, uint64(computeTime))

			startTime = time.Now()
			if vdf.Verify(output) {
				verificationTime := time.Since(startTime).Milliseconds()
				verificationTimes = append(verificationTimes, uint64(verificationTime))
			}
		}
	}
	fmt.Println("Algo", "Wesolowski", "Difficulty", difficulty, "Mean Compute Time", stats.Mean(computeTimes), "Median Compute Time", stats.Median(computeTimes))
	fmt.Println("Algo", "Wesolowski", "Difficulty", difficulty, "Mean Verfication Time", stats.Mean(verificationTimes), "Median Verification Time", stats.Median(verificationTimes))
}
