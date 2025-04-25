package engine

import (
	"fmt"

	"github.com/oscarboudap/finalert/internal/fetcher"
	"github.com/oscarboudap/finalert/internal/mlmodel"
	"github.com/oscarboudap/finalert/internal/notifier"
	"github.com/oscarboudap/finalert/internal/persistence"
	"github.com/oscarboudap/finalert/internal/types"
)

func Evaluate(alert types.Alert) (*types.Alert, error) {
	price, err := fetcher.GetPrice(alert.Symbol)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo precio: %w", err)
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

	// Si alerta clásica disparada
	if triggered && !alert.Triggered {
		alert.Triggered = true

		message := fmt.Sprintf(
			"🚨 ALERTA DISPARADA\n%s %s %.2f ✅\nPrecio Actual: %.2f",
			alert.Symbol, alert.Operator, alert.Value, price,
		)

		if alert.Channel == "telegram" {
			notifier.SendTelegramAlert(message)
		} else {
			fmt.Println(message)
		}

		return &alert, nil
	}

	// -------- Predicción IA (solo si no está disparada) --------

	// Fake delta (porque no hemos implementado un sistema real aún)
	delta1 := 0.0

	// Pedir predicción
	predictedDelta, err := mlmodel.Predict([]float32{float32(price), float32(delta1), 0, 0})
	if err != nil {
		return nil, fmt.Errorf("error haciendo predicción IA: %w", err)
	}

	// Enviar alerta predictiva
	predictiveAlert := fmt.Sprintf(
		"🚀 Predicción %s\nPrecio Actual: %.2f USD\nÚltima hora: %.2f%% 📊\nPróxima hora (predicción): %.2f%% 📈",
		alert.Symbol,
		price,
		delta1,
		predictedDelta,
	)

	if alert.Channel == "telegram" {
		notifier.SendTelegramAlert(predictiveAlert)
	} else {
		fmt.Println(predictiveAlert)
	}

	// Guardar predicción
	err = persistence.SavePrediction(alert.Symbol, float32(price), float32(delta1), predictedDelta)
	if err != nil {
		fmt.Println("Error guardando predicción:", err)
	}

	// Mini-reporte
	miniReport := fmt.Sprintf(
		"📊 Reporte FinAlert++\nActivo: %s\nReal última hora: %.2f%%\nPredicción próxima hora: %.2f%%\n",
		alert.Symbol,
		delta1,
		predictedDelta,
	)

	if alert.Channel == "telegram" {
		notifier.SendTelegramAlert(miniReport)
	} else {
		fmt.Println(miniReport)
	}

	return nil, nil
}
