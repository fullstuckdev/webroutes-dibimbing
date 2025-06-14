package config

import (
	"os"
	"time"
)

// buat ngambil JWT Secret
func GetJwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

// buat ngambil durasi expired dari JWT
func GetJWTExpirationDuration() time.Duration {
	duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	
	// jadi kalau variabel env JWT_EXPIRES_IN == null
	if err != nil {
		return time.Hour * 24 
	}

	return duration
}

