package helpers

import (
	"log"
	"time"
)

func DateFormatter(date string, fromLayout string, toLayout string) string {
	parsedDate, err := time.Parse(fromLayout, date)
	if err != nil {
		log.Printf("fail to parse date %s: %v", date, err)
		return date
	}

	return parsedDate.Format(toLayout)
}

func Min(x, y uint64) uint64 {
	if x > y {
		return y
	}
	return x
}
