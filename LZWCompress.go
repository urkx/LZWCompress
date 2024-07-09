// LZW compression and decompression module
package lzwcompress

// String -> UInt16 conversion: https://yuminlee2.medium.com/go-string-rune-and-byte-efd2aa6034f6

// LZW compression process
//
// Input:
// data string: Data to compress
//
// Returns:
// List containing the bytes (16-bit) to which input data has been encoded
func Compress(data string) []uint16 {
	// Build the initial dictionary with all 8-bit ASCII characters
	const size uint16 = 256 
	dict := make(map[string]uint16)
	for i := range size {
		dict[string(rune(i))] = i
	}
	// Add EOF character
	dict["EOF"] = size + 1
	idx := size + 1

	res := []uint16{}
	sec := ""
	for _, char := range data {
		aux := sec + string(char)
		_, ok := dict[aux]
		if ok {
			sec = aux
		} else {
			res = append(res, dict[sec])
			dict[aux] = idx
			idx++
			sec = string(char)
		}
	}
	
	if sec != "" {
		res = append(res, dict[sec])
	}

	return res
}

// LZW decompression process
//
// Input:
// data []uint16: LZW-encoded bytes (16-bit) list to decode
// 
// Returns:
// Original string decoded
func Decompress(data []uint16) string {
	// Build the initial dictionary with all 8-bit ASCII characters
	const size uint16 = 256
	dict := make(map[uint16]string)
	for i := range size {
		dict[i] = string(rune(i))
	}
	// Add EOF character
	dict[size] = "EOF"
	idx := size + 1
	
	cadena := string(rune(data[0])) // Start with first element of data
	decomp := cadena
	pop_data := append(data[:0], data[1:]...) // Remove first element of data in order to process it only 1 time
	for _, d := range pop_data {
		subpalabra := ""
		val, ok := dict[uint16(d)]
		if ok {
			subpalabra = val
		} else if uint16(d) == idx {
			subpalabra = cadena + string(rune(cadena[0]))
		}
		decomp = decomp + subpalabra
		dict[idx] = cadena + string(rune(subpalabra[0]))
		idx++
		cadena = subpalabra
	}
	return decomp
}