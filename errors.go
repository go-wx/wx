package wx

import "strings"

type Error struct {
	err     string
	context string
}

func NewWxErr(err string, context string) Error {
	return Error{err, context}
}

func (e Error) Error() string {
	return strings.ToLower(e.context + ": " + e.err)
}
