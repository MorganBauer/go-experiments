package main

import "fmt"

func isPalindrome(s string) bool {
     for index := range s {
     	 if s[index] != s[len(s) - index - 1] {
	    return false
	 }
     }
     return true
}

func main () {
     fmt.Println(true)
     fmt.Println(false)
     fmt.Println(isPalindrome("amanaplanacanalpanama"))
     fmt.Println(isPalindrome("paul"))
}