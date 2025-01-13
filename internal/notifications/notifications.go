package notifications

import (
	"jnsltk/look_away/internal/config"

	"github.com/gen2brain/beeep"
)

type Notifier struct {
	config config.NotificationConfig
}

func NewNotifier(cfg config.NotificationConfig) *Notifier {
	return &Notifier{config: cfg}
}

func (n *Notifier) Notify(message string) {
	if n.config.UseAlert {
		beeep.Alert("look_away reminder", message, "")
	} else {
		beeep.Notify("look_away reminder", message, "")
	}
}