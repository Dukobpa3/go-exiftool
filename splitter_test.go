package exiftool

import (
	"testing"
)

func TestJson(t *testing.T) {
	// flat
	token := "{'json':'like'}"
	a, b, c := JsonSplitter([]byte(token), true)
	t.Log(a)
	t.Log(string(b))
	t.Log(c)

	// prettyprint
	token = "{\n\t'json':'like'\n\t}"
	a, b, c = JsonSplitter([]byte(token), true)
	t.Log(a)
	t.Log(string(b))
	t.Log(c)

	// multiline
	token = "{'json':'like'}, {'json':'like'}"
	a, b, c = JsonSplitter([]byte(token), true)
	t.Log(a)
	t.Log(string(b))
	t.Log(c)

	// multiline
	token = "12.76\n" + string(endPattern) + "\n"
	a, b, c = JsonSplitter([]byte(token), true)
	t.Log(a)
	t.Log(string(b))
	t.Log(c)
}

func TestRegular(t *testing.T) {
	// flat
	token := "======== ./_MG_5111.JPG\nMIME Type                       : image/jpeg"
	a, b, c := JsonSplitter([]byte(token), true)
	t.Log(a)
	t.Log(string(b))
	t.Log(c)

	// multiline
	token = "======== ./_MG_5111.JPG\nMIME Type                       : image/jpeg\n======== ./_MG_5111.JPG\nMIME Type                       : image/jpeg"
	a, b, c = JsonSplitter([]byte(token), true)
	t.Log(a)
	t.Log(string(b))
	t.Log(c)
}
