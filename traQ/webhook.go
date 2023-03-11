package traq

import (
	"fmt"
	"os"

	traqwriter "github.com/ras0q/traq-writer"
)

var (
	webhookId     = os.Getenv("WEBHOOK_ID")
	webhookSecret = os.Getenv("WEBHOOK_SECRET")
	writer        = traqwriter.NewTraqWebhookWriter(webhookId, webhookSecret, traqwriter.DefaultHTTPOrigin)
)

func PostWebhook(content string) error {
	writer.SetEmbed(false)
	_, err := fmt.Fprint(writer, content)
	return err
}
