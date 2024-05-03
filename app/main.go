package main

import (
	"fmt"
	"github.com/Warh40k/entropy"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("example/book1")
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		os.Exit(1)
	}
	freqs, probs := entropy.GetFreqsProbs(data)
	entr := entropy.GetEntropy(probs)
	condProbs, condFreqs := entropy.GetCondProbs(data, freqs)
	condEntr := entropy.GetCondEntropy(probs, condProbs)
	condProbsXX := entropy.GetCondProbsXX(data, condFreqs)
	condEntrXX := entropy.GetCondEntropyXX(probs, condProbsXX)
	fmt.Println(condProbsXX)
	fmt.Println("Entropy H(X): ", entr)
	fmt.Println("Conditional entropy H(X|X): ", condEntr)
	fmt.Println("Conditional entropy H(X|XX): ", condEntrXX)
}
