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

func GetConditionalProbs(seq []byte, freqs [256]float64) [256][256]float64 {
	var condFreqs, condProbs [256][256]float64
	for i := range 256 {
		if freqs[i] != 0 {
			for j := 1; j < len(seq); j++ {
				if int(seq[j-1]) == i {
					condFreqs[i][seq[j]]++
				}
			}
			for j := range 256 {
				if freqs[j] != 0 {
					condProbs[i][j] = condFreqs[i][j] / freqs[i]
				}
			}
		}
	}

	return condProbs
}

func GetCondEntropy(probs [256]float64, condProbs [256][256]float64) float64 {
	var entropy float64
	for i, prob := range probs {
		var batch float64
		for _, condProb := range condProbs[i] {
			if condProb != 0 {
				batch += condProb * math.Log2(condProb)
			}
		}
		entropy -= prob * batch
	}
	return entropy
}
