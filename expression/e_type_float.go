package expression

import "strconv"

type EFloat float64

func (v EFloat) ToString() EString {
	return EString(strconv.FormatFloat(float64(v), 'f', -1, 64))
}
