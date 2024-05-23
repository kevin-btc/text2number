package alphatodigit

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var unitMap = map[string]int{
	"zero": 0, "un": 1, "une": 1, "deux": 2, "trois": 3, "quatre": 4, "cinq": 5,
	"six": 6, "sept": 7, "huit": 8, "neuf": 9, "dix": 10, "onze": 11,
	"douze": 12, "treize": 13, "quatorze": 14, "quinze": 15, "seize": 16,
}

var tensMap = map[string]int{
	"vingt": 20, "trente": 30, "quarante": 40, "cinquante": 50,
	"soixante": 60, "soixante-dix": 70, "quatre-vingt": 80, "quatre-vingt-dix": 90,
}

var magnitudeMap = map[string]int{
	"cent": 100, "cents": 100, "mille": 1000, "million": 1000000,
	"millions": 1000000, "milliard": 1000000000, "milliards": 1000000000,
}

func getAllKeys() []string {
	var keys []string
	for key := range unitMap {
		keys = append(keys, key)
	}
	for key := range tensMap {
		keys = append(keys, key)
	}
	for key := range magnitudeMap {
		keys = append(keys, key)
	}
	return keys
}

func ReplaceCurrencyWordsWithSymbols(sentence string) string {
    regexPattern := `(\d+(?:,\d+)?\s*)(euro|euros|dollar|dollars|livre|livres|yen)\b`
    regex := regexp.MustCompile(regexPattern)

    return regex.ReplaceAllStringFunc(sentence, func(match string) string {
        submatches := regex.FindStringSubmatch(match)
        numberWithOptionalSpaces := submatches[1]
        currencyWord := submatches[2]

        numberWithNoSpaces := strings.ReplaceAll(numberWithOptionalSpaces, " ", "")

        var symbol string
        switch currencyWord {
        case "euro", "euros":
            symbol = "€"
        case "dollar", "dollars":
            symbol = "$"
        case "livre", "livres":
            symbol = "£"
        case "yen":
            symbol = "¥"
        default:
            symbol = currencyWord
        }

        return numberWithNoSpaces + symbol
    })
}

func ReplaceHyphenatedWordsWithSpaces(sentence string) string {
	keys := getAllKeys()
	regexPattern := `\b(` + strings.Join(keys, "|") + `)(-(` + strings.Join(keys, "|") + `))+\b`
	regex := regexp.MustCompile(regexPattern)

	return regex.ReplaceAllStringFunc(sentence, func(match string) string {
		return strings.ReplaceAll(match, "-", " ")
	})
}

func ReplacePercentageWordsWithSymbol(sentence string) string {
    regexPattern := `(\d+)\s*(pourcent|pourcents)\b`
    regex := regexp.MustCompile(regexPattern)

    return regex.ReplaceAllStringFunc(sentence, func(match string) string {
        cleanMatch := regexp.MustCompile(`\s*(pourcent|pourcents)\b`).ReplaceAllString(match, "%")
        return cleanMatch
    })
}


type Options struct {
	ReplaceCurrencyWordsWithSymbols bool
	ReplacePercentageWordsWithSymbol bool
}

func Alpha2Digit(text string , options Options) (string, error) {

	transformedText := text    

    cmd := exec.Command("alpha2digit/search_and_replace_by_num", transformedText)

    output, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    transformedText = string(output)
	transformedText = ReplaceHyphenatedWordsWithSpaces(transformedText)


    if options.ReplaceCurrencyWordsWithSymbols {
        transformedText = ReplaceCurrencyWordsWithSymbols(transformedText)
    }

    if options.ReplacePercentageWordsWithSymbol {
        transformedText = ReplacePercentageWordsWithSymbol(transformedText)
    }

    return strings.Trim(transformedText, "\"'"), nil
}