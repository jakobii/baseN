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

	// base 4
	b4, _ := NewBase([]rune("0123"))
	d = b4.Decode([]byte("3333"))
	if d != 255 {
		t.Log(d)
		t.Error("got:", d, ", want: 3333")
	}

	d = b4.Decode([]byte("10000"))
	if d != 256 {
		t.Log(d)
		t.Error("got:", d, ", want: 10000")
	}

	d = b4.Decode([]byte("11013102213"))
	if d != 1340583 {
		t.Log(d)
		t.Error("got:", d, ", want: 1340583")
	}

	// base 16
	b16, _ := NewBase([]rune("0123456789abcdef"))
	d = b16.Decode([]byte("ff"))
	if d != 255 {
		t.Log(d)
		t.Error("got:", d, ", want: ff")
	}

	d = b16.Decode([]byte("1474a7"))
	if d != 1340583 {
		t.Log(d)
		t.Error("got:", d, ", want: 1340583")
	}
}

/*
func TestEncode(t *testing.T) {

	//base62 [0-9a-zA-Z]
	b62 := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b62r := Encode(62, b62)
	if b62r != "10" {
		t.Log(b62r)
		t.Error("Encode was incorrect, got:", b62r, ", want:", 10)
	}

	b62r = Encode(61, b62)
	if b62r != "Z" {
		t.Log(b62r)
		t.Error("Encode was incorrect, got:", b62r, ", want: Z")
	}

	// base 4  in emojis
	b4 := []rune("\xF0\x9F\x98\x81" + "\xF0\x9F\x98\xA1" + "\xF0\x9F\x98\xBC" + "\xF0\x9F\x98\x8B")

	b4r := Encode(255, b4)
	if b4r != "😋😋😋😋" {
		t.Log(b4r)
		t.Error("Encode was incorrect, got:", b4r, ", want: 😋😋😋😋")
	}

}
*/
