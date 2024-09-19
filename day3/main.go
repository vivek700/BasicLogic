package main

import "fmt"

var min_purchase_amount int = 150
var total_purchase_amount int = 500
var min_spend_amount int = 75

func isEligibleForDiscount(membership bool, purchase_amount int) string {
	if membership && purchase_amount >= min_purchase_amount {
		return "Eligible for discount"
	}
	return "Not eligible for discount"
}

func eligibleForLoyaltyReward(membership bool, total_purchase int) string {
	if membership && total_purchase > total_purchase_amount {
		return "Eligible for loyalty reward"
	}
	return "Not eligible for loyalty reward"
}

func eligibleForFreeshipping(membership bool, purchase_amount int) string {
	if membership && purchase_amount >= min_spend_amount {
		return "Eligible for free shipping"
	}
	return "Not eligible for free shipping"
}

func eligibleForBonusPoint(membership bool, purchase_amount int) string {
	min_spend_amount_for_members := 150
	min_spend_amount := 200
	if membership && purchase_amount >= min_spend_amount_for_members {
		return "Earns bonus points"
	} else if purchase_amount >= min_spend_amount {
		return "Earns bonus points"
	}
	return "Does not earn bonus points"
}

func main() {
	fmt.Println(isEligibleForDiscount(true, 150))
	fmt.Println(isEligibleForDiscount(false, 150))
	fmt.Println(isEligibleForDiscount(true, 100))
	fmt.Println(isEligibleForDiscount(true, 500))

	fmt.Println(eligibleForLoyaltyReward(true, 600))
	fmt.Println(eligibleForLoyaltyReward(true, 450))
	fmt.Println(eligibleForLoyaltyReward(false, 600))

	fmt.Println(eligibleForFreeshipping(true, 80))
	fmt.Println(eligibleForFreeshipping(false, 80))

	fmt.Println(eligibleForBonusPoint(true, 150))
	fmt.Println(eligibleForBonusPoint(false, 150))
	fmt.Println(eligibleForBonusPoint(false, 200))
}
