package parser

import "fmt"

type Datastore interface {
	Insert(query string, args ...any) error
	Query(query string, args ...any) (any, error)
}

type NullDatastore struct{}

func (d *NullDatastore) Insert(query string, args ...any) error {
	return fmt.Errorf("NullDatastore")
}

func (d *NullDatastore) Query(query string, args ...any) (any, error) {
	return nil, fmt.Errorf("NullDatastore")
}
