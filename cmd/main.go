package main

import (
	"fmt"
	"time"

	"github.com/oscarboudap/finalert/internal/engine"
	"github.com/oscarboudap/finalert/internal/types"
)

func main() {
	alerts := []types.Alert{
		{ID: "1", Symbol: "bitcoin", Operator: "<", Value: 30000, Triggered: false},
		{ID: "2", Symbol: "ethereum", Operator: ">", Value: 2000, Triggered: false},
	}

	for {
		for i, alert := range alerts {
			triggered, err := engine.Evaluate(alert)
			if err != nil {
				fmt.Println("Error evaluando alerta:", err)
				continue
			}
			if triggered {
				alerts[i].Triggered = true
			}
		}
		time.Sleep(10 * time.Second)
	}
}
