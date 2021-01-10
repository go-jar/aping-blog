package utils

import "testing"

func TestJwt(t *testing.T) {
	secretKey := GenSecretKey(6)
	t.Log(secretKey)

	token, err := GenerateToken(33, "felicity", secretKey, 60)
	if err != nil {
		t.Error(err)
	}

	t.Log(token)

	userClaim, err := ParseToken(token, secretKey)
	if err != nil {
		t.Error(err)
	}

	t.Log(userClaim.Id, userClaim.Username)
}
