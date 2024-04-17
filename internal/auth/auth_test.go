package auth

import (
	"os"
	"testing"
)

func TestJWT(t *testing.T) {

	// Define tests
	tests := map[string]struct {
		JwtSecret string
		UserId    int
		Email     string
	}{
		"env test": {"test", 1, "test@gmail.com"},
		"env 123":  {"123", 2, "123@gmail.com"},
	}

	// Run tests
	for i, v := range tests {
		t.Run(i, func(t *testing.T) {

			// Set env
			err := os.Setenv("JWT_SECRET", v.JwtSecret)
			if err != nil {
				t.Error(err)
			}

			jwtStr, err := GenerateJWT(v.UserId, v.Email)
			if err != nil {
				t.Error(err)
			}

			userInfo, err := ParseJWT(jwtStr)
			if err != nil {
				t.Error(err)
			}

			// Check for UserInfo mismatch
			if v.Email != userInfo.Email {
				t.Errorf("Email: generated %s, parsed %s", v.Email, userInfo.Email)
			}
			if v.UserId != userInfo.UserId {
				t.Errorf("UserId: generated %d, parsed %d", v.UserId, userInfo.UserId)
			}

		})
	}
}
