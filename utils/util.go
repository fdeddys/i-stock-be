package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset ...
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateRandomChar ...
func GenerateRandomChar() string {
	return StringWithCharset(8, charset)
}

// HashPassword ...
func HashPassword(pass string) string {

	hashPassword := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(hashPassword[:])

}

func ConvertDateFormat(reqDate string) string {
	reqDateLayout := "02-01-2006"
	dateLayout := "2006-01-02"

	date, _ := time.Parse(reqDateLayout, reqDate)

	return date.Format(dateLayout)
}

// ConvertTime ...
func ConvertTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")

}

func ConvertStringToTime(reqDate string) time.Time {
	reqDateLayout := "02-01-2006"

	date, _ := time.Parse(reqDateLayout, reqDate)
	return date
}


