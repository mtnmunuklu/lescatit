package security

import "golang.org/x/crypto/bcrypt"

// LoadServerTLSCredentials provides encrypt of password.
func EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// LoadServerTLSCredentials provides verify of password.
func VerifyPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
