package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CalculatePoints calculates the points for a receipt
func CalculatePoints(receipt *Receipt) int {
	points := 0

	// Points for alphanumeric characters in retailer name
	points += len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1))

	// Points if total is a round dollar amount
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// Points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Points for item descriptions with length multiple of 3
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Points if the day is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// Points if purchase time is between 2:00pm and 4:00pm
	time, _ := time.Parse("15:04", receipt.PurchaseTime)
	if time.Hour() == 14 || (time.Hour() == 15 && time.Minute() < 60) {
		points += 10
	}

	return points
}
