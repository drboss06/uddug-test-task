package testTask

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

var testData = []*Transaction{
	{
		4456,
		time.Unix(1616026248, 0),
	},
	{
		4231,
		time.Unix(1616022648, 0),
	},
	{
		5212,
		time.Unix(1616019048, 0),
	},
	{
		4321,
		time.Unix(1615889448, 0),
	},
	{
		4567,
		time.Unix(1615871448, 0),
	},
}

var exceptedData = []*Transaction{
	{
		4456,
		time.Unix(1616025600, 0),
	},
	{
		4231,
		time.Unix(1615939200, 0),
	},
	{
		4321,
		time.Unix(1615852800, 0),
	},
}

func TestFormatTransactions(t *testing.T) {
	var result = FormatTransactions(testData, time.Hour*24)

	var equalRes = slices.EqualFunc(result, exceptedData, func(e1 *Transaction, e2 *Transaction) bool {
		if e1.Timestamp == e2.Timestamp && e1.Value == e2.Value {
			return true
		} else {
			return false
		}
	})

	if !equalRes {
		fmt.Println("Resault")
		for _, i := range result {
			fmt.Println(i)
		}
		fmt.Println("ExceptedData")

		for _, i := range exceptedData {

			fmt.Println(i)
		}

		t.Fatal("Filed test")
	}
}
