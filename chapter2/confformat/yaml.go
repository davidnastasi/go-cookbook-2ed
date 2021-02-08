package confformat

import (
	"bytes"

	"github.com/ghodss/yaml"
)

// YAMLData is our common data struct with TOML struct tags
type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

// ToYAML dumps the YAMLData struct to a TOML format bytes.Buffer
func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode will decode into YAMLData
func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}
