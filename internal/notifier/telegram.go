package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	botToken = "7645583213:AAHazR64qZVsfk5tZ5iw8gU-0KviofKiN_g" // <-- aquí pones tu token de Telegram
	chatID   = "1862536733"                                     // <-- aquí pones tu chat ID de Telegram
)

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

// SendTelegramAlert envía un mensaje de alerta a Telegram
func SendTelegramAlert(text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	message := TelegramMessage{
		ChatID: chatID,
		Text:   text,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error serializando mensaje Telegram: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error enviando POST a Telegram: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error respuesta Telegram: %s", resp.Status)
	}

	return nil
}
