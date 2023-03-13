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

func PostWebhookOrPrint(content string) error {
	if webhookId == "" {
		fmt.Println(content)
		return nil
	}
	writer.SetEmbed(false)
	_, err := fmt.Fprint(writer, content)
	return err
}
