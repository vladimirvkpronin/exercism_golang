package cipher

import (
	//"math/rand"
	"strings"
	"unicode"
)

// Cipher interface defines the Encode and Decode methods

// shift implements a shift cipher with a fixed distance
type shift struct {
	distance int
}

// vigenere implements a Vigenère cipher with a key
type vigenere struct {
	key string
}

// NewCaesar creates a Caesar cipher (shift of 3)
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift creates a shift cipher with the given distance
// Distance must be between 1-25 or -1 to -25, not 0
func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return shift{distance: distance}
}

// NewVigenere creates a Vigenère cipher with the given key
// Key must contain only lowercase letters and not be all 'a's
func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	
	hasNonA := false
	for _, char := range key {
		if !unicode.IsLower(char) {
			return nil
		}
		if char != 'a' {
			hasNonA = true
		}
	}
	
	if !hasNonA {
		return nil
	}
	
	return vigenere{key: key}
}

// Encode implements shift cipher encoding
func (c shift) Encode(input string) string {
	return c.process(input, true)
}

// Decode implements shift cipher decoding
func (c shift) Decode(input string) string {
	return c.process(input, false)
}

// process handles both encoding and decoding for shift cipher
func (c shift) process(input string, encode bool) string {
	var result strings.Builder
	distance := c.distance
	if !encode {
		distance = -distance
	}
	
	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			// Apply shift with modular arithmetic
			shifted := int(char) + distance
			if shifted > 'z' {
				shifted -= 26
			} else if shifted < 'a' {
				shifted += 26
			}
			result.WriteRune(rune(shifted))
		}
	}
	
	return result.String()
}

// Encode implements Vigenère cipher encoding
func (v vigenere) Encode(input string) string {
	return v.process(input, true)
}

// Decode implements Vigenère cipher decoding
func (v vigenere) Decode(input string) string {
	return v.process(input, false)
}

// process handles both encoding and decoding for Vigenère cipher
func (v vigenere) process(input string, encode bool) string {
	var result strings.Builder
	keyIndex := 0
	
	for _, char := range strings.ToLower(input) {
		if unicode.IsLetter(char) {
			// Get the shift distance from the current key character
			keyChar := v.key[keyIndex%len(v.key)]
			shift := int(keyChar - 'a')
			if !encode {
				shift = -shift
			}
			
			// Apply the shift
			shifted := int(char) + shift
			if shifted > 'z' {
				shifted -= 26
			} else if shifted < 'a' {
				shifted += 26
			}
			
			result.WriteRune(rune(shifted))
			keyIndex++
		}
	}
	
	return result.String()
}