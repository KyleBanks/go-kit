// Package push provides GCM and APN push notification functionality.
package push

import (
	"fmt"
	"github.com/alexjlockwood/gcm"
	"github.com/anachronistic/apns"
)

// Pusher defines the interface for a push notification Pusher.
type Pusher interface {

	// SendMessage sends the payload provided as a JSON string to the deviceIds.
	SendMessage(message *PushMessage, deviceIds string) error
}

// PushMessage defines a message payload to be send to the client.
type PushMessage struct {

	// Message is the title or content to be displayed to the user.
	Message string

	// Data contains any additional data to be sent with the notification.
	Data map[string]interface{}
}

// NewAndroidPusher returns an AndroidPusher, initialized with the specified
// GCM API key.
func NewAndroidPusher(apiKey string) *AndroidPusher {
	return &AndroidPusher{
		gcm: gcm.Sender{
			ApiKey: apiKey,
		},
	}
}

// NewIosPusher returns an IosPusher, initialized with the specified
// certificate and key files.
func NewIosPusher(isProd bool, certificateFile string, keyFile string) *IosPusher {
	endpoint := APNS_ENDPOINT_SANDBOX
	if isProd {
		endpoint = APNS_ENDPOINT_PRODUCTION
	}

	return &IosPusher{
		apns: apns.NewClient(fmt.Sprintf("%v:%v", endpoint, APNS_PORT), certificateFile, keyFile),
	}
}
