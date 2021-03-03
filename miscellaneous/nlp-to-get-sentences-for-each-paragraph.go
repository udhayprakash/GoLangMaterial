package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	text := `EURUSD’s uptrend may be accelerated if US consumer confidence comes in below the 131.0 estimate and Fed Chairman Jerome Powell’s economic outlook bolsters market expectations of rate cuts. Overnight index swaps are pricing in a 100 percent probability of a cut from the July meeting through year-end. However, rhetoric from the central bank has not indicated that policymakers are feeling dovish to that degree.
 	However, hawkish members of the Fed are finding it increasingly difficult to justify their position in light of US growth. Since February, economic activity out of the US has been broadly underperforming relative to economists’ expectations – signaling that analysts are over estimating the economy’s strength. Inflationary pressure has also been waning alongside a deterioration in global trade due to the US-China trade war.`

	// count how many paragraph
	paragraph := strings.Split(text, "\n")
	//fmt.Println("Paragraph 0 : ", paragraph[0])
	//fmt.Println("Paragraph 1 : ", paragraph[1])

	// Create a new document for each paragraph
	for k, v := range paragraph {
		fmt.Println("Processing paragraph ", k, " : ")
		doc, err := prose.NewDocument(v)
		if err != nil {
			log.Fatal(err)
		}

		// Iterate over the doc's sentences:
		for _, sent := range doc.Sentences() {
			fmt.Println("[" + sent.Text)
		}
	}
}
