// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package selection

import (
	time "time"
)

type Event interface{}
type Like struct {
	Reaction  string    `json:"reaction"`
	Sent      time.Time `json:"sent"`
	Selection []string  `json:"selection"`
	Collected []string  `json:"collected"`
}
type Post struct {
	Message   string    `json:"message"`
	Sent      time.Time `json:"sent"`
	Selection []string  `json:"selection"`
	Collected []string  `json:"collected"`
}
