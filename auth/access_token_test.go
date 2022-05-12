package auth

import (
	"testing"
)

func TestNewAccessToken(t *testing.T) {
	at := GenerateAccessToken()
	if at.CheckExpired() {
		t.Error("new access token should be generated")
	}

	if at.AccessToken != "" {
		t.Error("new access token should be used")
	}

	if at.UserId != 0 {
		t.Error("new access token should be used")
	}
}
