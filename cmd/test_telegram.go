package main

import (
	"log"

	"github.com/oscarboudap/finalert/internal/notifier"
)

func main() {
	err := notifier.SendTelegramAlert("🚀 Test desde Go con éxito ✅")
	if err != nil {
		log.Fatal(err)
	}
}
