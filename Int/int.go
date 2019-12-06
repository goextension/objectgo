package Int

import "github.com/goextension/objectgo"

func New(i int) objectgo.Int {
	return objectgo.Int(i)

}
func toString(i objectgo.Int) objectgo.String {
	return i.toString()
}
