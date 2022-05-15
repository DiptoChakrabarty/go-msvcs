package access_token

import (
	"testing"
	"time"
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
		t.Error("new access token should have a user associated")
	}
}

func TestAccessTokenExpiry(t *testing.T) {
	at := AccessToken{}
	if !at.CheckExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if !at.CheckExpired() {
		t.Error("access token should not be expired")
	}
}
