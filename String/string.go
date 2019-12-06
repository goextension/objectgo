package String

import (
	"github.com/goextension/objectgo"
)

func New(s string) objectgo.String {
	return String(s)
}

func String(s string) objectgo.String {
	return objectgo.String(s)
}
