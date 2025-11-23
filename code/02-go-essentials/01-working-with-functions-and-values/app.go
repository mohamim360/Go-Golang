package main

import "fmt" // fmt package provides functions for formatted I/O (printing, scanning, etc.)

// main function — the entry point of every Go program
func main() {
	fmt.Print("Hello World!") // Print() prints text to the console without a newline at the end

	fmt.Print('Hello World!') // ❌ Invalid: single quotes are for single characters only

	fmt.Print(`Hello World!`) // ✔ Raw string literal: keeps text exactly as written


}

