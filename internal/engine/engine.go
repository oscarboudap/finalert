package engine

import (
	"fmt"

	"github.com/oscarboudap/finalert/internal/fetcher"
	"github.com/oscarboudap/finalert/internal/types"
)

func Evaluate(alert types.Alert) (bool, error) {
	price, err := fetcher.GetPrice(alert.Symbol)
	if err != nil {
		return false, err
	}

	triggered := false
	switch alert.Operator {
	case ">":
		triggered = price > alert.Value
	case "<":
		triggered = price < alert.Value
	case "==":
		triggered = price == alert.Value
	}

	if triggered && !alert.Triggered {
		fmt.Printf("🚨 ALERTA DISPARADA: %s %s %.2f (precio actual: %.2f)\n", alert.Symbol, alert.Operator, alert.Value, price)
	}

	return triggered, nil
}
