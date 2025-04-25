package mlmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// PredictFeatures representa el input para el microservicio FastAPI
type PredictFeatures struct {
	Features []float32 `json:"features"`
}

// Predict llama al microservicio de IA para obtener la predicción
func Predict(features []float32) (float32, error) {
	payload := PredictFeatures{Features: features}

	body, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("error serializando payload: %w", err)
	}

	resp, err := http.Post("http://localhost:8000/predict", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, fmt.Errorf("error llamando microservicio IA: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Prediction float32 `json:"prediction"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return 0, fmt.Errorf("error decodificando respuesta IA: %w", err)
	}

	return response.Prediction, nil
}
