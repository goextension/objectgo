package String

import (
	"github.com/goextension/objectgo"
)

type String = objectgo.String

func New(s string) String {
	return String(s)
}
