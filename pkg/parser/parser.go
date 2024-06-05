package parser

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Datastores struct {
	Connections []*Connection `yaml:"datastores"`
}

type Connection struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func Parse(r io.Reader) (map[string]*Connection, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var d Datastores
	err = yaml.Unmarshal(bs, &d)
	if err != nil {
		return nil, err
	}

	if len(d.Connections) == 0 {
		return nil, nil
	}

	cm := make(map[string]*Connection)
	for _, c := range d.Connections {
		cm[c.Name] = c
	}

	return cm, nil
}
