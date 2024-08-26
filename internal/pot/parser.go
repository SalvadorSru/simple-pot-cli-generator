package pot

import (
	"salvadorsru/po-generator/internal/action"
	"unicode"
)

func ParseEntryFromString(entryReference Reference, sentence string, parameters action.Parameters) Entry {

	var (
		entry             Entry
		parameterPosition = 0
		inString          = false
		inVariable        = false
		inNumber          = false
		startingQuote     rune
		currentString     []rune
	)

	entry.Reference = entryReference

	for position, letter := range sentence {
		if isQuote(letter) {
			if !inString {
				inString = true
				startingQuote = letter
				continue
			} else {
				if (position == 0 || sentence[position-1] != '\\') && letter == startingQuote {
					inString = false

					switch parameterPosition {
					case parameters.Singular:
						entry.Singular = []rune(currentString)
					case parameters.Plural:
						entry.Plural = []rune(currentString)
					case parameters.Context:
						entry.Context = []rune(currentString)
					}

					parameterPosition++
					currentString = nil
				}
			}
		}

		if !inString && unicode.IsPrint(letter) {
			inVariable = true
		}

		if inVariable && letter == ',' {
			inVariable = false
			parameterPosition++
		}

		if !inString && unicode.IsDigit(letter) {
			inNumber = true
		}

		if inNumber && letter == ',' {
			inNumber = false
			parameterPosition++
		}

		if inString {
			currentString = append(currentString, letter)
		}
	}

	return entry
}
