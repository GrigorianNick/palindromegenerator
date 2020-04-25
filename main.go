package main

import (
	"log"
	"sync"
)

// isPalindrome returns if a string is a palindrome
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// findAndInsert finds where the first dangler is, and then tries to insert a
// rune to fix it
func findAndInsert(s string) (int, string) {
	if IsPalindrome(s) {
		return 0, s
	}
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			var frontCount, backCount int
			var frontResult, backResult string
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				defer wg.Done()
				frontString := s[:i] + string(rune(s[len(s)-1-i])) + s[i:]
				frontCount, frontResult = findAndInsert(frontString)
			}()
			go func() {
				defer wg.Done()
				backString := s[:len(s)-i] + string(rune(s[i])) + s[len(s)-i:]
				backCount, backResult = findAndInsert(backString)
			}()
			wg.Wait()
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
func GenerateFrom(s string) (int, string) {
	return findAndInsert(s)
}

func main() {
	log.Println(GenerateFrom("TABABABAcT"))
}
