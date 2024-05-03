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

func GetCondProbs(seq []byte, freqs [256]float64) (map[byte]map[byte]float64, map[byte]map[byte]float64) {
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

	return condProbs, condFreqs
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

func GetCondProbsXX(seq []byte, condFreqs map[byte]map[byte]float64) map[byte]map[byte]map[byte]float64 {
	var condFreqsXX, condProbsXX = make(map[byte]map[byte]map[byte]float64), make(map[byte]map[byte]map[byte]float64)
	for i := 2; i < len(seq); i++ {
		cur := seq[i]
		prev := seq[i-1]
		prev2 := seq[i-2]
		if condFreqsXX[prev2] == nil {
			condFreqsXX[prev2] = make(map[byte]map[byte]float64)
		}
		if condFreqsXX[prev2][prev] == nil {
			condFreqsXX[prev2][prev] = make(map[byte]float64)
		}
		condFreqsXX[prev2][prev][cur]++
	}

	for prev2 := range condFreqsXX {
		if condProbsXX[prev2] == nil {
			condProbsXX[prev2] = make(map[byte]map[byte]float64)
		}
		for prev := range condFreqsXX[prev2] {
			if condProbsXX[prev2][prev] == nil {
				condProbsXX[prev2][prev] = make(map[byte]float64)
			}
			for cur := range condFreqsXX[prev2][prev] {
				condProbsXX[prev2][prev][cur] = condFreqsXX[prev2][prev][cur] / condFreqs[prev2][prev]
			}
		}
	}
	return condProbsXX
}

func GetCondEntropyXX(probs [256]float64, condProbs map[byte]map[byte]float64, condProbsXX map[byte]map[byte]map[byte]float64) float64 {
	var entropy float64
	for prev2 := range condProbsXX {
		for prev := range condProbsXX[prev2] {
			var temp float64
			for _, condProbXX := range condProbsXX[prev2][prev] {
				temp += condProbXX * math.Log2(condProbXX)
			}
			entropy += condProbs[prev2][prev] * probs[prev2] * temp
		}
	}
	return -entropy
}
