package objectgo

import "strconv"

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i Int) toString() String {
	return String(strconv.Itoa(int(i)))
}
