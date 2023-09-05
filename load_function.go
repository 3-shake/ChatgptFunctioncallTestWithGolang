package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func LoadFunctionDefinitions(directory string) ([]openai.FunctionDefinition, error) {
	var funcDefs []openai.FunctionDefinition

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("Error reading directory: %v", err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			data, err := ioutil.ReadFile(directory + "/" + f.Name())
			if err != nil {
				return nil, fmt.Errorf("Error reading file %s: %v", f.Name(), err)
			}

			// fmt.Printf("Loaded JSON from file %s: %s\n", f.Name(), string(data))

			var defs []openai.FunctionDefinition
			err = json.Unmarshal(data, &defs)
			if err != nil {
				return nil, fmt.Errorf("Error unmarshalling function definitions from file %s: %v", f.Name(), err)
			}

			funcDefs = append(funcDefs, defs...)
		}
	}
	// fmt.Printf("Loaded JSON: \n%s\n", funcDefs)

	return funcDefs, nil
}
