package envfmt

import (
	"fmt"
	"os"
	"strings"
)

// ToKeyValueArray returns an array of values formatted as "KEY=VALUE" for each pair in the map.
func ToKeyValueArray(keyValueMap map[string]string) []string {
	keyValueArray := make([]string, 0, len(keyValueMap))
	for key, value := range keyValueMap {
		keyValueArray = append(keyValueArray, fmt.Sprintf("%s=%s", key, value))
	}
	return keyValueArray
}

// ToKeyValueMap parses each string in the array as "KEY=VALUE", and returns a map of key-values.
func ToKeyValueMap(keyValueArray []string) map[string]string {
	keyValueMap := map[string]string{}
	for _, pair := range keyValueArray {
		key, value, ok := strings.Cut(pair, "=")
		if !ok {
			keyValueMap[pair] = ""
			continue
		}
		keyValueMap[key] = value
	}
	return keyValueMap
}

// ReadFile reads and parses the environment file.
// Each line of the file is expected to have the "KEY=VALUE" format.
func ReadFile(envFilePath string) (map[string]string, error) {
	bytes, err := os.ReadFile(envFilePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.ReplaceAll(string(bytes), "\r", ""), "\n")
	return ToKeyValueMap(lines), nil
}
