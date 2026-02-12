package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	filename := "./data.yaml"

	// Open the YAML file
	file, _ := os.Open(filename)
	defer file.Close()

	// Read the file contents
	data, err := io.ReadAll(file)

	//parse the YAML content
	var content any
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing YAML file %s: %v\n", filename, err)
		os.Exit(1)
	}

	if m, ok := content.(map[string]interface{}); ok {
		if name, exists := m["lang"]; exists {
			fmt.Println(name)
		} else {
			fmt.Println("The 'lang' field does not exist in the YAML file.")
		}
	} else {
		fmt.Println("The YAML content is not a valid map structure.")
	}
}
