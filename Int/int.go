package Int

import "github.com/goextension/objectgo"

type Int = objectgo.Int

func New(i int) Int {
	return Int(i)
}
