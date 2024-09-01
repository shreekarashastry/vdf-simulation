package vdfs

import (
	"fmt"
	"math/big"
	"time"

	"github.com/shreekarashastry/vdf-simulation/RQ1/stats"
	vdf "github.com/shreekarashastry/vdf/go_src/candidates/simple_vdf"
)

const p_difficulty = 30000000 // This value roughly corresponds to 10 seconds

// 1) 1 confirmations = 10 secs = 30,000,000
// 2) 10 confirmations = 100 secs = 300,000,000
// 3) 40 confirmations = 400 secs = 1,200,000,000
// 4) 70 confirmations = 700 secs = 2,100,000,000

func RunPietrzaks() {

	input, _ := new(big.Int).SetString("346666666627277", 0)
	var p, _ = new(big.Int).SetString("126493185890016866190387990037436305339", 0)
	var q, _ = new(big.Int).SetString("237515677732435432578220196406645605033", 0)

	//p*q=N
	var N = vdf.Mul(p, q)

	computeTimes := []uint64{}
	verificationTimes := []uint64{}

	for i := 0; i < num_iterations; i++ {
		start := time.Now()
		starting_value := vdf.ComputeVDF(p_difficulty, input, N)

		computeTime := time.Since(start).Milliseconds()
		computeTimes = append(computeTimes, uint64(computeTime))

		start = time.Now()
		verif := true
		for _, statement := range vdf.VerifyVDF(starting_value, N) {
			if !statement {
				verif = false
			}
		}

		if !verif {
			fmt.Println("verification failed")
			return
		} else {
			verificationTime := time.Since(start).Milliseconds()
			verificationTimes = append(verificationTimes, uint64(verificationTime))
		}
	}

	fmt.Println("Algo", "Pietrzaks", "Difficulty", difficulty, "Mean Compute Time", stats.Mean(computeTimes), "Median Compute Time", stats.Median(computeTimes))
	fmt.Println("Algo", "Pietrzaks", "Difficulty", difficulty, "Mean Verfication Time", stats.Mean(verificationTimes), "Median Verification Time", stats.Median(verificationTimes))

}
