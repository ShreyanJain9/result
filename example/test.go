package main

import (
	"fmt"
	"github.com/ShreyanJain9/result"
	"os"
)

func WriteStringToFile(filename string, data string) (res result.Result[int]) {
	defer result.Catch(&res)

	result.Try(fmt.Println(filename, data))

	file := result.Try(os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644))
	result.Try(file.Write([]byte(data)))

	return result.Ok(0)
}

func main() {
	WriteStringToFile("hello.txt", "hello").Print()
}
