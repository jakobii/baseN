package baseN

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var e []byte

	// base 4
	b4, _ := NewBase([]rune("0123"))
	e = b4.Encode(255)
	if string(e) != "3333" {
		t.Log(string(e))
		t.Error("got:", string(e), ", want: 3333")
	}

	e = b4.Encode(1340583)
	if string(e) != "11013102213" {
		t.Log(string(e))
		t.Error("got:", string(e), ", want: 11013102213")
	}

	// base 16
	b16, _ := NewBase([]rune("0123456789abcdef"))
	e = b16.Encode(255)
	if string(e) != "ff" {
		t.Log(string(e))
		t.Error("got:", string(e), ", want: ff")
	}

	e = b16.Encode(1340583)
	if string(e) != "1474a7" {
		t.Log(string(e))
		t.Error("got:", string(e), ", want: 1474a7")
	}
}

func TestDecode(t *testing.T) {
	var d uint64
	var err error

	// base 4
	b4, err := NewBase([]rune("0123"))
	if err != nil {
		t.Log(err)
	}

	d, err = b4.Decode([]byte("3333"))
	if err != nil {
		t.Log(err)
	} else if d != 255 {
		t.Log(d)
		t.Error("got:", d, ", want: 3333")
	}

	d, err = b4.Decode([]byte("10000"))
	if err != nil {
		t.Log(err)
	} else if d != 256 {
		t.Log(d)
		t.Error("got:", d, ", want: 10000")
	}


	d, err = b4.Decode([]byte("11013102213"))
	if err != nil {
		t.Log(err)
	} else if d != 1340583 {
		t.Log(d)
		t.Error("got:", d, ", want: 1340583")
	}

	// base 16
	b16, _ := NewBase([]rune("0123456789abcdef"))
	d, err = b16.Decode([]byte("ff"))
	if d != 255 {
		t.Log(d)
		t.Error("got:", d, ", want: ff")
	}

	d, err = b16.Decode([]byte("1474a7"))
	if err != nil {
		t.Log(err)
	} else if d != 1340583 {
		t.Log(d)
		t.Error("got:", d, ", want: 1340583")
	}
}