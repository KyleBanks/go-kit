package push

import (
	"github.com/KyleBanks/go-kit/log"
	"github.com/alexjlockwood/gcm"
)

const (
	// MAX_RETRIES defines the number of times to retry message sender if an error occurs.
	MAX_RETRIES = 2
)

// AndroidPusher supports sending of GCM push notifications to Android devices.
type AndroidPusher struct {
	gcm gcm.Sender
}

// SendMessage sends a JSON payload to the specified DeviceIds through the GCM service.
func (a *AndroidPusher) SendMessage(message *PushMessage, deviceIds ...string) error {
	notif := map[string]interface{}{
		"data": message.Data,
	}
	if len(message.Message) > 0 {
		notif["message"] = message.Message
	}

	msg := gcm.NewMessage(notif, deviceIds...)
	log.Infof("Sending Android Notification: {DeviceIds: %v, Payload: %v}", deviceIds, msg)

	if _, err := a.gcm.Send(msg, MAX_RETRIES); err != nil {
		return err
	}

	return nil
}
