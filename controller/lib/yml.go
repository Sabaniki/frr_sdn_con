package controller

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Host struct {
	Name string `yaml:"name"`
}

func DeserializeHosts(filePath string) ([]Host, error) {
	var yml []Host
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return yml, err
	}
	if err := yaml.Unmarshal(raw, &yml); err != nil {
		return yml, err
	}
	return yml, nil
}
