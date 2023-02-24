package main

import (
	"fmt"
	"time"
)

func calculate(birthDateString string) bool {

	// declare var
	birthDate, _ := time.Parse("2006-01-02", birthDateString)
	startDate, _ := time.Parse("2006-01-02", "2021-06-01")
	endDate, _ := time.Parse("2006-01-02", "2021-08-31")

	//get Age at first date and last date of Vaccine
	ageAtStartDate := startDate.Sub(birthDate).Hours() / 24 / 365
	ageAtEndDate := endDate.Sub(birthDate).Hours() / 24 / 365
	//fmt.Println(ageAtStartDate)
	//fmt.Println(ageAtEndDate)
	if ageAtStartDate > 65 || (ageAtStartDate > 0.5 && ageAtStartDate < 2) {
		//fmt.Println("1")
		return true
	} else if (ageAtStartDate < 65 && ageAtEndDate >= 65) || ((ageAtStartDate < 0.5 && ageAtEndDate >= 0.5) && ageAtEndDate < 2) {
		//fmt.Println("2")
		return true
	} else if ageAtEndDate < 65 || (ageAtEndDate < 0.5 && ageAtEndDate < 0.5) || (ageAtStartDate > 2 && ageAtStartDate < 10) {
		//fmt.Println("3")
		return false
	}

	return false
}

func getStartDate(birthDateString string) string {
	//res := "2021/06/01"
	// declare var
	birthDate, _ := time.Parse("2006-01-02", birthDateString)
	startDate, _ := time.Parse("2006-01-02", "2021-06-01")
	endDate, _ := time.Parse("2006-01-02", "2021-08-31")
	res := startDate.Format("2006-01-02")

	//get Age at first date and last date of Vaccine
	ageAtStartDate := startDate.Sub(birthDate).Hours() / 24 / 365
	ageAtEndDate := endDate.Sub(birthDate).Hours() / 24 / 365

	//fmt.Println(ageAtStartDate)
	if (ageAtStartDate < 65 && ageAtEndDate >= 65) || ((ageAtStartDate < 0.5 && ageAtEndDate >= 0.5) && ageAtEndDate < 2) {
		if ageAtStartDate < 65 && ageAtEndDate >= 65 {
			birthDate = birthDate.AddDate(65, 0, 0)
			res = birthDate.Format("2006-01-02")
		} else if (ageAtStartDate < 0.5 && ageAtEndDate >= 0.5) && ageAtEndDate < 2 {
			birthDate = birthDate.AddDate(0, 6, 0)
			res = birthDate.Format("2006-01-02")
		}
	}
	return res
}

func getLastDate(birthDateString string) string {
	//res := "2021/06/01"
	// declare var
	birthDate, _ := time.Parse("2006-01-02", birthDateString)
	startDate, _ := time.Parse("2006-01-02", "2021-06-01")
	endDate, _ := time.Parse("2006-01-02", "2021-08-31")
	res := endDate.Format("2006-01-02")

	//get Age at first date and last date of Vaccine
	ageAtStartDate := startDate.Sub(birthDate).Hours() / 24 / 365
	ageAtEndDate := endDate.Sub(birthDate).Hours() / 24 / 365

	//fmt.Println(ageAtStartDate)
	if ageAtStartDate < 2 && ageAtEndDate > 2 {
		birthDate = birthDate.AddDate(2, 0, 0)
		res = birthDate.Format("2006-01-02")
	}
	return res
}

func Results(gender string, birthDateString string) (string, string, string) {
	eligibleFlag := "N"
	serviceStartDate := ""
	serviceEndDate := ""

	birthDate, err := time.Parse("2006-01-02", birthDateString)
	if err != nil {
		eligibleFlag = "X"
		serviceStartDate = "invalid date"
		serviceEndDate = "invalid date"
		return eligibleFlag, serviceStartDate, serviceEndDate
	}
	if birthDate.Year() > 2100 { // handle กรณีใส่เป็น พศ ครับ ขอ assume ว่าไม่มีใครอายุเยอะขนาดนี้นะครับ 55555555
		birthDate = birthDate.AddDate(-543, 0, 0)
		birthDateString = birthDate.Format("2006-01-02")
		//fmt.Println(birthDateString)
	}
	_ = gender // decoy ไม่มีให้ response แต่โจทย์บอกให้รับ paramater มาด้วย เลยสร้าง decoy ไว้เพื่อให้ build ได้ครับ

	isCanVac := calculate(birthDateString) // เช็คว่ามีสิทธิฉีด vaccine ไหม ถ้าไม่จะ return defult value
	if isCanVac {
		eligibleFlag = "Y"
		serviceStartDate = getStartDate(birthDateString)
		serviceEndDate = getLastDate(birthDateString)
	}
	return eligibleFlag, serviceStartDate, serviceEndDate
}

func main() {
	eligibleFlag, serviceStartDate, serviceEndDate := Results("Woman", "2564-02-28")
	fmt.Println("eligible Flag = ", eligibleFlag)
	fmt.Println("service Start Date = ", serviceStartDate)
	fmt.Println("service End Date = ", serviceEndDate)
}
