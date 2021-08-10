package main

import (
	"fmt"
	"io"
)

type phoneReader string

func (phone phoneReader) Read(p []byte) (int, error) {

	count := 0
	for i := 0; i < len(phone); i++ {
		if phone[i] >= '0' && phone[i] <= '9' {
			p[count] = phone[i]
			count++
		}
	}

	return count, io.EOF
}

func main() {
	var phone phoneReader // Enter the phone in console for example
	fmt.Print("Enter the num: ")
	fmt.Scan(&phone)

	buffer := make([]byte, len(phone))
	phone.Read(buffer)
	fmt.Println(string(buffer))
}