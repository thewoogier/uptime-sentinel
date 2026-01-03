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
func LoadTargets(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	// TODO: We're assuming the JSON structure is just an array of strings?
	// Or maybe it expects an object? Let's assume object with "targets" key based on struct above,
	// but let's write code that tries to parse it into the struct.
	var cfg Config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	return cfg.Targets, nil
}
