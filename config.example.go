// Package main provides a sample configuration structure for the POC application
package main

// Config holds application configuration
type Config struct {
	Server   ServerConfig   `json:"server"`
	Greeting GreetingConfig `json:"greeting"`
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// GreetingConfig holds greeting-related configuration
type GreetingConfig struct {
	DefaultName string `json:"default_name"`
}

// Example usage (commented out to keep main.go simple):
/*
func loadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host: "localhost",
			Port: "8080",
		},
		Greeting: GreetingConfig{
			DefaultName: "Guest",
		},
	}
}
*/
