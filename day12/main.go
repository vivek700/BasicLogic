package main

import (
	"fmt"
	"strings"
)

func countAppearance(s []string, keyword string) int {
	counter := 0

	for _, v := range s {
		if strings.Contains(v, keyword) {
			counter += 1
		}
	}
	return counter

}

func main() {

	topics := []string{"Economy is booming", "Healthcare debate continues", "Economy vs Healthcare"}
	topics1 := []string{"Politics today", "Sports news", "Weather forecast"}

	topics2 := []string{"Politics today", "Sports news", "Politics and society"}
	topics3 := []string{"Economy update", "Health tips", "Sports highlights"}

	topics4 := []string{"Cricket match analysis", "Football transfer news", "Cricket World Cup"}
	topics5 := []string{"Basketball scores", "Tennis updates", "Olympic games"}

	topics6 := []string{"Health and wellness", "Global economy", "Health tips for winter"}
	topics7 := []string{"Technology trends", "Market analysis", "Education reforms"}

	fmt.Println(countAppearance(topics, "Economy"))
	fmt.Println(countAppearance(topics1, "Economy"))

	fmt.Println(countAppearance(topics2, "Politics"))
	fmt.Println(countAppearance(topics3, "Politics"))

	fmt.Println(countAppearance(topics4, "Cricket"))
	fmt.Println(countAppearance(topics5, "Cricket"))

	fmt.Println(countAppearance(topics6, "Health"))
	fmt.Println(countAppearance(topics7, "Health"))
}
