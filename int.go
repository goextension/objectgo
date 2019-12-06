package objectgo

import "strconv"

func (i Int) String() string {
	return strconv.Itoa(int(i))
}
