package xor

func XORStreamCrypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result) // Ensure result channel is closed when function exits

	for {
		textByte, textOk := <-textCh
		if !textOk {
			// textCh is closed or empty
			return
		}

		keyByte, keyOk := <-keyCh
		if !keyOk {
			// keyCh is closed or empty
			return
		}

		// Perform XOR operation
		xorResult := textByte ^ keyByte
		result <- xorResult
	}
}
