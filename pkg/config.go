package pkg

import "os"

// Returns a mockable JSON File location
type ConfigProvider interface {
	GetResourcePath() string
}

// Retrieve the configuration from the environment variables
type EnvReader struct{}

func NewEnvReader() EnvReader {
	return EnvReader{}
}

func (EnvReader) GetResourcePath() string {
	return os.Getenv("RESOURCES")
}
