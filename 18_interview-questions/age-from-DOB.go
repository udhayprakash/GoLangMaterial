package main

import (
	"fmt"
	"strconv"
	"time"
)

func monthToInt(month string) int {
	switch month {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	default:
		panic("Unrecognized month")
	}
}

func ageFromDateOfBirth(dob string) (int, int) {
	layOut := "02/01/2006" // dd/mm/yyyy
	dobTime, err := time.Parse(layOut, dob)

	if err != nil {
		panic(err)
	}

	var ageYear, leapAge int
	ageYear = time.Now().Year() - dobTime.Year()

	// the trick here is to combine the day and month into an integer of string type

	dobDayMonth, _ := strconv.Atoi(strconv.Itoa(dobTime.Day()) + strconv.Itoa(monthToInt(dobTime.Month().String())))
	nowDayMonth, _ := strconv.Atoi(strconv.Itoa(time.Now().Day()) + strconv.Itoa(monthToInt(time.Now().Month().String())))

	// if the DOB's day + month is larger than today's day + month
	// then the age is still younger by 1 year

	if dobDayMonth > nowDayMonth {
		ageYear = ageYear - 1
	}

	if dobDayMonth == 292 { // dob on 29th Feb - leap year
		leapAge = ageYear / 4
	} else {
		leapAge = 0
	}

	return ageYear, leapAge
}

func main() {

	// dd/mm/yyyy format
	dob := "28/10/1978"
	age, _ := ageFromDateOfBirth(dob)
	fmt.Println(dob, " age is : ", age)

	dob1 := "01/10/1978"
	age1, _ := ageFromDateOfBirth(dob1)
	fmt.Println(dob1, " age is : ", age1)

	dob2 := "02/05/1978"
	age2, _ := ageFromDateOfBirth(dob2)
	fmt.Println(dob2, " age is : ", age2)

	dob3 := "03/05/1978"
	age3, _ := ageFromDateOfBirth(dob3)
	fmt.Println(dob3, " age is : ", age3)

	// February 29, also known as leap day or leap year day,
	// is a date added to most years that are divisible by 4,
	// such as 2008, 2012, 2016, 2020, and 2024.

	dob4 := "29/02/2008" // leap year age test
	age4, leap := ageFromDateOfBirth(dob4)
	fmt.Println(dob4, " age is : ", age4)
	fmt.Println(dob4, " date of birth is 29th Feb and the leap age is : ", leap)

	dob5 := "29/02/1936" // leap year age test
	age5, leap1 := ageFromDateOfBirth(dob5)
	fmt.Println(dob5, " age is : ", age5)
	fmt.Println(dob5, " date of birth is 29th Feb and the leap age is : ", leap1)
}
