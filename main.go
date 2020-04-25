package main

import "log"

type Generator struct {
}

// NewGenerator returns a pointer to a new generator
func NewGenerator() *Generator {
	return &Generator{}
}

// isPalindrome returns if a string is a palindrome
func (generator *Generator) isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// findAndInsert finds where the first dangler is, and then tries to insert a
// rune to fix it
func (generator *Generator) findAndInsert(s string) (int, string) {
	if generator.isPalindrome(s) {
		return 0, s
	}
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			frontString := s[:i] + string(rune(s[len(s)-1-i])) + s[i:]
			backString := s[:len(s)-i] + string(rune(s[i])) + s[len(s)-i:]
			frontCount, frontResult := generator.findAndInsert(frontString)
			backCount, backResult := generator.findAndInsert(backString)
			if frontCount < backCount {
				return frontCount + 1, frontResult
			} else {
				return backCount + 1, backResult
			}
		}
	}
	return -1, s
}

// GenerateFrom generates the shortest pallindrome from a string
func (generator *Generator) GenerateFrom(s string) (int, string) {
	// Early check
	if generator.isPalindrome(s) {
		return 0, s
	}
	return generator.findAndInsert(s)
}

func main() {
	gen := NewGenerator()
	log.Println(gen.GenerateFrom("TATTA"))
}
