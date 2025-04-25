package types

type Event struct {
	AlertID      string
	Symbol       string
	Operator     string
	Value        float64
	CurrentPrice float64
	TriggeredAt  int64
	Message      string // Nuevo campo para pasar mensajes completos
}
