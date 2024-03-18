package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("mega  super world!")
	result := buffer.String()
	fmt.Println(result)

}
