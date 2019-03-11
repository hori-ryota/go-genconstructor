package example

import "time"

//go:generate go-genconstructor

//genconstructor
type Person struct {
	id        string `required:""`
	name      string `required:""`
	tags      []string
	createdAt time.Time `required:"time.Now()"`
}

//genconstructor -p
type PersonService struct {
	id string `required:""`
}
