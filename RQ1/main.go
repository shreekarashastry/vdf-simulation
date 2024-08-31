package main

import (
	"github.com/shreekarashastry/vdf-simulation/RQ1/ethash"
	"github.com/shreekarashastry/vdf-simulation/RQ1/sha256"
)

const consensus = "ethash"

func main() {
	switch consensus {
	case "sha":
		sha256.RunShaSim()
	case "ethash":
		ethash.RunEthash()
	}
}
