package main

import (
	"fmt"
	"time"

	"github.com/oscarboudap/finalert/internal/engine"
	"github.com/oscarboudap/finalert/internal/persistence"
	"github.com/oscarboudap/finalert/internal/types"
)

func main() {
	// Iniciar base de datos SQLite
	err := persistence.InitDB("finalert.db")
	if err != nil {
		fmt.Println("Error iniciando base de datos:", err)
		return
	}

	// Definir alertas
	alerts := []types.Alert{
		{ID: "1", Symbol: "bitcoin", Operator: ">", Value: 1000, Triggered: false, Channel: "telegram"},
		{ID: "2", Symbol: "ethereum", Operator: "<", Value: 3000, Triggered: false, Channel: "console"},
	}

	// Bucle principal
	for {
		for i, alert := range alerts {
			event, err := engine.Evaluate(alert)
			if err != nil {
				fmt.Println("Error evaluando alerta:", err)
				continue
			}
			if event != nil {
				alerts[i].Triggered = true
			}
		}
		time.Sleep(60 * time.Second) // Cada 10 segundos reevalúa
	}
}
