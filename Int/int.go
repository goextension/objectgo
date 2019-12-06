package Int

import "github.com/goextension/objectgo"

func New(i int) objectgo.Int {
	return objectgo.Int(i)

}
func ToString(i objectgo.Int) objectgo.String {
	return i.ToString()
}
