package jwt_test

import (
	"testing"

	"github.com/anurag925/aero/core"
	"github.com/anurag925/aero/lib/utils/jwt"
	gojwt "github.com/golang-jwt/jwt"
)

func TestEncodeDecode(t *testing.T) {
	core.Test()
	payload := gojwt.MapClaims{"user_id": 123, "username": "test_user"}

	// Encode the payload
	tokenString, err := jwt.Encode(payload)
	if err != nil {
		t.Errorf("Error encoding JWT: %v", err)
	}

	// Decode the token
	decodedPayload, err := jwt.Decode(tokenString)
	if err != nil {
		t.Errorf("Error decoding JWT: %v", err)
	}

	// Check if the decoded payload matches the original payload
	if !claimsMatch(payload, decodedPayload) {
		t.Errorf("Decoded payload does not match original payload")
	}
}

// Helper function to compare two JWT claims maps
func claimsMatch(claims1, claims2 gojwt.MapClaims) bool {
	if len(claims1) != len(claims2) {
		return false
	}
	for key, val1 := range claims1 {
		val2, ok := claims2[key]
		if !ok || val1 != val2 {
			return false
		}
	}
	return true
}
