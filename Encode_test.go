package baseN

import (
	"testing"
)

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
