// Code generated by go-genconstructor; DO NOT EDIT.

package example

import (
	"time"
)

func NewPerson(
	id string,
	name string,
) Person {
	return Person{
		id:        id,
		name:      name,
		createdAt: time.Now(),
	}
}
