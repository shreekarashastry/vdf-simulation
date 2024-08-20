package sha256

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/shreekarashastry/vdf-simulation/RQ1/stats"
)

const (
	max_time       = 5 * time.Second
	max_iterations = 50
)

func RunShaSim() {
	input := "hello world"
	sum := sha256.Sum256([]byte(input))

	shaHashesEachRound := []uint64{}
	for i := 0; i < max_iterations; i++ {
		shaHashes := 0

		var done bool
		timer := time.NewTimer(max_time)
		// Run sha vdf for max_time
		for !done {
			select {
			case <-timer.C:
				done = true
			default:
				sha256.Sum256(sum[:])
				shaHashes++
			}
		}
		shaHashesEachRound = append(shaHashesEachRound, uint64(shaHashes))
	}

	fmt.Println("Algo", "Sha256", "Duration", max_time.Seconds(), "s", "Mean", stats.Mean(shaHashesEachRound), "Median", stats.Median(shaHashesEachRound))
}
