package stats

import (
	"math"
	"sort"
)

func Median(data []uint64) uint64 {
	dataCopy := make([]uint64, len(data))
	copy(dataCopy, data)

	sort.Slice(dataCopy, func(i, j int) bool { return i < j })

	var median uint64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}

func Mean(data []uint64) uint64 {
	var total uint64 = 0
	for _, v := range data {
		total += v
	}

	return uint64(math.Round(float64(total) / float64(len(data))))
}
