package notifier

import (
	"fmt"

	"github.com/oscarboudap/finalert/internal/types"
)

func notifyConsole(event types.Event) error {
	fmt.Printf("🚨 ALERTA CONSOLE: %s %s %.2f ✅ (precio actual: %.2f)\n",
		event.Symbol, event.Operator, event.Value, event.CurrentPrice)
	return nil
}
