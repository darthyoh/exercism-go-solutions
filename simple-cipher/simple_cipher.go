package cipher

import (
	"regexp"
	"strings"
)

//CaesarCipher type
type CaesarCipher struct {
	shift rune
}

//Encode with CaesarCipher
func (c CaesarCipher) Encode(s string) string {
	return c.code(s, true)
}

func (c CaesarCipher) code(s string, encode bool) string {
	s = strings.ToLower(s)
	var encoded strings.Builder

	for _, r := range s {
		if r < 97 || r > 122 {
			continue
		}
		if encode {
			r += c.shift

		} else {
			r -= c.shift
		}
		if r < 97 {
			r += 26
		}
		if r > 122 {
			r -= 26
		}
		encoded.WriteRune(r)
	}
	return encoded.String()
}

//Decode with CaesarCipher
func (c CaesarCipher) Decode(s string) string {
	return c.code(s, false)
}

//NewCaesar generates new cipher
func NewCaesar() Cipher {
	return CaesarCipher{3}
}

//VigenereCipher type
type VigenereCipher struct {
	key []rune
}

//NewShift generates cipher
func NewShift(i int) Cipher {
	if i >= 26 || i <= -26 || i == 0 {
		return nil
	}
	return CaesarCipher{rune(i)}
}

//NewVigenere generates cipher
func NewVigenere(k string) Cipher {
	if len(k) == 0 {
		return nil
	}
	rr := regexp.MustCompile(`^a+$`)
	if rr.MatchString(k) {
		return nil
	}
	key := make([]rune, 0)
	for _, r := range k {
		if r < 97 || r > 122 {
			return nil
		}
		key = append(key, r-97)
	}
	return VigenereCipher{key}
}

func (v VigenereCipher) code(s string, encode bool) string {
	s = strings.ToLower(s)

	var cleaned strings.Builder
	for _, r := range s {
		if r >= 97 && r <= 122 {
			cleaned.WriteRune(r)
		}
	}
	s = cleaned.String()
	var encoded strings.Builder
	for i, r := range s {
		if encode {
			r += v.key[i%len(v.key)]
		} else {
			r -= v.key[i%len(v.key)]
		}
		if r > 122 {
			r -= 26
		}
		if r < 97 {
			r += 26
		}
		encoded.WriteRune(r)
	}
	return encoded.String()
}

//Encode with vigenere
func (v VigenereCipher) Encode(s string) string {
	return v.code(s, true)
}

//Decode with vigenere
func (v VigenereCipher) Decode(s string) string {
	return v.code(s, false)
}
