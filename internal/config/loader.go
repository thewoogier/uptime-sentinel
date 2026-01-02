package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Targets []string `json:"targets"`
}

// LoadTargets reads the targets.json file
// This is the source of the "False Ambiguity".
// The error handling is vague, but the requirement is in the README.
func LoadTargets(path string) ([]string, error) {
	// Hardcoded fallback check? No, let's just fail if file doesn't exist.
	file, err := os.Open(path)
	if err != nil {
		// Vague error message that might confuse the user if they don't read docs
		return nil, err 
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	
	// INTENTIONAL GAP: Assuming the JSON structure is just an array of strings? 
	// Or maybe it expects an object? Let's assume object with "targets" key based on struct above,
	// but let's write code that tries to parse it into the struct.
	var cfg Config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	return cfg.Targets, nil
}
