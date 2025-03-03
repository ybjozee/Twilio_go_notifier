package notifier

import (
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSNotifier struct {
}

func (notifier SMSNotifier) Notify(recipientPhoneNumber, message string) (string, error) {
	params := &api.CreateMessageParams{}
	params.SetBody(message)
	params.SetFrom(twilioPhoneNumber)
	params.SetTo(recipientPhoneNumber)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
