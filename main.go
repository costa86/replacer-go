package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
)

func handleFailure(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}

func readJSONFile(filePath string) (map[string]interface{}, error) {
	file, err := os.ReadFile(filePath)
	handleFailure(err)

	var jsonData map[string]interface{}
	err = json.Unmarshal(file, &jsonData)
	handleFailure(err)

	return jsonData, nil
}

func readFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	handleFailure(err)

	return string(file), nil
}

func replacePlaceholders(content string, jsonData map[string]interface{}) string {
	re := regexp.MustCompile(`{{(.*?)}}`)

	return re.ReplaceAllStringFunc(content, func(match string) string {
		key := re.FindStringSubmatch(match)[1]

		if value, ok := jsonData[key]; ok {
			return fmt.Sprint(value)
		}

		return match
	})
}

func writeFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

const version = "1.0.0"

func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Replacer\nReplaces entries in a file based on a placeholder\nPlaceholder: {{}}\nUsage: ./replacer <json_dict_file> <file_to_modify>\nDeveloped by PCA team\nVersion: %s\n", version)
		os.Exit(1)
	}

	jsonFilePath := os.Args[1]
	jsonData, err := readJSONFile(jsonFilePath)

	handleFailure(err)

	filePath := os.Args[2]
	fileContent, err := readFile(filePath)

	handleFailure(err)

	modifiedContent := replacePlaceholders(fileContent, jsonData)

	err = writeFile(filePath, modifiedContent)
	handleFailure(err)

	fmt.Println("Replacement complete.")
}
