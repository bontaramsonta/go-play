package encodings

func Base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	n := int(bits)
	if n > 7 {
		return ""
	}
	return string(base8Alphabet[n])
}
