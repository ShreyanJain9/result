package main

import (
	"fmt"
	"github.com/ShreyanJain9/result"
)

func WriteStringToFile(filename Result[string], data Result[string]) (res Result[int]) {
	defer result.Catch(&res)

	filename := filename.Throw()
	data := data.Throw()

	result.Wrap(fmt.Println(filename)).Throw()
	result.Wrap(fmt.Println(data)).Throw()

}
