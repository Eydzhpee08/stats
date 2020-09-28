package stats

import (
	"github.com/Eydzhpee08/newbank/pkg/types"
)

func Avg(payments []types.Payment) (money types.Money) {
	k := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			money += payment.Amount
			k++
		}
	}
	return money / types.Money(k)
}

func TotalInCategory(payments []types.Payment, category types.Category) (money types.Money) {
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			money += payment.Amount
		}
	}
	return
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	moneys := map[types.Category]types.Money{}
	quantity := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			moneys[payment.Category] += payment.Amount
			quantity[payment.Category]++
		}
	}

	for key := range moneys {
		moneys[key] /= quantity[key]
	}
	return moneys
}

func PeriodsDynamic(first, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	for key, amount := range second {
		result[key] += amount
	}
	for key, amount := range first {
		result[key] -= amount
	}
	return result
}