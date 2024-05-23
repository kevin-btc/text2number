package main

import (
	"strings"
	"testing"

	ttn "github.com/kevin-btc/text2number/alpha2digit"
)

func TestAlpha2Digit(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		options ttn.Options
		want    string
		wantErr bool
	}{
		{
			name: "Simple number",
			text: "zéro, un, deux, trois",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: false,
				ReplacePercentageWordsWithSymbol: false,
			},
			want: "0, 1, 2, 3",
			wantErr: false,
		},
		{
			name: "Hyphenated numbers",
			text: "vingt-deux trente-trois",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: false,
				ReplacePercentageWordsWithSymbol: false,
			},
			want: "22 33",
			wantErr: false,
		},
		{
			name: "Currency replacement",
			text: "5 dollars 10 euros",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: true,
				ReplacePercentageWordsWithSymbol: false,
			},
			want: "5$ 10€",
			wantErr: false,
		},
		{
			name: "Percentage conversion",
			text: "25 pourcents",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: false,
				ReplacePercentageWordsWithSymbol: true,
			},
			want: "25%",
			wantErr: false,
		},
		{
			name: "Complex hyphenated",
			text: "quatre-vingt-dix-sept",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: false,
				ReplacePercentageWordsWithSymbol: false,
			},
			want: "97",
			wantErr: false,
		},
		{
			name: "Magnitudes",
			text: "quatre mille deux cent trente-six",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: false,
				ReplacePercentageWordsWithSymbol: false,
			},
			want: "4236",
			wantErr: false,
		},
		{
			name: "Complex number with magnitude",
			text: "un million trois cent mille deux cent quatre-vingt",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: false,
				ReplacePercentageWordsWithSymbol: false,
			},
			want: "1300280",
			wantErr: false,
		},
		{
			name: "Non numeric input",
			text: "some non-numeric words",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: false,
				ReplacePercentageWordsWithSymbol: false,
			},
			want: "some non-numeric words",
			wantErr: false,
		},
		{
			name: "Multiple options enabled",
			text: "100 dollars et 50 pourcents",
			options: ttn.Options{
				ReplaceCurrencyWordsWithSymbols: true,
				ReplacePercentageWordsWithSymbol: true,
			},
			want: "100$ et 50%",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ttn.Alpha2Digit(tt.text, tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Alpha2Digit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if strings.TrimSpace(got) != strings.TrimSpace(tt.want) {
				t.Errorf("Alpha2Digit() got = [%v], want [%v]", got, tt.want)
			}

		})
	}
}
