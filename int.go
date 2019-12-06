package objectgo

import "strconv"

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i Int) ToString() String {
	return String(strconv.Itoa(int(i)))
}
