package stats

import (
	"github.com/Muhammad-21/bank/pkg/bank/types"
)



func Avg(payments []types.Payment)  types.Money{
	total:=types.Money(0)
	for i:=0; i<len(payments);i++{
		payment:=payments[i]
		total+=payment.Amount
	}
	average:=total/types.Money(len(payments))
	return types.Money(average)
}