package lzwcompress

import (
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	res := Compress("tres tristes tigres tragaban trigo en un trigal")
	want := []uint16{116, 114, 101, 115, 32, 257, 105, 115, 116, 259, 261, 105, 103, 258, 260, 257, 97, 103, 97, 98, 97, 110, 261, 114, 268, 111, 32, 101, 278, 117, 278, 262, 274, 108}
	if len(res) != len(want) {
		t.Fatal("Result size != want size")
	}

	if !reflect.DeepEqual(want, res) {
		t.Fatal("Result != want")
	}
}

func TestDecompress(t *testing.T) {
	res := Decompress([]uint16{116, 114, 101, 115, 32, 257, 105, 115, 116, 259, 261, 105, 103, 258, 260, 257, 97, 103, 97, 98, 97, 110, 261, 114, 268, 111, 32, 101, 278, 117, 278, 262, 274, 108})
	want := "tres tristes tigres tragaban trigo en un trigal"

	if res != want {
		t.Fatalf("Decompression process failed: %s", res)
	}
}