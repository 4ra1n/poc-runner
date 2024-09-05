package expression

import "strconv"

type EInt int

func (v EInt) ToString() EString {
	return EString(strconv.Itoa(int(v)))
}
