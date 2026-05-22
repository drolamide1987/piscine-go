package main

import "fmt"

func CountAlpha(arg string) int {

	count := 0
	for _, char:= range arg {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			count++
		}
	}
	return count
}

// func CountAlpha(s string) int {
// 	t := 0
// 	for _, c := range s {
// 		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
// 			t++
// 		}
// 	}
// 	return t
// }

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}