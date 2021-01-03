package main

import (
	"fmt"
	"unicode"

	"github.com/chrisport/go-lang-detector/langdet/langdetdef"
)

func main() {

	// "love" translated to different languages
	input := []string{"爱", "好き", "ආදරය", "cinta", "حب", "مينه", "ฮัก", "माया"}

	for k, v := range input {

		if checkifLatin(v) {
			fmt.Println(k, "["+v+"] is a Latin or a derivative of Latin ? ", checkifLatin(v))
		} else if checkifArabic(v) {
			fmt.Println(k, "["+v+"] is an Arabic or a derivative of Arabic ? ", checkifArabic(v))
		} else if checkifThai(v) {
			fmt.Println(k, "["+v+"] is a Thai or a derivative of Thai ? ", checkifThai(v))
		} else if checkifHan(v) {
			fmt.Println(k, "["+v+"] is a Han(Chinese) or a derivative of Han(Chinese) ? ", checkifHan(v))
		} else if checkifDevanagari(v) {
			fmt.Println(k, "["+v+"] is a Devanagari or a derivative of Devanagari ? ", checkifDevanagari(v))
		} else {
			// throw in an unknown script Japanese and Sinhala just to test out. To detect, use unicode.Sinhala and unicode.Hiragana
			fmt.Println(k, "["+v+"] unable to detect script type. Need new function to detect this script type. ")
		}

	}

	fmt.Println("----------------------------------------------------------------------------")

	// alternative method, but....it is not what I want.
	detector := langdetdef.NewWithDefaultLanguages()
	result := detector.GetClosestLanguage(input[4])
	fmt.Println("GetClosestLanguage returns : ", result)

}

// checkifLatin - check if input is Latin scripts or not.
func checkifLatin(input string) bool {

	var isLatin = false

	for _, v := range input {
		if unicode.In(v, unicode.Latin) {
			isLatin = true
		} else {
			isLatin = false
		}
	}
	return isLatin
}

// checkifThai - check if input is Thai scripts or not.
func checkifThai(input string) bool {

	var isThai = false

	for _, v := range input {
		if unicode.In(v, unicode.Thai) {
			isThai = true
		} else {
			isThai = false
		}
	}
	return isThai
}

// checkifArabic is Arabic characters or not.
func checkifArabic(input string) bool {

	var isArabic = false

	for _, v := range input {
		if unicode.In(v, unicode.Arabic) {
			isArabic = true
		} else {
			isArabic = false
		}
	}
	return isArabic
}

// checkifHan is Han(Chinese) characters or not.
func checkifHan(input string) bool {

	var isHan = false

	for _, v := range input {
		if unicode.In(v, unicode.Han) {
			isHan = true
		} else {
			isHan = false
		}
	}
	return isHan
}

// checkifDevanagari - check if input is Devanagari scripts or not.
func checkifDevanagari(input string) bool {

	var isDevanagari = false

	for _, v := range input {
		if unicode.In(v, unicode.Devanagari) {
			isDevanagari = true
		} else {
			isDevanagari = false
		}
	}
	return isDevanagari
}
