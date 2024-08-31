package ethash

import (
	"fmt"
	"time"

	engine "github.com/shreekarashastry/ethash/ethash"
	"github.com/shreekarashastry/vdf-simulation/RQ1/stats"
)

var (
	max_time       = 5 * time.Second
	num_iterations = 50
)

func RunEthash() {
	var ethashHashesEachRound []uint64
	ethash := engine.New(engine.Config{DatasetDir: ".",
		DatasetsInMem:  1,
		DatasetsOnDisk: 1})
	for i := 00; i < num_iterations; i++ {

		data := []byte("hello world!")
		stop := make(chan struct{})

		timer := time.NewTimer(max_time)
		// Run sha vdf for max_time
		go ethash.SealRandomData(nil, data, stop)
		select {
		case <-timer.C:
			ethashHashesEachRound = append(ethashHashesEachRound, uint64(ethash.NumHashes()))
			close(stop)
		}
	}
	fmt.Println("Algo", "Ethash", "Duration", max_time.Seconds(), "s", "Mean", stats.Mean(ethashHashesEachRound), "Median", stats.Median(ethashHashesEachRound))
}
