package app

import "gopkg.in/mgo.v2"

// IPersistance - interface for user settings and general purpose storage
type IPersistance interface {
	GetCollection(string) *mgo.Collection
	DropDatabase() error
}
