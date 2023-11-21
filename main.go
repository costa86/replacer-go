package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <jsonFilePath> <fileToModify>")
		os.Exit(1)
	}

	// Read JSON file
	jsonFilePath := os.Args[1]
	jsonData, err := readJSONFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		os.Exit(1)
	}

	// Read the content of the file
	filePath := os.Args[2]
	fileContent, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		os.Exit(1)
	}

	// Replace placeholders in the file with JSON values
	modifiedContent := replacePlaceholders(fileContent, jsonData)

	// Write the modified content back to the file
	err = writeFile(filePath, modifiedContent)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		os.Exit(1)
	}

	fmt.Println("Replacement complete.")
}

func readJSONFile(filePath string) (map[string]interface{}, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func readFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func replacePlaceholders(content string, jsonData map[string]interface{}) string {
	re := regexp.MustCompile(`{{(.*?)}}`)

	return re.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the key from the placeholder
		key := re.FindStringSubmatch(match)[1]

		// Look up the key in the JSON data
		if value, ok := jsonData[key]; ok {
			return fmt.Sprint(value)
		}

		// If the key is not found, return the original placeholder
		return match
	})
}

func writeFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
