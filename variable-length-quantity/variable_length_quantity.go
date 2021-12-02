package variablelengthquantity

import (
	"fmt"
	"strconv"
)

//EncodeVarint encodes
func EncodeVarint(input []uint32) []byte {
	encoded := make([]byte, 0)
	for _, i := range input {
		encoded = append(encoded, encodeVarint(i)...)
	}
	return encoded
}

//DecodeVarint encodes
func DecodeVarint(input []byte) ([]uint32, error) {
	if input[len(input)-1]>>7 != 0 {
		return nil, fmt.Errorf("invalid sequence")
	}

	decoded := make([]uint32, 0)

	temp := ""

	for _, b := range input {
		temp += fmt.Sprintf("%07b", b&127)
		if b>>7 == 0 {
			v, _ := strconv.ParseInt(temp, 2, 64)
			decoded = append(decoded, uint32(v))
			temp = ""
		}
	}

	return decoded, nil
}

func encodeVarint(input uint32) []byte {
	encoded := make([]byte, 0)
	for {
		byt := input & 127
		encoded = append(encoded, byte(byt))
		input >>= 7
		if input == 0 {
			break
		}
	}
	reversed := make([]byte, 0)
	for i := len(encoded) - 1; i >= 0; i-- {
		if i == 0 {
			reversed = append(reversed, encoded[i])
		} else {
			reversed = append(reversed, encoded[i]|128)
		}
	}
	return reversed
}
