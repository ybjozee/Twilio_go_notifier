package notifier

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
)

var client = twilio.NewRestClient()
var twilioPhoneNumber = os.Getenv("TWILIO_PHONE_NUMBER")
var twilioWhatsAppNumber = os.Getenv("TWILIO_WHATSAPP_NUMBER")

type Notifier interface {
	Notify(recipientPhoneNumber, message string) (string, error)
}

func GetNotifier(medium string) (Notifier, error) {
	switch medium {
	case "sms":
		return SMSNotifier{}, nil
	case "voice":
		return VoiceNotifier{}, nil
	case "whatsApp":
		return WhatsAppNotifier{}, nil
	default:
		return nil, fmt.Errorf("unsupported medium %s provided", medium)
	}
}

func GetSupportedMedia() []string {
	return []string{
		"sms",
		"voice",
		"whatsApp",
	}
}
