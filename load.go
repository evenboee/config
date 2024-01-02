package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func loadFiles(filenames ...string) (map[string]string, error) {
	vals := make(map[string]string)
	for _, filename := range filenames {
		val, err := loadFile(filename)
		if err != nil {
			return nil, err
		}
		for k, v := range val {
			vals[k] = v
		}
	}
	return vals, nil
}

func loadFile(filename string) (map[string]string, error) {
	ext := filepath.Ext(filename)
	switch ext {
	case ".yaml", ".yml":
		return loadYAMLFile(filename)
	case ".json":
		return loadJSONFile(filename)
	default:
		return loadEnvFile(filename)
	}
}

func loadYAMLFile(filename string) (map[string]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	vals := make(map[string]any)
	err = yaml.Unmarshal(content, &vals)
	if err != nil {
		return nil, err
	}

	return json2env(vals), nil
}

func loadJSONFile(filename string) (map[string]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	vals := make(map[string]any)
	err = json.Unmarshal(content, &vals)
	if err != nil {
		return nil, err
	}

	return json2env(vals), nil
}

func loadEnvFile(filename string) (map[string]string, error) {
	return godotenv.Read(filename)
}

// json2env converts a map[string]any to a map[string]string.
// flattens the json structure into a single level.
// also works for yaml.
func json2env(data map[string]any) map[string]string {
	flattened := make(map[string]string)

	for key, value := range data {
		switch val := value.(type) {
		case map[string]any:
			for k, v := range json2env(val) {
				flattened[key+"_"+k] = v
			}
		case any:
			flattened[key] = fmt.Sprintf("%v", val)
		}
	}

	return flattened
}
