package main

import "fmt"

func main() {
	// Integer Types
	var integerNum int = 42 // Platform dependent (32 or 64 bit)
	// var int8Num int8 = 127          // -128 to 127
	// var int16Num int16 = 32767      // -32768 to 32767
	// var int32Num int32 = 2147483647 // -2147483648 to 2147483647
	// var int64Num int64 = 9223372036854775807

	fmt.Printf("Integer: %v\n", integerNum) // %v is default format
	fmt.Printf("Integer: %d\n", integerNum) // %d is integer format
}
