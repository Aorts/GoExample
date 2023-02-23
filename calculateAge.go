package main

import (
	"fmt"
	"time"
)

func calculate(birthDateString string) bool {

	// declare var
	birthDate, _ := time.Parse("2006/01/02", birthDateString)
	startDate, _ := time.Parse("2006/01/02", "2021/06/01")
	endDate, _ := time.Parse("2006/01/02", "2021/08/31")

	//get Age at first date and last date of Vaccine
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

func getStartDate(birthDateString string) string {
	//res := "2021/06/01"
	// declare var
	birthDate, _ := time.Parse("2006/01/02", birthDateString)
	startDate, _ := time.Parse("2006/01/02", "2021/06/01")
	endDate, _ := time.Parse("2006/01/02", "2021/08/31")
	res := startDate.Format("2006/01/02")

	//get Age at first date and last date of Vaccine
	ageAtStartDate := startDate.Sub(birthDate).Hours() / 24 / 365
	ageAtEndDate := endDate.Sub(birthDate).Hours() / 24 / 365

	//fmt.Println(ageAtStartDate)
	if (ageAtStartDate < 65 && ageAtEndDate >= 65) || ((ageAtStartDate < 0.5 && ageAtEndDate >= 0.5) && ageAtEndDate < 2) {
		if ageAtStartDate < 65 && ageAtEndDate >= 65 {
			birthDate = birthDate.AddDate(65, 0, 0)
			res = birthDate.Format("2006/01/02")
		} else if (ageAtStartDate < 0.5 && ageAtEndDate >= 0.5) && ageAtEndDate < 2 {
			res = startDate.Format("2006/01/02")
		}
	}
	return res
}

func Results(gender string, birthDateString string) (string, string, string) {
	eligibleFlag := "N"
	serviceStartDate := ""
	serviceEndDate := ""

	isCanVac := calculate(birthDateString)
	if isCanVac {
		eligibleFlag = "Y"
		serviceStartDate = getStartDate(birthDateString)
	}

	return eligibleFlag, serviceStartDate, serviceEndDate
}

func main() {
	eligibleFlag, serviceStartDate, serviceEndDate := Results("man", "2021/01/05")
	fmt.Println("eligible Flag = ", eligibleFlag)
	fmt.Println("service Start Date = ", serviceStartDate)
	fmt.Println("service End Date = ", serviceEndDate)
}
