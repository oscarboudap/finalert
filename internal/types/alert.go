package types

type Alert struct {
	ID        string  `json:"id"`
	Symbol    string  `json:"symbol"`
	Operator  string  `json:"operator"` // "<", ">", "=="
	Value     float64 `json:"value"`
	Triggered bool    `json:"triggered"`
	Channel   string  `json:"channel"` // "console", "telegram", "email", etc.
}
