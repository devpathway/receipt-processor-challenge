package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CalculatePoints calculates the points for a given receipt
func CalculatePoints(receipt *Receipt) int {
	points := 0

	// Rule 1: Alphanumeric characters in retailer name
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alphanumeric.FindAllString(receipt.Retailer, -1))

	// Rule 2: 50 points if total is a round dollar amount (e.g., 10.00)
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0 // Invalid total, no points
	}
	if total == math.Floor(total) {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Item descriptions with length multiple of 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2)) // 20% of item price, rounded up
			}
		}
	}

	// Rule 6: 6 points if the day is odd
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && date.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if time is between 2:00 PM and 4:00 PM
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
