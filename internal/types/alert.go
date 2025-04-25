package types

type Alert struct {
	ID        string  // ID único de la alerta
	Symbol    string  // "bitcoin", "ethereum", etc.
	Operator  string  // "<", ">", "=="
	Value     float64 // Umbral de precio
	Triggered bool    // Se activó ya o no
}
