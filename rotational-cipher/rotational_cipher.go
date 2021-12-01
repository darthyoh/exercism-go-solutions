package rotationalcipher

import "unicode"

//RotationalCipher encode an input string with a key
func RotationalCipher(input string, key int) string {
	encoded := ""
	for _, r := range input {
		if unicode.IsLetter(r) {
			offset := rune(97)
			if unicode.IsUpper(r) {
				offset = rune(65)
			}
			encoded += string(((r - offset + rune(key)) % 26) + offset)
		} else {
			encoded += string(r)
		}

	}
	return encoded
}
