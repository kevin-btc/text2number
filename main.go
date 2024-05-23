package main

import (
	"fmt"

	ttn "github.com/kevin-btc/text2number/alpha2digit"
)

func main() {
	sentence := "deux mille vingt trois europe perd 3000 euros et gagne 5000 dollars mais reperd 1 euro cent pourcent et 2pourcent\n un virgule cinq-milliards d'euros et un million de dollars\n"
	output, err := ttn.Alpha2Digit(sentence, ttn.Options{
		ReplaceCurrencyWordsWithSymbols: true,
		ReplacePercentageWordsWithSymbol: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}