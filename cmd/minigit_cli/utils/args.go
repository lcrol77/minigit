package utils

import (
	"errors"
	"fmt"
)

type ArgsIterator struct {
	idx  int
	args []string
}

type ArgsError struct {
	Err error
}

func (e *ArgsError) Error() string {
	return fmt.Sprintf("err %v", e.Err)
}

func New(args []string) *ArgsIterator {
	ai := ArgsIterator{}
	ai.idx = 0
	ai.args = args
	ai.Next()
	return &ai
}

func (a *ArgsIterator) Next() (string, error) {
	if a.idx == len(a.args) {
		return "", &ArgsError{Err: errors.New("next index is undefined")}
	}
	a.idx++
	return a.args[a.idx-1], nil
}
