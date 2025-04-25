package main

func bulkSend(numMessages int) float64 {
	cost := float64(0)
	for i := range numMessages {
		cost += 1 + 0.01*float64(i)
	}
	return cost
}
