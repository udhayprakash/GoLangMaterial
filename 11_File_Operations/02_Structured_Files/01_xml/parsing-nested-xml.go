package main

import (
	"encoding/xml"
	"fmt"

	"log"
	"strings"
)

func main() {
	const dat1 = `
<?xml version="1.0" encoding="UTF-8"?>
<someprefix:Response Ticket="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" When="1985-03-24T20:45:18Z" Passthrough="https://somewebsite.com/prettypage" xmlns:someprefix="a:b:c:d:e:f:g" VersionNumber="3.6">
	<someotherprefix:Owner xmlns:someotherprefix="h:i:j:k:l:m:n">https://somewebsite.com</someotherprefix:Owner>
	<someprefix:Status>
		<someprefix:StatusCode Value="GoodThingsHappened"/>
	</someprefix:Status>
	<someotherprefix:Response2 VersionNumber="3.6" Ticket="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" When="1985-03-24T20:45:18Z" xmlns:someotherprefix="h:i:j:k:l:m:n">
		<someotherprefix:Owner>https://somewebsite.com</someotherprefix:Owner>
		<qq:Signature xmlns:qq="http://www.somewebsite.com">
			<qq:SpeculativeInfo xmlns:qq="http://www.somewebsite.com">
				<qq:BarbequeMethod Method="http://www.somewebsite.com"/>
				<qq:SauteMethod Method="http://somewebsite.com"/>
				<qq:Reference USB="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX">
					<qq:Functions>
						<qq:Function Method="http://www.somewebsite.com"/>
						<qq:Function Method="http://www.somewebsite.com">
							<ec:ExecutiveNamespaces xmlns:ec="http://www.somewebsite.com" PrefixList="ys"/>
						</qq:Function>
					</qq:Functions>
					<qq:HoolaHoopMethod Method="http://somewebsitehere.com"/>
					<qq:HoolaHoopValue>BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB</qq:HoolaHoopValue>
				</qq:Reference>
			</qq:SpeculativeInfo>
			<qq:LongStringValue>AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA</qq:LongStringValue>
			<qq:HeaderInfo>
				<qq:Data>
					<qq:SubData>AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA</qq:SubData>
				</qq:Data>
			</qq:HeaderInfo>
		</qq:Signature>
		<someotherprefix:Subject>
			<someotherprefix:NameID Format="a:b:c:d:e:f:g">ABCDEFG</someotherprefix:NameID>
			<someotherprefix:SubjectConfirmation Method="a:b:c:d:e:f:g:h">
				<someotherprefix:SubjectConfirmationData NotOnOrAfter="1996-10-21T01:02:17Z" Recipient="https://somewebsite.com"/>
			</someotherprefix:SubjectConfirmation>
		</someotherprefix:Subject>
		<someotherprefix:Conditions NotBefore="1992-05-12T22:47:17Z" NotOnOrAfter="1991-02-11T04:04:97Z">
			<someotherprefix:PeopleRestriction>
				<someotherprefix:People>a:b:c:d:e:f</someotherprefix:People>
			</someotherprefix:PeopleRestriction>
		</someotherprefix:Conditions>
		<someotherprefix:TknStatement TknInstant="1998-04-12T06:33:33Z" TokenIndex="ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ">
			<someotherprefix:TknContext>
				<someotherprefix:TknContextClassRef>a:b:c:d:e:f:g</someotherprefix:TknContextClassRef>
			</someotherprefix:TknContext>
		</someotherprefix:TknStatement>
		<someotherprefix:TopLevel>
			<someotherprefix:SubLevel1 xmlns:someotherprefix="h:i:j:k:l:m:n" Name="https://somewebsite.com" NameFormat="a:b:c:d:e:f:g">
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">IWantThisValue1</someotherprefix:SubLevel2>
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">IWantThisValue2</someotherprefix:SubLevel2>
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">IWantThisValue3</someotherprefix:SubLevel2>
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">IWantThisValue4</someotherprefix:SubLevel2>
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">IWantThisValue5</someotherprefix:SubLevel2>
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">IWantThisValue6</someotherprefix:SubLevel2>
			</someotherprefix:SubLevel1>
			<someotherprefix:SubLevel1 xmlns:someotherprefix="h:i:j:k:l:m:n" Name="https://somewebsite.com" NameFormat="a:b:c:d:e:f:g:h">
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">Something@Email.com</someotherprefix:SubLevel2>
			</someotherprefix:SubLevel1>
			<someotherprefix:SubLevel1 xmlns:someotherprefix="h:i:j:k:l:m:n" Name="https://somewebsite.com/suffix" NameFormat="a:b:c:d:e:f">
				<someotherprefix:SubLevel2 xmlns:someotherprefix="h:i:j:k:l:m:n">99999999999999999</someotherprefix:SubLevel2>
			</someotherprefix:SubLevel1>
		</someotherprefix:TopLevel>
	</someotherprefix:Response2>
</someprefix:Response>
`

	reader := strings.NewReader(dat1)
	decoder := xml.NewDecoder(reader)

	//https://stackoverflow.com/questions/59615418/how-to-parse-xml-in-slice-format
	type StuffList struct {
		Stuff []string `xml:"Response2>TopLevel>SubLevel1>SubLevel2"`
	}

	var sl StuffList

	err := decoder.Decode(&sl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rl:", sl.Stuff)
}
