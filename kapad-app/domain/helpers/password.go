package helpers

import "golang.org/x/crypto/bcrypt"

func HashValueUsingBcrypt(value string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashed)
}

func IsHashedAndPlanStringEqual(hashed string, plan string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plan))
	if err != nil {
		return false
	}
	return true
}
