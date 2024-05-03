package main

import (
	"fmt"
	"github.com/Warh40k/entropy"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("example/bib")
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		os.Exit(1)
	}
	freqs, probs := entropy.GetFreqsProbs(data)
	entr := entropy.GetEntropy(probs)
	cond := entropy.GetConditionalProbs(data, freqs)
	condEntr := entropy.GetCondEntropy(probs, cond)
	fmt.Println(cond)
	fmt.Println("Entropy: ", entr)
	fmt.Println("Conditional entropy: ", condEntr)
}
