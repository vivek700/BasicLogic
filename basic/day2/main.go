package main

import "fmt"

func eligibleForVote(age int) string {
	var min_age int = 18
	if age >= min_age {
		return "Eligible to vote"
	}
	return "Not eligible to vote"
}

func eligibleForDrivingLicense(age int) string {
	min_age := 16
	if age >= min_age {
		return "Eligible for driving license"
	}
	return "Not eligible for driving license"
}

func eligibleForSeniorCitizenDiscount(age int) string {
	min_age := 60

	if age >= min_age {
		return "Eligible for senior citizen discount"
	}
	return "Not eligible for senior citizen discount"
}

func main() {

	fmt.Println("Checking eligibility for voting...")
	fmt.Println(eligibleForVote(20))
	fmt.Println(eligibleForVote(16))
	fmt.Println(eligibleForVote(23))
	fmt.Println(eligibleForVote(18))

	fmt.Println("Checking eligibility for driving license...")
	fmt.Println(eligibleForDrivingLicense(14))
	fmt.Println(eligibleForDrivingLicense(16))
	fmt.Println(eligibleForDrivingLicense(17))

	fmt.Println("Checking eligibility for senior citizen discount...")
	fmt.Println(eligibleForSeniorCitizenDiscount(65))
	fmt.Println(eligibleForSeniorCitizenDiscount(55))
	fmt.Println(eligibleForSeniorCitizenDiscount(60))

}
