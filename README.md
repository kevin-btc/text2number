# AlphaToDigit

## Introduction

AlphaToDigit is a Go package that provides a function to convert alphabetic words to their corresponding digit representations.

## Installation

To use AlphaToDigit, you need to have Go installed on your system. You can install AlphaToDigit by running the following command:

```

go get github.com/kevin-btc/text2number/alpha2digit

```

## Usage

```go
package main

import (
	"fmt"
	ttn "github.com/kevin-btc/text2number/alpha2digit"
)

func main() {
	sentence := "I have two apples and three bananas."
	digitSentence, err := ttn.Alpha2Digit(sentence)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(digitSentence)
}

// "I have 2 apples and 3 bananas"
```

## Function

### Alpha2Digit

`func Alpha2Digit(sentence string) (string, error)`

Alpha2Digit takes an alphabetic sentence as input and returns the sentence with alphabetic words replaced by their corresponding digit representations.
