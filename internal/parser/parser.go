package parser

import (
	"encoding/json"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func Parse(path string) (FileOpenAPI, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return FileOpenAPI{}, err
	}
	if strings.HasSuffix(path, "yaml") || strings.HasSuffix(path, "yml") {
		return FileOpenAPI{
			Filetype: "yaml",
			Spec:     parseYaml(file),
		}, nil
	} else if strings.HasSuffix(path, "json") {
		return FileOpenAPI{
			Filetype: "json",
			Spec:     parseJson(file),
		}, nil
	} else {
		return FileOpenAPI{}, ErrWrongFileFormat
	}
}

func parseYaml(file []byte) OpenAPI {
	o := OpenAPI{}
	err := yaml.Unmarshal(file, &o)
	if err != nil {
		panic(err)
	}
	return o
}

func parseJson(file []byte) OpenAPI {
	o := OpenAPI{}
	err := json.Unmarshal(file, &o)
	if err != nil {
		panic(err)
	}
	return o
}
