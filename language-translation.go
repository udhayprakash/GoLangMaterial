package main

import (
	//"fmt" -- replace with message package in this example

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"golang.org/x/text/message"
)

func main() {

	// define a direct translation from English to Dutch
	message.SetString(language.Dutch, "In %v people speak %v.", "In %v spreekt men %v.")

	fr := language.French
	region, _ := fr.Region()
	for _, tag := range []string{"en", "nl"} {
		p := message.NewPrinter(language.Make(tag))

		p.Printf("In %v people speak %v.", display.Region(region), display.Language(fr))
		p.Println()
	}

	// define a direct translation from English to Chinese
	message.SetString(language.Chinese, "In %v people speak %v.", "在%v说%v.") // * Must match for translation to work

	zh := language.Chinese
	region, _ = zh.Region()
	for _, tag := range []string{"en", "zh"} {
		p := message.NewPrinter(language.Make(tag))

		p.Printf("In %v people speak %v.", display.Region(region), display.Language(zh)) // * Must match for translation to work
		p.Println()
	}

	// define a direct translation from English to Malay
	message.SetString(language.Malay, "In %v people speak %v.", "Orang %v berbahasa %v.")

	msmy := language.Malay
	region, _ = msmy.Region()
	for _, tag := range []string{"en", "ms"} {
		p := message.NewPrinter(language.Make(tag))

		p.Printf("In %v people speak %v.", display.Region(region), display.Language(msmy))
		p.Println()
	}
}

/*
Ref:
https://godoc.org/golang.org/x/text/message

https://godoc.org/golang.org/x/text/language/display#Dictionary.Languages
*/
