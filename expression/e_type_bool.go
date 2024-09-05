package expression

import "strconv"

type EBool bool

func (v EBool) ToString() EString {
	return EString(strconv.FormatBool(bool(v)))
}
