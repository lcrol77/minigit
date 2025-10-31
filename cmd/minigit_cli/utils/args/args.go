package args

import (
	"errors"
	"fmt"
)

type ArgsIterator struct {
	head *Arg
	curr *Arg
}

type Arg struct {
	literal string
	next    *Arg
}

type ArgsError struct {
	Err error
}

func (e *ArgsError) Error() string {
	return fmt.Sprintf("err %v", e.Err)
}

func New(args []string) *ArgsIterator {
	a := ArgsIterator{}
	a.head = &Arg{}
	a.curr = a.head
	for _, arg := range args {
		a.curr.next = &Arg{literal: arg}
		a.curr = a.curr.next
	}
	a.curr = a.head
	a.Next()
	return &a
}

func (a *ArgsIterator) HasNext() bool {
	return a.curr.next != nil
}

func (a *ArgsIterator) Next() (string, error) {
	if a.curr == nil || !a.HasNext() {
		return "", &ArgsError{errors.New("No next argument")}
	}
	a.curr = a.curr.next
	return a.curr.literal, nil
}
