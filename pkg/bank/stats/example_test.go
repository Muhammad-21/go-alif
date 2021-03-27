package stats

import (
	"fmt"

	"github.com/Muhammad-21/bank/pkg/bank/types")


func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 10_000,
		},
		{
			Amount: 20_000,
		},
	}
	average:=Avg(payments)
	fmt.Println(average)
//Output:
//15000
}