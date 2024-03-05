package notification

import (
	"strings"
	"testing"
)

func TestGetRandomCatMessage(t *testing.T) {
	message := getRandomCatMessage()
	if message == "" {
		t.Errorf("Expected a non-empty string, got an empty string")
	}
	if !strings.Contains(message, "にゃん") {
		t.Errorf("Expected the message to contain 'にゃん', but it didn't. Message: %s", message)
	}
}
