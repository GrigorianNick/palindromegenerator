package main

import (
	"log"
	"sync"
)

// GenerateFrom generates the shortest pallindrome from a string
func GenerateFrom(s string) (int, string) {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			var frontCount, backCount int
			var frontResult, backResult string
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				defer wg.Done()
				frontString := s[:i] + string(rune(s[len(s)-1-i])) + s[i:]
				frontCount, frontResult = GenerateFrom(frontString)
			}()
			go func() {
				defer wg.Done()
				backString := s[:len(s)-i] + string(rune(s[i])) + s[len(s)-i:]
				backCount, backResult = GenerateFrom(backString)
			}()
			wg.Wait()
			if frontCount < backCount {
				return frontCount + 1, frontResult
			} else {
				return backCount + 1, backResult
			}
		}
	}
	return 0, s
}

func main() {
	log.Println(GenerateFrom("GOB"))
}
