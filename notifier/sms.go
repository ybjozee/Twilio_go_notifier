package notifier

import (
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSNotifier struct {
}

func (notifier SMSNotifier) Notify(recipientNumber, message string) (string, error) {
	params := &api.CreateMessageParams{}
	params.SetBody(message)
	params.SetFrom(phoneNumber)
	params.SetTo(recipientNumber)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
