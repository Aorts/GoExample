package main

import (
	"fmt"
	"time"
)

type Response struct {
	eligibleFlag     string
	serviceStartDate string
	serviceEndDate   string
}

func calculate(birthDateString string) bool {
	birthDate, _ := time.Parse("2006/01/02", birthDateString)
	startDate, _ := time.Parse("2006/01/02", "2021/06/01")
	endDate, _ := time.Parse("2006/01/02", "2021/08/31")
	//age := time.Since(birthDate).Hours() / 24 / 365
	ageAtStartDate := startDate.Sub(birthDate).Hours() / 24 / 365
	ageAtEndDate := endDate.Sub(birthDate).Hours() / 24 / 365
	//fmt.Println(ageAtStartDate)
	//fmt.Println(ageAtEndDate)
	if ageAtStartDate > 65 || (ageAtStartDate > 0.5 && ageAtStartDate < 2) {
		return true
	} else if (ageAtStartDate < 65 && ageAtEndDate >= 65) || ((ageAtStartDate < 0.5 && ageAtEndDate >= 0.5) && ageAtEndDate < 2) {
		return true
	} else if ageAtEndDate < 65 || (ageAtEndDate < 0.5 && ageAtEndDate < 0.5) || (ageAtStartDate > 2 && ageAtStartDate < 10) {
		return false
	}
	return false
}

func main() {
	aage := calculate("1956/09/01")
	fmt.Println(aage)
}
