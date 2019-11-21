package baseN

import (
	"errors"
	"fmt"
)

// Base stores calculated properties based used to convert bases
type Base struct {
	charset []byte
	powers  []uint64
}

const maxPower uint64 = 18446744073709551615

// New is like NewBase except it panics if there is an error
func New(charSet []rune) Base {
	b, err := NewBase(charSet)
	if err != nil {
		panic(err)
	}
	return b
}

// NewBase creates a new base encoder/decoder from the callers specified character set.
// error is returned when the provided character set is not a distinct list of characters.
func NewBase(charSet []rune) (Base, error) {
	var b Base

	// byte rune array into base.charset
	//check chars for duplicates in the character set.
	b.charset = make([]byte, 0, len(charSet))
	for _, c := range charSet {
		if _, ok := b.Index(byte(c)); !ok {
			b.charset = append(b.charset, byte(c))
		} else {
			return b, errors.New(fmt.Sprintf("invalid character set. duplicate character '%s' in character set.", string(c)))
		}
	}

	// create powers table.
	b.powers = make([]uint64, 0, 0)
	power := b.N()
	var p uint64
	for i := 0; p <= maxPower; i++ {
		p = pow(power, uint(i))
		if p == 0 && i > 0 {
			break
		}
		b.powers = append(b.powers, p)
	}
	return b, nil
}

// Encode will convert the provided uint64 into a encoded byte slice.
func (b *Base) Encode(source uint64) []byte {
	var e []byte
	var rem uint64 = source
	var char uint64
	var init bool
	var reset bool
	var base uint64 = uint64(b.N())
	for i := 0; rem > 0; reset = false {
		if (*b).powers[i] > rem {
			if !init {
				init = true

				// generate char array
				e = make([]byte, i+1)

				// zero fill
				for j := 0; j < i; j++ {
					e[j] = (*b).charset[0]
				}
			}
			// n / l
			if i-1 < 0 {
				char = rem
			} else {
				char = rem / (*b).powers[i-1]
			}
			e[i] = (*b).charset[char]

			// (high/base)*char
			// this should never result in a float
			sector := ((*b).powers[i] / base) * char
			rem = rem - sector

			// start next search at zero
			reset = true
			i = 0
		}
		if !reset {
			i++
		}
	}

	// reverse the byte array
	// to read from left to right
	i := 0
	j := len(e) - 1
	for i <= j {
		e[i], e[j] = e[j], e[i]
		i++
		j--
	}

	return e
}

// Decode takes a base encoded byte slice and converts it back to a uint64 number.
func (b *Base) Decode(source []byte) (uint64, error) {
	var total uint64
	var base uint64
	var high uint64
	var sectorSize uint64

	// reverse the source array to read from right to left
	var e []byte = make([]byte, 0, len(source))
	e = append(e, source...)
	i := 0
	j := len(e) - 1
	for i <= j {
		e[i], e[j] = e[j], e[i]
		i++
		j--
	}

	// start from the last index.
	// last index identifies its power.
	//   e.g. base 4: 255 = 3333
	//       4 3 2 1 <-
	//   | 1 _ _ _ _
	//   v 2
	//     3
	//     4
	//
	base = uint64(b.N())
	for i := len(e) - 1; i >= 0; i-- {
		high = b.powers[i]
		sectorSize = high / base
		if charIndex, ok := b.Index(e[i]); ok {
			// no need to add zero
			if charIndex == 0 {
				continue
			}
			total += sectorSize * uint64(charIndex)
		} else {
			return 0, errors.New(fmt.Sprintf("invalid character '%s'.", string(e[i])))
		}
	}
	return total, nil
}

func (b *Base) N() uint {
	return uint(len((*b).charset))
}

func (b *Base) Index(B byte) (index uint, ok bool) {
	for index, char := range (*b).charset {
		if char == B {
			return uint(index), true
		}
	}
	return 0, false
}

// x^y
func pow(n uint, power uint) uint64 {
	var t uint64 = uint64(n)
	var i uint
	for i < power {
		t *= uint64(n)
		i++
	}
	return t
}
