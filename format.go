package testTask

import (
	"sort"
	"time"
)

type Transaction struct {
	Value     int
	Timestamp time.Time
}

func FormatTransactions(transactions []*Transaction, interval time.Duration) []*Transaction {
	intervalMap := make(map[int]Transaction)

	for _, transaction := range transactions {
		unixSec := transaction.Timestamp.Unix()
		intervalDuration := int(unixSec - unixSec%int64(interval.Seconds()))

		if intervalMap[intervalDuration].Timestamp.Compare(transaction.Timestamp) == -1 {
			intervalMap[intervalDuration] = *transaction
		}
	}

	var result []*Transaction

	for key, value := range intervalMap {
		tr := &Transaction{
			Value:     value.Value,
			Timestamp: time.Unix(int64(key), 0),
		}

		result = append(result, tr)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Timestamp.Compare(result[j].Timestamp) == 1 {
			return true
		} else {
			return false
		}
	})

	return result
}
