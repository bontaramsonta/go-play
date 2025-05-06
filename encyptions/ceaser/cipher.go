package ceaser

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	result := ""
	for _, char := range text {
		result += getOffsetChar(char, key)
	}
	return result
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	offsetVal := ((int(c) - int('a')) + offset%26) % 26
	if offsetVal < 0 {
		offsetVal += 26
	}
	return string(alphabet[offsetVal])
}
