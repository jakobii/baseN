package baseN

import (
	"math"
)

func Encode(n int, charSet []rune) string {

	// use number of characters as base
	var base int = len(charSet)

	// set the array length
	_, _, encodeIndex := getChar(base, n)

	// add 1 to compensate for 0 index
	var reverseRunes = make([]rune, encodeIndex+1)

	// set all characters to 0
	for k, _ := range reverseRunes {
		reverseRunes[k] = 48 // 48 = 0
	}

	// current number
	var cn int = n
	for i := 0; ; i++ {
		charIndex, charValue, encodeIndex := getChar(base, cn)
		if charValue <= 0 {
			break
		}
		cn = cn - charValue
		reverseRunes[encodeIndex] = charSet[charIndex]
	}

	// convert runes to string
	var encoding string
	for i := len(reverseRunes) - 1; i >= 0; i-- {
		encoding += string(reverseRunes[i])
	}
	return encoding
}

func getChar(base int, number int) (int, int, int) {

	if number == 0 {
		return 0, 0, 0
	}

	var encodeIndex int
	var power int
	for i := 0; ; i++ {
		power = int(math.Pow(float64(base), float64(i)))

		//fmt.Println(base, "^", i, "=", power)

		if power > number {
			encodeIndex = i - 1
			break
		}
	}

	if number < base {
		return number, number, encodeIndex
	}

	// (<pow> / <base> = <character domain percentage> ) * <base> = <charSet index>
	charDomains := power / base

	// convert to int64 and throw away percentages to get lowest letter association.
	charSetIndex := int(float64(number) / float64(power) * float64(base))

	// what is the value of the character within pow?
	charSetValue := charDomains * charSetIndex

	return charSetIndex, charSetValue, encodeIndex
}
