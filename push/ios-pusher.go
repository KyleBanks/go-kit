package push

import (
	"github.com/KyleBanks/go-kit/log"
	"github.com/anachronistic/apns"
)

const (
	// APNS_ENDPOINT_SANDBOX is the endpoint of the sandbox APNS service.
	APNS_ENDPOINT_SANDBOX = "gateway.sandbox.push.apple.com"

	// APNS_ENDPOINT_PRODUCTION is the endpoint of the production APNS service.
	APNS_ENDPOINT_PRODUCTION = "gateway.push.apple.com"

	// APNS_PORT is the port of the APNS service.
	APNS_PORT = 2195
)

// IosPusher supports sending of APNS push notifications to iOS devices.
type IosPusher struct {
	apns *apns.Client
}

// SendMessage sends a JSON payload to the specified DeviceIds through the APNS service.
func (i *IosPusher) SendMessage(message *PushMessage, deviceIds ...string) error {
	for _, deviceId := range deviceIds {
		// Construct the APNS payload...
		payload := apns.NewPayload()
		if len(message.Message) > 0 {
			payload.Alert = message.Message
		}
		if len(message.IosSound) > 0 {
			payload.Sound = message.IosSound
		}

		pn := apns.NewPushNotification()
		pn.DeviceToken = deviceId
		pn.AddPayload(payload)

		for key := range message.Data {
			pn.Set(key, message.Data[key])
		}

		payloadStr, _ := pn.PayloadString()
		log.Infof("Sending iOS Notification: {DeviceId: %v, Payload: %v}", deviceId, payloadStr)

		// Send it off
		res := i.apns.Send(pn)
		if res.Error != nil {
			log.Error("Error sending iOS notification:", res.Error)
			continue
		}
	}

	return nil
}
