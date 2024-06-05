package main

type Datastore interface {
	Insert(query string, args ...any) error
	Query(query string, args ...any) (any, error)
}
