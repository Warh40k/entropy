package entropy

import "math"

func GetFreqsProbs(seq []byte) ([256]float64, [256]float64) {
	var freqs, probs [256]float64
	length := float64(len(seq))
	for _, b := range seq {
		freqs[b]++
	}
	for j := range 256 {
		probs[j] = freqs[j] / length
	}

	return freqs, probs
}

func GetEntropy(probs [256]float64) float64 {
	var entropy float64
	for _, p := range probs {
		if p == 0 {
			continue
		}
		entropy -= p * math.Log2(p)
	}
	return entropy
}

func GetConditionalProbs(seq []byte, freqs [256]float64) map[byte]map[byte]float64 {
	var condFreqs, condProbs = make(map[byte]map[byte]float64), make(map[byte]map[byte]float64)
	for i := 1; i < len(seq); i++ {
		cur := seq[i]
		prev := seq[i-1]
		if condFreqs[prev] == nil {
			condFreqs[prev] = make(map[byte]float64)
		}
		condFreqs[prev][cur]++
	}

	for prev := range condFreqs {
		if condProbs[prev] == nil {
			condProbs[prev] = make(map[byte]float64)
		}
		for cur := range condFreqs[prev] {
			condProbs[prev][cur] = condFreqs[prev][cur] / freqs[prev]
		}
	}

	return condProbs
}

func GetCondEntropy(probs [256]float64, condProbs map[byte]map[byte]float64) float64 {
	var entropy float64
	for prev, prob := range probs {
		var temp float64
		for _, condProb := range condProbs[byte(prev)] {
			temp += condProb * math.Log2(condProb)
		}
		entropy += prob * temp
	}
	return -entropy
}
