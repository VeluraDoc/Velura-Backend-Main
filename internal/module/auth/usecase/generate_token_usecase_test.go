package auth_usecase

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken_Success(t *testing.T) {
	const secret = "testsecret"
	const expirationHours = "1"
	if err := os.Setenv("JWT_SECRET", secret); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}
	if err := os.Setenv("JWT_EXPIRATION", expirationHours); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}

	t.Cleanup(func() {
		os.Unsetenv("JWT_SECRET")
		os.Unsetenv("JWT_EXPIRATION")
	})

	userID := "u-123"

	tokenStr, err := GenerateToken(userID)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}
	if tokenStr == "" {
		t.Fatalf("expected non-empty token string")
	}

	parsed, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.Fatalf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		t.Fatalf("failed to parse/validate token: %v", err)
	}
	if !parsed.Valid {
		t.Fatalf("token is not valid")
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("claims is not MapClaims")
	}

	if gotID, _ := claims["id"].(string); gotID != userID {
		t.Fatalf("id claim mismatch: got %q want %q", gotID, userID)
	}
}
