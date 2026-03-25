package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// Config holds application configuration
type Config struct {
	WorkerCount int    `json:"workerCount"`
	QueueName   string `json:"queueName"`
	Timeout     int    `json:"timeout"`
}

// LoadConfig reads and parses the configuration file
func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	if config.WorkerCount <= 0 {
		return nil, errors.New("worker count must be greater than 0")
	}

	return &config, nil
}

// Retry executes a function with retry logic
func Retry(attempts int, delay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return err
}

// GenerateID creates a unique identifier
func GenerateID() string {
	return time.Now().Format("20060102150405") + "-" + RandomString(8)
}

// RandomString generates a random string of given length
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}