package helper
import (
	"math/rand"
	"strconv"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano()) // Seed only once when the program starts
}

func GenerateNumericString() string {
	length := rand.Intn(63) + 1 // Random length between 1 and 64
	result := ""
	for i := 0; i < length; i++ {
		result += strconv.Itoa(rand.Intn(10)) // Random digit between 0-9
	}
	return result
}

func GenerateRandomString() string {
	length := rand.Intn(63) + 1 // Random length between 1 and 64
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateStringWithBigSize() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 66)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}