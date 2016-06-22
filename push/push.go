// Package push provides GCM and APN push notification functionality.
package push

import "github.com/alexjlockwood/gcm"

// Pusher defines the interface for a push notification Pusher.
type Pusher interface {

	// SendMessage sends the payload provided as a JSON string to the deviceIds.
	SendMessage(payload map[string]interface{}, deviceIds string) error
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
