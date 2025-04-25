package persistence

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite" // <<< Cambiado aquí
)

var db *sql.DB

// Inicializar base de datos
func InitDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath) // <<< Ojo: ahora es "sqlite"
	if err != nil {
		return fmt.Errorf("error abriendo SQLite: %w", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS predictions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		symbol TEXT,
		timestamp DATETIME,
		price REAL,
		delta_real REAL,
		delta_predicted REAL,
		direction_correct BOOLEAN
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creando tabla: %w", err)
	}

	return nil
}

// Guardar predicción
func SavePrediction(symbol string, price float32, deltaReal float32, deltaPredicted float32) error {
	directionCorrect := (deltaReal > 0 && deltaPredicted > 0) || (deltaReal < 0 && deltaPredicted < 0)

	_, err := db.Exec(`
		INSERT INTO predictions (symbol, timestamp, price, delta_real, delta_predicted, direction_correct)
		VALUES (?, ?, ?, ?, ?, ?)
	`, symbol, time.Now(), price, deltaReal, deltaPredicted, directionCorrect)
	if err != nil {
		return fmt.Errorf("error guardando predicción: %w", err)
	}

	return nil
}
