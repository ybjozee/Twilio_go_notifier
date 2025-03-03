package notifier

import (
	"fmt"

	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type WhatsAppNotifier struct {
}

func (notifier WhatsAppNotifier) Notify(recipientPhoneNumber, message string) (string, error) {
	params := &api.CreateMessageParams{}
	params.SetFrom(fmt.Sprintf("whatsapp:%s", twilioWhatsAppNumber))
	params.SetTo(fmt.Sprintf("whatsapp:%s", recipientPhoneNumber))
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
