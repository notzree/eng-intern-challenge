package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	var input string
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string to translate")
		return
	}
	// Join with a space to preservce spaces in input.
	input = strings.TrimSpace(strings.Join(os.Args[1:], " "))
	isEnglish := IsEnglish(input)
	var result string
	var err error
	if isEnglish {
		result, err = EnglishToBrailleTranslator(input)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		result, err = BrailleToEnglishTranslator(input) //todo: Implement
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(result)
}

// IsEnglish is a function that returns true if the input is english, and false if the input is braille.
func IsEnglish(input string) bool {
	const (
		dotRune      = '.'
		capitalORune = 'O'
	)
	for _, r := range input {
		isBrailleRune := (r == dotRune || r == capitalORune)
		if !isBrailleRune {
			return true
		}
	}
	return false
}

func EnglishToBrailleTranslator(input string) (string, error) {
	const (
		CapitalFollows string = ".....O"
		DecimalFollows string = ".O...O"
		NumberFollows  string = ".O.OOO"
	)
	lookup := GetEnglishToBrailleLookup()
	// inputRune = rune(input)
	var result strings.Builder
	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		if unicode.IsDigit(char) {
			var number strings.Builder
			j := i
			for j < len(input) && unicode.IsDigit(rune(input[j])) {
				numberChar := rune(input[j])
				if braille, ok := lookup[numberChar]; ok {
					number.WriteString(braille)
				} else {
					log.Fatal("unable to convert from %v", numberChar)
				}
				j += 1
			}
			result.WriteString(NumberFollows)
			result.WriteString(number.String())
			i = j - 1
		} else {
			if unicode.IsUpper(char) {
				result.WriteString(CapitalFollows)
			}
			if braille, ok := lookup[unicode.ToLower(char)]; ok {
				result.WriteString(braille)
			} else {
				log.Fatal("unable to convert from %v", char)
			}
		}
	}
	return result.String(), nil
}

// GetEnglishToBrailleLookup returns the character lookup table for converting English into Braille.
func GetEnglishToBrailleLookup() map[rune]string {
	return map[rune]string{
		// Lowercase letters and space
		'a': "O.....", 'b': "O.O...", 'c': "OO....", 'd': "OO.O..", 'e': "O..O..",
		'f': "OOO...", 'g': "OOOO..", 'h': "O.OO..", 'i': ".OO...", 'j': ".OOO..",
		'k': "O...O.", 'l': "O.O.O.", 'm': "OO..O.", 'n': "OO.OO.", 'o': "O..OO.",
		'p': "OOO.O.", 'q': "OOOOO.", 'r': "O.OOO.", 's': ".OO.O.", 't': ".OOOO.",
		'u': "O...OO", 'v': "O.O.OO", 'w': ".OOO.O", 'x': "OO..OO", 'y': "OO.OOO",
		'z': "O..OOO", ' ': "......",

		// Numbers
		'1': "O.....", '2': "O.O...", '3': "OO....", '4': "OO.O..", '5': "O..O..",
		'6': "OOO...", '7': "OOOO..", '8': "O.OO..", '9': ".OO...", '0': ".OOO..",
	}
}
