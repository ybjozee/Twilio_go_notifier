package notifier

import (
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/twilio/twilio-go/twiml"
)

type VoiceNotifier struct {
}

func (notifier VoiceNotifier) Notify(recipientPhoneNumber, message string) (string, error) {
	params := &api.CreateCallParams{}
	params.SetFrom(twilioPhoneNumber)
	params.SetTo(recipientPhoneNumber)

	sayMessage, err := sayTwiml(message)
	if err != nil {
		return "", err
	}
	params.SetTwiml(sayMessage)

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func sayTwiml(message string) (string, error) {
	say := &twiml.VoiceSay{
		Message: message,
	}

	verbList := []twiml.Element{say}
	sayTwiml, err := twiml.Voice(verbList)

	if err != nil {
		return "", err
	}

	return sayTwiml, nil
}
